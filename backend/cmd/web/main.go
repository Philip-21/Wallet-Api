package main

import (
	"log"
	"os"
	"wallet-api/config"
	"wallet-api/database"
	"wallet-api/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)


const portNumber = ":8080"

var client *mongo.Client

func main() {
	//setting up a logger to write to the terminal,helps in writing the client and server errors
	//info log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//error log
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
      //app configurations
	app := config.AppConfig{
		Infolog: infoLog,
		Errorlog: errorLog,
		Models: database.New(client),
	}
	//db conections
	mongoClient, err := database.ConnectToDB()
	if err != nil {
		app.Errorlog.Panic("Error in connecting to db..... ", err)
	}
	client = mongoClient
	//app.Models = database.New(client)
	
	app.Infolog.Println("connected to mongo successfully ")

	//running the application
	app.Infolog.Printf("Starting application on port %s", portNumber)
	r := Routes(handlers.Repo)
	err = r.Run(portNumber)
	if err != nil {
		app.Errorlog.Fatal("Cant run app :", err)
	}

}
