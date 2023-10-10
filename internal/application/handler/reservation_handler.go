package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create a new reservation
	// ...
	petRepo := repository.NewPetRepository() // Alterado de NewAnimalRepository para NewPetRepository
	animal, err := petRepo.GetPetByID("1")   // Alterado de GetAnimalByID para GetPetByID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	checkIn := time.Now()
	checkOut := checkIn.Add(24 * time.Hour)
	newReservation := reservation.NewReservation("1", animal, checkIn, checkOut)
	json.NewEncoder(w).Encode(newReservation)
}

func ListReservationsHandler(w http.ResponseWriter, r *http.Request) {
	reservationRepo := repository.NewReservationRepository()
	reservations, err := reservationRepo.GetAllReservations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reservations)
}

func GetReservationHandler(w http.ResponseWriter, r *http.Request) {
	// ... implement your code
}
