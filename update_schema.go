package main

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func updateDefinition(c echo.Context) error {
	var schema DeviceSchema
	if err := json.NewDecoder(c.Request().Body).Decode(&schema); err != nil {
		return c.JSON(500, err)
	}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if checkError(err) {
		return c.JSON(500, err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"version":         schema.Version,
			"name":            schema.Name,
			"type":            schema.Type,
			"public":          schema.Public,
			"description":     schema.Description,
			"latest_firmware": schema.LatestFirmware,
			"communication":   schema.DeviceComm,
			"commands":        schema.Commands,
			"command_format":  schema.CommandsFormat,
			"parameters":      schema.Parameters,
		},
	}
	schema.UpdatedAt = time.Now()
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if checkError(err) {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}
