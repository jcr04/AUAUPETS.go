package main

import (
	"log"

	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/api"
	"github.com/jcr04/AUAUPETS.go/internal/infrastructure/database"
)

func main() {
	log.Println("Iniciando o servidor...") // Log de início

	if err := database.Connect(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	api.StartServer() // Suponho que isso bloqueie, se não, você vai querer esperar por isso de alguma forma.

	log.Println("Servidor encerrado.") // Log de encerramento
}
