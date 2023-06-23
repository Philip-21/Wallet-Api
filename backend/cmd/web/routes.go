package main

import (
	"wallet-api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes (app*handlers.Repository)*gin.Engine{

	router := gin.Default()

	return router 

}