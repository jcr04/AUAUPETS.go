package repository

import (
	"errors"

	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
)

type ReservationRepository struct {
	// ...
}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{}
}

func (repo *ReservationRepository) Save(res *reservation.Reservation) error {
	// Simulate a database save
	// ...
	return nil
}

func (repo *ReservationRepository) GetAllReservations() ([]*reservation.Reservation, error) {
	// Simulate a database fetch
	return []*reservation.Reservation{
		// ... list of reservations
	}, nil
}

func (repo *ReservationRepository) GetReservationByID(id string) (*reservation.Reservation, error) {
	// Simular uma busca no banco de dados
	if id == "1" {
		return &reservation.Reservation{
			ID: "1",
			// ... outros campos
		}, nil
	}

	// Se a reserva não foi encontrada ou outro erro ocorreu
	return nil, errors.New("reservation not found")
}
