package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/golang-default/src/configuration/logger"
	"github.com/jimmmisss/golang-default/src/controller/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("About to start user application")
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
