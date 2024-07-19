package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinfjq/crud-golang/src/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{serviceInterface}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
