package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcr04/AUAUPETS.go/internal/application/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/animals", handler.CreateAnimalHandler).Methods("POST")
	return r
}

func StartServer() {
	http.HandleFunc("/animals", handler.CreateAnimalHandler)
	http.HandleFunc("/reservations", handler.CreateReservationHandler)
	http.ListenAndServe(":8080", nil)
}
