package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevinfjq/crud-golang/src/controller"
	"github.com/kevinfjq/crud-golang/src/controller/routes"
	"github.com/kevinfjq/crud-golang/src/model/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serv := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(serv)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
