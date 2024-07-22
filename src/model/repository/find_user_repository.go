package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/kevinfjq/crud-golang/src/configuration/logger"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
	"github.com/kevinfjq/crud-golang/src/model"
	"github.com/kevinfjq/crud-golang/src/model/entity"
	"github.com/kevinfjq/crud-golang/src/model/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository", zap.String("journey", "findUserByEmail"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmailRepository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository", zap.String("journey", "findUserByID"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByIDRepository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
