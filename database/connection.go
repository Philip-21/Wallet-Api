package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectToDB() (*mongo.Client, error) {
	//mongoURL := os.Getenv("MONGO_URL")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURL := os.Getenv("MONGO_URL")
	clientoptions := options.Client().ApplyURI(mongoURL)

	conn, err := mongo.Connect(ctx, clientoptions)
	if err != nil {
		log.Println("Error in connecting....", err)
	}
	log.Println("connected to mongo")
	//context to disconnect mongo
	ct, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection when the function exits due to an error
	defer func() {
		if err = conn.Disconnect(ct); err != nil {
			panic(err)
		}
	}()
	return conn, nil
}
