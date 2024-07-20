package converter

import (
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/model/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
