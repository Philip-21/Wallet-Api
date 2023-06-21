package main

import (
	"log"
	"os"
	"wallet-api/config"
	"wallet-api/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var app *config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

var client *mongo.Client

func main() {
	//setting up a logger to write to the terminal,helps in writing the client and server errors
	//info log
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//error log
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.Errorlog = errorLog
	app.Infolog = infoLog

	//db conections
	mongoClient, err := database.ConnectToDB()
	if err != nil {
		app.Errorlog.Panic("Error in connecting to db..... ", err)
	}
	client = mongoClient
	app.Models = database.New(client)
	// app := config.AppConfig{
	// 	Models: database.New(client),
	// }
	app.Infolog.Println("connected to mongo successfully ")
}
