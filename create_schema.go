package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func createDefinition(c echo.Context) error {
	var schema DeviceSchema
	if err := c.Bind(schema); err != nil {
		return err
	}
	schema.CreatedAt, schema.UpdatedAt = time.Now(), time.Now()
	insertResult, err := DB.Collection.InsertOne(context.TODO(), schema)
	if checkError(err) {
		return c.JSON(500, err)
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		schema.ID = oid
		return c.JSON(201, schema)
	} else {
		return c.JSON(500, err)
	}
}
