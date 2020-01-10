package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName                 string        = "cclunch"
	dbURI                  string        = "mongodb://localhost:27017"
	dbCollectionLunchOrder string        = "lunchorder"
	dbRequestDuration      time.Duration = 10 * time.Second
	port                   string        = ":3001"
)

var dbClient *mongo.Client

type LunchOrder struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OptionID primitive.ObjectID `json:"option_id,omitempty" bson:"option_id,omitempty"`
	Day      int                `json:"day,omitempty" bson:"day,omitempty"`
}

func CreateLunchOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var lunchOrder LunchOrder
	json.NewDecoder(request.Body).Decode(&lunchOrder)
	collection := dbClient.Database(dbName).Collection(dbCollectionLunchOrder)
	ctx, _ := context.WithTimeout(context.Background(), dbRequestDuration)

	result, err := collection.InsertOne(ctx, lunchOrder)
	if err != nil {
		fmt.Printf("An error occurred: %+v", err)
	}
	json.NewEncoder(response).Encode(result)
}

func GetLunchOrderByUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	userID, _ := primitive.ObjectIDFromHex(params["user_id"])
	var lunchOrders []LunchOrder

	collection := dbClient.Database(dbName).Collection(dbCollectionLunchOrder)
	ctx, _ := context.WithTimeout(context.Background(), dbRequestDuration)

	cursor, err := collection.Find(ctx, LunchOrder{UserID: userID})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var lunchOrder LunchOrder
		cursor.Decode(&lunchOrder)
		lunchOrders = append(lunchOrders, lunchOrder)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(lunchOrders)
}

func main() {
	fmt.Println("Application start...")

	ctx, _ := context.WithTimeout(context.Background(), dbRequestDuration)
	dbClient, _ = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	router := mux.NewRouter()

	router.HandleFunc("/lunchorder", CreateLunchOrder).Methods("POST")
	router.HandleFunc("/lunchorder/{user_id}", GetLunchOrderByUser).Methods("GET")

	http.ListenAndServe(port, router)
}