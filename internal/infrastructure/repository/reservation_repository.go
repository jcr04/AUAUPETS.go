package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jcr04/AUAUPETS.go/internal/domain/animal"
	"github.com/jcr04/AUAUPETS.go/internal/domain/reservation"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/database"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository() *ReservationRepository {
	db := database.GetDB()
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	return &ReservationRepository{
		db: db,
	}
}

func (repo *ReservationRepository) Save(res *reservation.Reservation) error {
	query := `
        INSERT INTO reservations (id, animal_id, check_in, check_out)
        VALUES ($1, $2, $3, $4);
    `
	_, err := repo.db.Exec(query, res.ID, res.Animal.ID, res.CheckIn, res.CheckOut)
	if err != nil {
		log.Printf("Error saving reservation: %v", err)
		return fmt.Errorf("error saving reservation: %w", err)
	}
	return nil
}

func (repo *ReservationRepository) GetAllReservations() ([]*reservation.Reservation, error) {
	rows, err := repo.db.Query(`
        SELECT 
            reservations.id, 
            Pets.id, 
            Pets.name, 
            Pets.breed, 
            Pets.age, 
            reservations.check_in, 
            reservations.check_out
        FROM 
            reservations
        JOIN 
            Pets ON reservations.animal_id = Pets.id
    `)
	if err != nil {
		log.Printf("Error querying all reservations: %v", err)
		return nil, fmt.Errorf("error querying all reservations: %w", err)
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var res reservation.Reservation
		res.Animal = &animal.Animal{} // Inicialize res.Animal antes da chamada Scan
		err := rows.Scan(
			&res.ID,
			&res.Animal.ID,
			&res.Animal.Name,
			&res.Animal.Breed,
			&res.Animal.Age,
			&res.CheckIn,
			&res.CheckOut,
		)
		if err != nil {
			log.Printf("Error scanning reservation row: %v", err)
			return nil, fmt.Errorf("error scanning reservation row: %w", err)
		}
		reservations = append(reservations, &res)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error after iterating through reservation rows: %v", err)
		return nil, fmt.Errorf("error after iterating through reservation rows: %w", err)
	}

	return reservations, nil
}

func (repo *ReservationRepository) GetReservationByID(id string) (*reservation.Reservation, error) {
	log.Printf("GetReservationByID called with id: %v", id) // Log the id value
	log.Printf("Database connection object: %+v", repo.db)  // Log the database connection object

	if repo.db == nil {
		log.Println("Database connection is nil in GetReservationByID")
		return nil, fmt.Errorf("database connection is nil")
	}

	var res reservation.Reservation
	res.Animal = &animal.Animal{} // Inicialize res.Animal antes da chamada Scan
	query := `
        SELECT 
            reservations.id, 
            Pets.id, 
            Pets.name, 
            Pets.breed, 
            Pets.age, 
            reservations.check_in, 
            reservations.check_out
        FROM 
            reservations
        JOIN 
            Pets ON reservations.animal_id = Pets.id
        WHERE 
            reservations.id = $1;
    `
	err := repo.db.QueryRow(query, id).Scan(
		&res.ID,
		&res.Animal.ID,
		&res.Animal.Name,
		&res.Animal.Breed,
		&res.Animal.Age,
		&res.CheckIn,
		&res.CheckOut,
	)

	if err != nil {
		log.Printf("Error querying reservation by ID: %v", err)
		return nil, fmt.Errorf("error querying reservation by ID: %w", err)
	}
	return &res, nil
}
