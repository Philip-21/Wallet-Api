package config

import (
	"log"
	"wallet-api/database"
)

type AppConfig struct {
	Infolog  *log.Logger
	Errorlog *log.Logger
	Models database.Models
}
