package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

type AnimalHandler struct {
	petRepo *repository.PetRepository
}

func NewAnimalHandler(petRepo *repository.PetRepository) *AnimalHandler {
	return &AnimalHandler{petRepo: petRepo}
}

func (h *AnimalHandler) CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	var newAnimal animal.Animal
	if err := json.NewDecoder(r.Body).Decode(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.petRepo.CreateAnimal(&newAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newAnimal) // Retorne o animal criado com o ID atribuído
}

func (h *AnimalHandler) GetAnimalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pet, err := h.petRepo.GetPetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pet)
}

func (h *AnimalHandler) ListAnimalsHandler(w http.ResponseWriter, r *http.Request) {
	animals, err := h.petRepo.ListAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(animals)
}

func (h *AnimalHandler) UpdateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedAnimal animal.Animal
	if err := json.NewDecoder(r.Body).Decode(&updatedAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.petRepo.UpdateAnimal(id, &updatedAnimal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedAnimal) // Retornando o animal atualizado
}

func (h *AnimalHandler) DeleteAnimalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.petRepo.DeleteAnimal(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Retornando um status 204 No Content
}
