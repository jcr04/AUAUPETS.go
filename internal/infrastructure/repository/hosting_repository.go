// C:\Users\Usuario\go\AUAUPETS.go\internal\infrastructure\repository\hosting_repository.go

package repository

import (
	"strconv"

	"github.com/jcr04/AUAUPETS.go/internal/domain/hosting"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/database"
)

type HostingRepository struct{}

func NewHostingRepository() *HostingRepository {
	return &HostingRepository{}
}

func (repo *HostingRepository) Save(h *hosting.Hosting) error {
	db := database.GetDB() // Obtém a conexão com o banco de dados

	query := `
        INSERT INTO hostings (name, reservation_id)
        VALUES ($1, $2)
        RETURNING id;
    `

	var id int
	err := db.QueryRow(query, h.Name, h.Reservation.ID).Scan(&id)
	if err != nil {
		return err
	}

	h.ID = strconv.Itoa(id) // Atualiza o ID da hospedagem com o ID retornado pelo banco de dados
	return nil
}

func (repo *HostingRepository) GetAllHostings() ([]*hosting.Hosting, error) {
	db := database.GetDB() // Obtém a conexão com o banco de dados

	rows, err := db.Query("SELECT * FROM hostings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hostings []*hosting.Hosting
	for rows.Next() {
		var h hosting.Hosting
		err := rows.Scan(&h.ID, &h.Name, &h.ReservationID) // Agora há três argumentos para Scan
		if err != nil {
			return nil, err
		}
		hostings = append(hostings, &h)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hostings, nil
}
