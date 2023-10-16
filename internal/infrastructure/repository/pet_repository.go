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
	err := repo.db.QueryRow(
		"SELECT ID, Name, Breed, Age, Weight, Health, Alergic FROM Pets WHERE ID = $1", id,
	).Scan(&pet.ID, &pet.Name, &pet.Breed, &pet.Age, &pet.Weight, &pet.Health, &pet.Alergic)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (repo *PetRepository) CreateAnimal(a *animal.Animal) error {
	query := `
        INSERT INTO Pets (Name, Breed, Age, Weight, Health, Alergic)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING ID;
    `
	return repo.db.QueryRow(query, a.Name, a.Breed, a.Age, a.Weight, a.Health, a.Alergic).Scan(&a.ID)
}

func (repo *PetRepository) ListAnimals() ([]*animal.Animal, error) {
	rows, err := repo.db.Query("SELECT * FROM Pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []*animal.Animal
	for rows.Next() {
		var a animal.Animal
		err := rows.Scan(&a.ID, &a.Name, &a.Breed, &a.Age, &a.Weight, &a.Health, &a.Alergic)
		if err != nil {
			return nil, err
		}
		animals = append(animals, &a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return animals, nil
}

func (repo *PetRepository) UpdateAnimal(id string, a *animal.Animal) error {
	query := `
		UPDATE Pets
		SET Name = $2, Breed = $3, Age = $4, Weight = $5, Health = $6, Alergic = $7
		WHERE ID = $1;
	`
	_, err := repo.db.Exec(query, id, a.Name, a.Breed, a.Age, a.Weight, a.Health, a.Alergic)
	return err
}

func (repo *PetRepository) DeleteAnimal(id string) error {
	query := `
		DELETE FROM Pets
		WHERE ID = $1;
	`
	_, err := repo.db.Exec(query, id)
	return err
}
