package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:  []string{os.Getenv("ALLOWED_HOSTS")},
		SSLRedirect:   true,
		STSSeconds:    31536000,
		FrameDeny:     true,
		IsDevelopment: false,
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			if err != nil {
				c.Abort()
				return
			}

			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	r := gin.Default()
	auth := gin.BasicAuth(gin.Accounts{
		"username": "password",
	})
	r.Use(auth)
	r.Use(secureFunc)

	r.GET("/device-definitions/:id", readDefinition)

	r.POST("/device-definitions/create", createDefinition)

	r.PUT("/device-definitions/:id", updateDefinition)

	r.DELETE("/device-definitions/:id", deleteDefinition)

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":8004"))
}
