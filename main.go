package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevinfjq/crud-golang/src/configuration/db/mongodb"
	"github.com/kevinfjq/crud-golang/src/controller"
	"github.com/kevinfjq/crud-golang/src/controller/routes"
	"github.com/kevinfjq/crud-golang/src/model/repository"
	"github.com/kevinfjq/crud-golang/src/model/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connnect to database, error=%s", err)
		return
	}
	userRepository := repository.NewUserRepository(dbConnection)
	serv := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(serv)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
