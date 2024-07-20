package converter

import (
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/model/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
