package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createDefinition(c *gin.Context) {
	var definition DeviceDefinition
	err := json.NewDecoder(c.Request.Body).Decode(&definition)
	checkError(err, c)
	insertResult, err := DB.Collection.InsertOne(context.TODO(), definition)
	checkError(err, c)
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		definition.ID = oid
		c.JSON(201, definition)
	} else {
		c.AbortWithStatus(500)
	}
}
