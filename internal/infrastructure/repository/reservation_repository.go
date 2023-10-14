package repository

import (
	"errors"
	"sync"

	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
)

type ReservationRepository struct {
	// Simula um banco de dados de reservas usando um mapa.
	db     map[string]*reservation.Reservation
	dbLock sync.RWMutex // Usado para proteger o acesso concorrente ao mapa.
}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{
		db: make(map[string]*reservation.Reservation),
	}
}

func (repo *ReservationRepository) Save(res *reservation.Reservation) error {
	repo.dbLock.Lock()
	defer repo.dbLock.Unlock()

	// Simula a operação de salvamento no banco de dados
	if _, exists := repo.db[res.ID]; exists {
		return errors.New("reservation with the same ID already exists")
	}

	repo.db[res.ID] = res
	return nil
}

func (repo *ReservationRepository) GetAllReservations() ([]*reservation.Reservation, error) {
	repo.dbLock.RLock()
	defer repo.dbLock.RUnlock()

	// Simula a operação de busca no banco de dados
	reservations := []*reservation.Reservation{}
	for _, res := range repo.db {
		reservations = append(reservations, res)
	}

	return reservations, nil
}

func (repo *ReservationRepository) GetReservationByID(id string) (*reservation.Reservation, error) {
	repo.dbLock.RLock()
	defer repo.dbLock.RUnlock()

	// Simula a operação de busca no banco de dados
	if res, exists := repo.db[id]; exists {
		return res, nil
	}

	return nil, errors.New("reservation not found")
}
