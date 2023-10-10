package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

func CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	var newAnimal animal.Animal
	if err := json.NewDecoder(r.Body).Decode(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	petRepo := repository.NewPetRepository()
	if err := petRepo.CreateAnimal(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newAnimal) // Retorne o animal criado com o ID atribuído
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
