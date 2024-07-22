package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validade userId", err, zap.String("journey", "deleteUser"))
		errorMessage := rest_err.NewBadRequestError("User id is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("deleteUser controller executed successfully", zap.String("journey", "deleteUser"))
	c.Status(http.StatusOK)
}
