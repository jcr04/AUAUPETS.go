package main

import (
	"log"

	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/api"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/database"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	api.StartServer()
}
