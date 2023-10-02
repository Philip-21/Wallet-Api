package handlers

import (
	"net/http"
	"time"
	"wallet-api/config"
	"wallet-api/database"

	"github.com/gin-gonic/gin"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}
type userpayload struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type loginpayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *Repository) Home(c *gin.Context) {
	c.JSON(http.StatusOK, "wallet system up and running")
}

func (m *Repository) Signup(c *gin.Context) {
	var requestPayload userpayload
	err := c.ShouldBindJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in binding json": err.Error()})
		return
	}
	event := database.User{
		Name:      requestPayload.Name,
		Email:     requestPayload.Email,
		Password:  requestPayload.Password,
		CreatedAt: requestPayload.CreatedAt,
		UpdatedAt: requestPayload.UpdatedAt,
	}
	err = m.App.Models.UserModel.InsertUser(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error inserting to database": err.Error()})
	}
	c.JSON(http.StatusAccepted, "user created")
	//c.Redirect(http.StatusSeeOther, "/user/login")
}

func (m *Repository) Login(c *gin.Context) {
	var login loginpayload
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in binding json": err.Error()})
		return
	}
	event := database.User{
		Name:     login.Email,
		Password: login.Password,
	}
	_, err = m.App.Models.UserModel.AuthUser(event.Name, event.Password)
	if err != nil {
		c.JSON((http.StatusInternalServerError), gin.H{"invalid details": err.Error()})
	}
	//jwt implementation

}
