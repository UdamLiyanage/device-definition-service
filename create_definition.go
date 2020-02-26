package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createDefinition(c echo.Context) error {
	var definition DeviceDefinition
	if err := c.Bind(definition); err != nil {
		return err
	}
	insertResult, err := DB.Collection.InsertOne(context.TODO(), definition)
	if checkError(err) {
		return c.JSON(500, err)
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		definition.ID = oid
		return c.JSON(201, definition)
	} else {
		return c.JSON(500, err)
	}
}
