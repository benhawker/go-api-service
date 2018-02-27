package routes

import (
	"github.com/gorilla/mux"
	"github.com/benhawker/go-api-service/internal/handlers/time-slots"
)

func Register(router *mux.Router) {
	ts := timeslots.NewHandler()
	router.HandleFunc("/time_slots", ts.ServeHTTP)
}