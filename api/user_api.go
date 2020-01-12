package api

import (
	"cc-lunch-backend/config"
	"cc-lunch-backend/entities"
	"cc-lunch-backend/models"
	"encoding/json"
	"net/http"
)

const colUser = "user"

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userModel := models.UserModel{
		Db:         db,
		Collection: colUser,
	}

	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	loggedInUser, findErr := userModel.FindUser(user.Username, user.Password)
	if findErr != nil {
		respondWithError(w, http.StatusBadRequest, findErr.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, loggedInUser)
}
