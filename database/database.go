package database

import (
	"os"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect with database
func Connect() (mongo.Client, context.Context, string, context.CancelFunc) {
	MONGO_CONNECTION_STRING := os.Getenv("MONGO_CONNECTION_STRING")
	DATABASE_NAME := "upfastmain"

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_CONNECTION_STRING))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	fmt.Println("Connected to Database")

	return *client, ctx, DATABASE_NAME, cancel
}