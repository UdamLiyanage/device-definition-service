package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateDefinition(c *gin.Context) {
	var definition DeviceDefinition
	err := json.NewDecoder(c.Request.Body).Decode(&definition)
	checkError(err, c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
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
	checkError(err, c)
	c.JSON(200, res)
}
