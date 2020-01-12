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

const collectionName = "lunchorder"

func GetAllLunchOrders(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		lunchOrderModel := models.LunchOrderModel{
			Db:         db,
			Collection: collectionName,
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
			Collection: collectionName,
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
			Collection: collectionName,
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

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
