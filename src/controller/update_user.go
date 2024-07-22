package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/configuration/validation"
	"github.com/kevinfjq/crud-golang/src/controller/model/request"
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		restError := rest_err.NewBadRequestError("User id is not a valid hex id")
		c.JSON(restError.Code, restError)
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	updatedDomain, err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err, zap.String("jorney", "updateUser"))
		c.JSON(err.Code, err)
	}

	logger.Info("updateUser controller executed successfully", zap.String("journey", "updateUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(updatedDomain))
}
