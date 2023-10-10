package api

import (
	"net/http"

	"github.com/usuario/go/auaupets/internal/application/handler"
)

func StartServer() {
	http.HandleFunc("/animals", handler.CreateAnimalHandler)
	http.HandleFunc("/reservations", handler.CreateReservationHandler)
	http.ListenAndServe(":8080", nil)
}
