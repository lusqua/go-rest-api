package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/lusqua/go-rest-api/database"
	"go.mongodb.org/mongo-driver/bson"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := database.Users.Find(ctx, bson.M{})
	
	if err != nil {
			log.Fatal(err)
	}
	var contacts_list []bson.M
	if err = cursor.All(ctx, &contacts_list); err != nil {
			log.Fatal(err)
	}
	
}
