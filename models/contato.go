package models

import "go.mongodb.org/mongo-driver/mongo"

type Contato struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var Contacts *mongo.Collection