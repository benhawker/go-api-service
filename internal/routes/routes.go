package routes

import (
	"github.com/gorilla/mux"
	"github.com/benhawker/go-api-service/internal/handlers/friends"
)

func Register(router *mux.Router) {
	friends := friends.NewHandler()
	subscriptions := subscriptions.NewHandler()
	updates := updates.NewHandler()

	router.HandleFunc("/friends", handler.Show)
	router.HandleFunc("/friends", handler.Create).Methods("POST")

	router.HandleFunc("/subscriptions", handler.Create).Methods("POST")
	router.HandleFunc("/subscriptions", handler.Index)

	router.HandleFunc("/friends", handler.Index)
	router.HandleFunc("/friends", handler.Index)

}