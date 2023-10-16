package hosting

import (
	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
)

type Hosting struct {
	ID            string
	Name          string
	Reservation   *reservation.Reservation
	ReservationID string `json:"reservation_id"`
}

func NewHosting(id, name string, reservation *reservation.Reservation) *Hosting {
	return &Hosting{
		ID:          id,
		Name:        name,
		Reservation: reservation,
	}
}
