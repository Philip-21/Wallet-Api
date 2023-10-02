package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUrl = "mongodb://mongo:27017"
)

var ctx = context.TODO()

func ConnectToDB() (*mongo.Client, error) {
	clientoptions := options.Client().ApplyURI(mongoUrl)

	conn, err := mongo.Connect(ctx, clientoptions)
	if err != nil {
		log.Println("Error in connectig....", err)
	}
	log.Println("connected to mongo")
	//context to disconnect mongo
	ct, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection
	defer func() {
		if err = conn.Disconnect(ct); err != nil {
			panic(err)
		}
	}()
	return conn, nil
}
