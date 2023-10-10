package handler

import (
	"AUAUPets/internal/domain/reservation"
	"AUAUPets/internal/infrastructure/repository"
	"encoding/json"
	"net/http"
	"time"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create a new reservation
	// ...
	animalRepo := repository.NewAnimalRepository()
	animal, _ := animalRepo.GetAnimalByID("1")
	checkIn := time.Now()
	checkOut := checkIn.Add(24 * time.Hour)
	newReservation := reservation.NewReservation("1", animal, checkIn, checkOut)
	json.NewEncoder(w).Encode(newReservation)
}
