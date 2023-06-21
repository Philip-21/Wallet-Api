package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *User) InsertUser(entry User) error {
	collection := client.Database("user").Collection("user")

	_, err := collection.InsertOne(context.TODO(), User{
		Name:      entry.Name,
		Email:     entry.Email,
		Password:  entry.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting User into Collection", err)
		return err
	}
	log.Println("User inserted")
	return nil
}

func (u *User) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("user").Collection("user")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry User
	//extracting based on the id 
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}
	return &entry , nil
}
