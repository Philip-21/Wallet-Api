package main

import (
	"wallet-api/handlers"
	"wallet-api/helpers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Routes(app *handlers.Repository) *gin.Engine {

	router := gin.Default()

	//register websocket handlers 
	router.GET("/ws",helpers.HandleWebsocket)
	//start the broadcast routine
	go helpers.HandleBroadcast()

	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https//*", "http://*"},//communicatng with port 3000 enabled
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Access-Control-Request-Method", "Access-Control-Request-Headers", "Accept", "Authorization", " Accept-Encoding",
			"Content-Type", "Connection", " Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header", "Token", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link", "Content-Length"},
		AllowCredentials: true, //AllowCredentials indicates whether the request can include user credentials like cookies

		MaxAge: 300, //MaxAge indicates how long (with second-precision) the results of a preflight request can be cached
	}))

	return router

}
