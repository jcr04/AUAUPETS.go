package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
)

func CreateAnimalHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create a new animal
	// ...
	newAnimal := animal.NewAnimal("1", "Fido", "Labrador")
	json.NewEncoder(w).Encode(newAnimal)
}
