package reservation

import (
	"AUAUPets/internal/domain/animal"
	"time"
)

type Reservation struct {
	ID       string
	Animal   *animal.Animal
	CheckIn  time.Time
	CheckOut time.Time
}

func NewReservation(id string, animal *animal.Animal, checkIn, checkOut time.Time) *Reservation {
	return &Reservation{id, animal, checkIn, checkOut}
}
