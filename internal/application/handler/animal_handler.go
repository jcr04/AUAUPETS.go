package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

func CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create a new animal
	// ...
	newAnimal := animal.NewAnimal("1", "Fido", "Labrador", "15")
	json.NewEncoder(w).Encode(newAnimal)
}

func GetAnimalHandler(w http.ResponseWriter, r *http.Request) {
	petRepo := repository.NewPetRepository()
	pet, err := petRepo.GetPetByID("1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pet)
}
