package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

func generateUniqueReservationID() string {
	// Gere um ID único aleatório
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		panic(err) // Tratar erro de geração de ID único
	}

	// Converta o ID em uma string hexadecimal
	return hex.EncodeToString(id)
}

type ReservationHandler struct {
	reservationRepo *repository.ReservationRepository
	animalRepo      *repository.PetRepository
}

func NewReservationHandler(reservationRepo *repository.ReservationRepository, animalRepo *repository.PetRepository) *ReservationHandler {
	return &ReservationHandler{reservationRepo: reservationRepo, animalRepo: animalRepo}
}

// CreateReservationHandler cria uma nova reserva com base na solicitação.
func (h *ReservationHandler) CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var newReservation reservation.Reservation
	if err := json.NewDecoder(r.Body).Decode(&newReservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Busque o animal no banco de dados com base nas informações da reserva
	animal, err := h.animalRepo.GetPetByID(newReservation.Animal.ID)
	if err != nil {
		http.Error(w, "Failed to get pet by ID", http.StatusInternalServerError)
		return
	}

	// Defina o animal na reserva
	newReservation.Animal = animal

	// Defina a data de check-in e check-out, como você estava fazendo anteriormente
	checkIn := time.Now()
	checkOut := checkIn.Add(24 * time.Hour)

	// Defina as datas de check-in e check-out na reserva
	newReservation.CheckIn = checkIn
	newReservation.CheckOut = checkOut

	// Gere um ID único para a reserva
	newReservation.ID = generateUniqueReservationID() // Você precisa implementar essa função

	if err := h.reservationRepo.Save(&newReservation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newReservation)
}

// ListReservationsHandler lista todas as reservas existentes.
func (h *ReservationHandler) ListReservationsHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenha todas as reservas (substitua com a lógica real)
	reservations, err := h.reservationRepo.GetAllReservations()
	if err != nil {
		http.Error(w, "Failed to get reservations", http.StatusInternalServerError)
		return
	}

	// Configure o cabeçalho da resposta com o tipo de conteúdo JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifique a lista de reservas em JSON e envie-a como resposta
	json.NewEncoder(w).Encode(reservations)
}

// GetReservationHandler obtém uma reserva com base em seu ID.
func (h *ReservationHandler) GetReservationHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenha o ID da reserva da solicitação (substitua com a lógica real)
	vars := mux.Vars(r)
	reservationID := vars["id"]

	// Obtenha a reserva com base no ID (substitua com a lógica real)
	reservation, err := h.reservationRepo.GetReservationByID(reservationID)
	if err != nil {
		http.Error(w, "Failed to get reservation", http.StatusInternalServerError)
		return
	}

	// Configure o cabeçalho da resposta com o tipo de conteúdo JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifique a reserva em JSON e envie-a como resposta
	json.NewEncoder(w).Encode(reservation)
}
