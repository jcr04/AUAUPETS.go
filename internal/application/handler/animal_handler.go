package handler

import (
	"AUAUPets/internal/domain/animal"
	"encoding/json"
	"net/http"
)

func CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create a new animal
	// ...
	newAnimal := animal.NewAnimal("1", "Fido", "Labrador")
	json.NewEncoder(w).Encode(newAnimal)
}
