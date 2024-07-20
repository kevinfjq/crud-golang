package repository

import (
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	//UpdateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
