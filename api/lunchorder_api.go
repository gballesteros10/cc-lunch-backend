package api

import (
	"cc-lunch-backend/config"
	"cc-lunch-backend/entities"
	"cc-lunch-backend/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const colLunchOrder = "lunchorder"

func GetAllLunchOrders(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		lunchOrderModel := models.LunchOrderModel{
			Db:         db,
			Collection: colLunchOrder,
		}

		lunchOrders, findErr := lunchOrderModel.FindAll()
		if findErr != nil {
			respondWithError(w, http.StatusBadRequest, findErr.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, lunchOrders)
		}
	}
}

func GetLunchOrdersByUser(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		lunchOrderModel := models.LunchOrderModel{
			Db:         db,
			Collection: colLunchOrder,
		}

		params := mux.Vars(r)
		userID, _ := primitive.ObjectIDFromHex(params["user_id"])
		lunchOrders, findErr := lunchOrderModel.FindByUser(userID)
		if findErr != nil {
			respondWithError(w, http.StatusBadRequest, findErr.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, lunchOrders)
		}
	}
}

func CreateLunchOrder(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		lunchOrderModel := models.LunchOrderModel{
			Db:         db,
			Collection: colLunchOrder,
		}

		var lunchOrder entities.LunchOrder
		json.NewDecoder(r.Body).Decode(&lunchOrder)

		lunchOrder, findErr := lunchOrderModel.Create(lunchOrder)
		if findErr != nil {
			respondWithError(w, http.StatusBadRequest, findErr.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, lunchOrder)
		}
	}
}
