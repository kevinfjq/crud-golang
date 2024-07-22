package repository

import (
	"context"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/model/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) UpdateUser(id string, domain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init updateUserRepository repository", zap.String("journey", "updateUserRepository"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(domain)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "updateUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	logger.Info("updateUser repository executed successfully", zap.String("userId", id), zap.String("journey", "updateUser"))
	return ur.FindUserByID(id)
}
