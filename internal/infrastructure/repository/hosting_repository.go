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
