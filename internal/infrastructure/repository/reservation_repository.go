package repository

import (
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
