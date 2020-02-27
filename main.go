package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status} path=${path} latency=${latency_human}\n",
	}))
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == os.Getenv("API_AUTH_USERNAME") && password == os.Getenv("API_AUTH_PASSWORD") {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Recover())

	e.GET("/device-definitions/:id", readDefinition)
	e.GET("/device-definitions/all/:id", readAllDefinitions)

	e.POST("/device-definitions/create", createDefinition)

	e.PUT("/device-definitions/:id", updateDefinition)

	e.DELETE("/device-definitions/:id", deleteDefinition)

	return e
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":8004"))
}
