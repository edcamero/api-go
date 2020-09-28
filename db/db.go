package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:mongo13@mascota.sdg9g.mongodb.net/mascota?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conecto")

	return client
}



