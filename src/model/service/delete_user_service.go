package service

import (
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(id string) *rest_err.RestErr {
	logger.Info("Init deleteUser service", zap.String("journey", "deleteUser"))
	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Error trying to call deleteUser repository", err, zap.String("journey", "deleteUser"))
		return err
	}
	logger.Info("deleteUser service executed successfully", zap.String("userId", id), zap.String("journey", "deleteUser"))
	return nil
}
