package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contato struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Owner string `json:"owner"`
}

var Contacts *mongo.Collection

func FindContact(name string) []bson.M {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filterCursor, err := Contacts.Find(ctx, bson.M{"name": name})
	if err != nil {
			log.Fatal(err)
	}

	var contactsFiltered []bson.M
	if err = filterCursor.All(ctx, &contactsFiltered); err != nil {
			log.Fatal(err)
	}
	return contactsFiltered
}