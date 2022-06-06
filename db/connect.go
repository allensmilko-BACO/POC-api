package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dataBase := os.Getenv("DB_URI")
	connectUri := fmt.Sprintf("mongodb+srv://%s:%s@%s", user, pass, dataBase)
	clientOpts := options.Client().ApplyURI(connectUri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(client)
	fmt.Println("Congratulations, you're already connected to MongoDB!")

	return client
}