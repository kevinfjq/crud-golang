package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller", zap.String("journey", "findUserByID"))
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validade userId", err, zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError("User id is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	domain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID service", err, zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByID controller executed successfully", zap.String("journey", "findUserByID"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))
	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validade userEmail", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("User email is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	domain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail service", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("findUserByEmail controller executed successfully", zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
