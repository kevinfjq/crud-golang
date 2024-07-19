package view

import (
	"github.com/kevinfjq/crud-golang/src/controller/model/response"
	"github.com/kevinfjq/crud-golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
