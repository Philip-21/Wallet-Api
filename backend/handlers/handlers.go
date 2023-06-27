package handlers

import (
	"wallet-api/config"

	"github.com/gin-gonic/gin"
)
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(c *gin.Context) {
	c.File("./frontend/src/components/home.js")
}

func (m *Repository) Signup(c *gin.Context) {



}

func (m *Repository) Login(c *gin.Context) {}



