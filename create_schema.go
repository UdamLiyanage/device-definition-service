package main

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func createDefinition(c echo.Context) error {
	var schema DeviceSchema
	if err := json.NewDecoder(c.Request().Body).Decode(&schema); err != nil {
		panic(err)
	}
	schema.CreatedAt, schema.UpdatedAt = time.Now(), time.Now()
	insertResult, err := collection.InsertOne(context.TODO(), schema)
	if checkError(err) {
		return c.JSON(500, err.Error())
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		schema.ID = oid
		return c.JSON(201, schema)
	} else {
		return c.JSON(500, err)
	}
}
