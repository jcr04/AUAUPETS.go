package repository

import (
	"database/sql"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/database"
)

type PetRepository struct {
	db *sql.DB
}

func NewPetRepository() *PetRepository {
	return &PetRepository{db: database.GetDB()}
}

func (repo *PetRepository) GetPetByID(id string) (*animal.Animal, error) {
	var pet animal.Animal
	err := repo.db.QueryRow("SELECT * FROM Pets WHERE ID = $1", id).Scan(&pet.ID, &pet.Name, &pet.Breed, &pet.Age)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}
