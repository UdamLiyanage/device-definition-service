package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readDefinition(c *gin.Context) {
	var definition DeviceDefinition
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	err = DB.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&definition)
	checkError(err, c)
	c.JSON(200, definition)
}
