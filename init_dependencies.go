package main

import (
	"github.com/kevinfjq/crud-golang/src/controller"
	"github.com/kevinfjq/crud-golang/src/model/repository"
	"github.com/kevinfjq/crud-golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(dbConnection *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(dbConnection)
	serv := service.NewUserDomainService(userRepository)
	return controller.NewUserControllerInterface(serv)
}
