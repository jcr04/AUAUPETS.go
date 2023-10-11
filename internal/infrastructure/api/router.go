package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcr04/AUAUPETS.go/internal/application/handler"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Crie uma nova instância de PetRepository
	petRepo := repository.NewPetRepository()
	// Crie uma nova instância de AnimalHandler
	animalHandler := handler.NewAnimalHandler(petRepo)

	// Agora use a instância de AnimalHandler para chamar os métodos
	r.HandleFunc("/animals", animalHandler.CreateAnimalHandler).Methods("POST")
	r.HandleFunc("/animals", animalHandler.ListAnimalsHandler).Methods("GET")
	r.HandleFunc("/animals/{id}", animalHandler.GetAnimalHandler).Methods("GET")
	r.HandleFunc("/animals/{id}", animalHandler.UpdateAnimalHandler).Methods("PUT")
	r.HandleFunc("/animals/{id}", animalHandler.DeleteAnimalHandler).Methods("DELETE")

	return r
}

func StartServer() {
	http.Handle("/", NewRouter()) // Use NewRouter aqui para registrar as rotas
	http.ListenAndServe(":8080", nil)
}
