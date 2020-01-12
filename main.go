package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"cc-lunch-backend/api"
	"cc-lunch-backend/config"
)

func main() {
	fmt.Println("Application start...")

	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})

	// lunchorder api
	router.HandleFunc("/lunchorder", api.CreateLunchOrder).Methods("POST")
	router.HandleFunc("/lunchorder/{user_id}", api.GetLunchOrdersByUser).Methods("GET")
	router.HandleFunc("/lunchorder", api.GetAllLunchOrders).Methods("GET")

	// user api
	router.HandleFunc("/login", api.LoginUser).Methods("POST")

	err := http.ListenAndServe(config.Port, handlers.CORS(corsObj)(router))
	if err != nil {
		fmt.Println(err)
	}
}
