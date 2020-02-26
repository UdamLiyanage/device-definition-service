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
