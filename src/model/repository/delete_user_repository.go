package repository

import (
	"context"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUserRepository"))
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collectionName)
	userId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userId}}

	_, err := collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}
	logger.Info("deleteUser repository executed successfully", zap.String("userId", id), zap.String("journey", "deleteUser"))
	return nil
}
