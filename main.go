package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/routes"
)

func main() {
	logs.Info("About to start aplication")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
