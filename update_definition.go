package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateDefinition(c echo.Context) error {
	var definition DeviceDefinition
	if err := c.Bind(definition); err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"version":    definition.Version,
			"name":       definition.Name,
			"type":       definition.Type,
			"events":     definition.Events,
			"actions":    definition.Actions,
			"parameters": definition.Parameters,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}
