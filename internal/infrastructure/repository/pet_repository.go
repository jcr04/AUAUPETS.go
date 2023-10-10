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

func (repo *PetRepository) CreateAnimal(a *animal.Animal) error {
	query := `
        INSERT INTO Pets (Name, Breed, Age, CheckIn, CheckOut)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING ID;
    `
	return repo.db.QueryRow(query, a.Name, a.Breed, a.Age, a.CheckIn, a.CheckOut).Scan(&a.ID)
}
