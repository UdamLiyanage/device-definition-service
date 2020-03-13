package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strconv"
)

func readDefinition(c echo.Context) error {
	var schema DeviceSchema
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&schema)
	if checkError(err) {
		if err == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
		return c.JSON(500, err)
	}
	return c.JSON(200, schema)
}

func readAllDefinitions(c echo.Context) error {
	var (
		res Pager
		err error
	)
	if lim := c.QueryParam("limit"); lim == "" {
		res.Limit = 20
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(n)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$gt", objID}}},
			}
			res.FirstPage = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$lt", objID}}},
			}
			res.FirstPage = false
		} else {
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
			}
			res.FirstPage = true
		}
	} else {
		res.Limit, err = strconv.ParseInt(lim, 32, 64)
		if err != nil {
			log.Println("Error Occurred: ", err)
			return c.JSON(500, err)
		}
		if n := c.QueryParam("next"); n != "" {
			objID, err := primitive.ObjectIDFromHex(c.QueryParam("next"))
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$gt", objID}}},
			}
			res.FirstPage = false
		} else if p := c.QueryParam("previous"); p != "" {
			objID, err := primitive.ObjectIDFromHex(p)
			if err != nil {
				return c.JSON(500, err)
			}
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
				{"_id", bson.D{{"$lt", objID}}},
			}
			res.FirstPage = false
		} else {
			res.Filter = bson.D{
				{"device_token", c.Param("token")},
			}
			res.FirstPage = true
		}
	}
	res.URL = os.Getenv("SELF_ADDRESS") + c.Request().URL.Path
	return c.JSON(200, res.Paginate())
}
