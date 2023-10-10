package api

import (
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/application/handler"
)

func StartServer() {
	http.HandleFunc("/animals", handler.CreateAnimalHandler)
	http.HandleFunc("/reservations", handler.CreateReservationHandler)
	http.ListenAndServe(":8080", nil)
}
