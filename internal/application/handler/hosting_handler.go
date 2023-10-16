package handler

import (
	"encoding/json"
	"log" // Importe o pacote log
	"net/http"

	"github.com/jcr04/AUAUPETS.go/internal/domain/hosting"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/repository"
)

type HostingHandler struct {
	hostingRepo     *repository.HostingRepository
	reservationRepo *repository.ReservationRepository
}

func NewHostingHandler(hostingRepo *repository.HostingRepository, reservationRepo *repository.ReservationRepository) *HostingHandler {
	log.Println("NewHostingHandler chamado") // Log de entrada
	if hostingRepo == nil || reservationRepo == nil {
		log.Println("Repositórios não foram inicializados corretamente")
		return nil
	}
	return &HostingHandler{hostingRepo: hostingRepo, reservationRepo: reservationRepo}
}

func (h *HostingHandler) CreateHostingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateHostingHandler chamado") // Log de entrada

	var newHosting hosting.Hosting
	if err := json.NewDecoder(r.Body).Decode(&newHosting); err != nil {
		log.Printf("Erro decodificando o corpo da requisição: %v\n", err) // Log do erro
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if h.reservationRepo == nil {
		log.Println("reservationRepo é nil")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Use newHosting.ReservationID ao invés de newHosting.Reservation.ID
	res, err := h.reservationRepo.GetReservationByID(newHosting.ReservationID)
	if err != nil {
		log.Printf("Erro ao obter reserva por ID: %v\n", err) // Log do erro
		http.Error(w, "Failed to get reservation by ID", http.StatusInternalServerError)
		return
	}
	newHosting.Reservation = res

	if err := h.hostingRepo.Save(&newHosting); err != nil {
		log.Printf("Erro ao salvar hospedagem: %v\n", err) // Log do erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(newHosting)
	if err != nil {
		log.Printf("Erro codificando resposta JSON: %v\n", err) // Log do erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("CreateHostingHandler concluído com sucesso") // Log de saída
}

func (h *HostingHandler) ListHostingHandler(w http.ResponseWriter, r *http.Request) {
	hostings, err := h.hostingRepo.GetAllHostings()
	if err != nil {
		log.Printf("Erro ao obter todas as hospedagens: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(hostings)
	if err != nil {
		log.Printf("Erro codificando resposta JSON: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
