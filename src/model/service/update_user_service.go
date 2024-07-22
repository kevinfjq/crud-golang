package service

import (
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(id string, userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init updateUser services", zap.String("journey", "updateUserService"))
	return ud.userRepository.UpdateUser(id, userDomain)
}
