package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Connect inicializa a conexão com o MongoDB e define a variável global DB.
// Retorna um erro se a conexão falhar.
func Connect(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao MongoDB: %w", err)
	}

	// Valida a conexão com um ping
	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("erro ao pingar o MongoDB: %w", err)
	}

	log.Println("Conexão com MongoDB estabelecida com sucesso!")
	return nil
}
