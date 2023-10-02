package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type Payment struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	User        *User     `bson:"user,omitempty"`
	Amount      float64   `bson:"amount" json:"amount"`
	Description string    `bson:"description" json:"description"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type Models struct {
	PaymentModel Payment
	UserModel    User
}
//New() is used to initialize the database models and establish the connection to the MongoDB database
func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		UserModel:    User{},
		PaymentModel: Payment{},
	}
}
