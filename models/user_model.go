package models

import (
	"cc-lunch-backend/config"
	"cc-lunch-backend/entities"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Db         *mongo.Database
	Collection string
}

func (userModel UserModel) FindUser(username, password string) (entities.User, error) {
	fmt.Println("==========FindUser()")

	collection := userModel.Db.Collection(userModel.Collection)
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.DBRequestDuration)
	defer cancelFunc()

	var user entities.User
	err := collection.FindOne(ctx, entities.User{Username: username}).Decode(&user)
	if err != nil {
		return entities.User{}, err
	}

	if user.Password != password {
		return entities.User{}, fmt.Errorf("Invalid username/password")
	}

	return user, nil
}
