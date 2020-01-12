package models

import (
	"cc-lunch-backend/config"
	"cc-lunch-backend/entities"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LunchOrderModel struct {
	Db         *mongo.Database
	Collection string
}

func (lunchOrderModel LunchOrderModel) FindAll() ([]entities.LunchOrder, error) {
	fmt.Println("==========FindAll()")
	collection := lunchOrderModel.Db.Collection(lunchOrderModel.Collection)
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.DBRequestDuration)
	defer cancelFunc()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	lunchOrders := []entities.LunchOrder{}
	for cursor.Next(ctx) {
		var lo entities.LunchOrder
		cursor.Decode(&lo)

		if lo.OptionID != nil {
			lunchOrders = append(lunchOrders, lo)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return lunchOrders, nil
}

func (lunchOrderModel LunchOrderModel) FindByUser(userID primitive.ObjectID) ([]entities.LunchOrder, error) {
	fmt.Println("==========FindByUser()")

	collection := lunchOrderModel.Db.Collection(lunchOrderModel.Collection)
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.DBRequestDuration)
	defer cancelFunc()
	cursor, err := collection.Find(ctx, entities.LunchOrder{UserID: userID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	lunchOrders := []entities.LunchOrder{}
	for cursor.Next(ctx) {
		var lo entities.LunchOrder
		cursor.Decode(&lo)

		if lo.OptionID != nil {
			lunchOrders = append(lunchOrders, lo)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return lunchOrders, nil
}

func (lunchOrderModel LunchOrderModel) Create(lunchOrder entities.LunchOrder) (entities.LunchOrder, error) {
	fmt.Println("==========Create()")

	var existingLunchOrder entities.LunchOrder
	collection := lunchOrderModel.Db.Collection(lunchOrderModel.Collection)
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.DBRequestDuration)
	defer cancelFunc()

	err := collection.FindOne(ctx, entities.LunchOrder{UserID: lunchOrder.UserID, Day: lunchOrder.Day}).Decode(&existingLunchOrder)
	if err != nil && err != mongo.ErrNoDocuments {
		return entities.LunchOrder{}, err
	}

	// user&&day is NOT existing? create : update
	var createUpdateErr error
	if err == mongo.ErrNoDocuments {
		_, createUpdateErr = collection.InsertOne(ctx, lunchOrder)
	} else {
		_, createUpdateErr = collection.UpdateOne(ctx, bson.D{{"_id", existingLunchOrder.ID}}, bson.D{{"$set", bson.D{{"option_id", lunchOrder.OptionID}}}})
	}

	if createUpdateErr != nil {
		return entities.LunchOrder{}, createUpdateErr
	}
	return lunchOrder, nil
}
