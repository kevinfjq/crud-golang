package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/validation"
	"github.com/kevinfjq/crud-golang/src/controller/model/request"
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/view"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	createdDomain, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err, zap.String("jorney", "createUser"))
		c.JSON(err.Code, err)
	}

	logger.Info("CreateUser controller executed successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(createdDomain))
}
