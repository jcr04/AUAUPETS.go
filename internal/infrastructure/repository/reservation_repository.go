package repository

import (
	"AUAUPets/internal/domain/reservation"
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
