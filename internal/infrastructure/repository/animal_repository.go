package repository

import (
	"encoding/json"
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
)

type AnimalHandler struct {
	petRepo *PetRepository
}

func NewAnimalHandler(petRepo *PetRepository) *AnimalHandler {
	return &AnimalHandler{petRepo: petRepo}
}

func (h *AnimalHandler) CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	var newAnimal animal.Animal
	if err := json.NewDecoder(r.Body).Decode(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode adicionar validações adicionais para os dados do animal se necessário.

	if err := h.petRepo.CreateAnimal(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)    // Retorna o status 201 Created
	json.NewEncoder(w).Encode(newAnimal) // Retorna o animal criado com o ID atribuído
}
