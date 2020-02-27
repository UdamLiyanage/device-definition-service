package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readDefinition(c echo.Context) error {
	var definition DeviceDefinition
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	err = DB.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&definition)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, definition)
}

func readAllDefinitions(c echo.Context) error {
	var definitions []DeviceDefinition
	userID := c.Param("id")
	filter := bson.D{{"uid", userID}}
	cur, err := DB.Collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var definition DeviceDefinition
		err := cur.Decode(&definition)
		if err != nil {
			panic(err)
		}
		definitions = append(definitions, definition)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return c.JSON(200, definitions)
}
