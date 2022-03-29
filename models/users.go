package models

import (
	"context"
	"log"
	"time"

	"github.com/lusqua/go-rest-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


var Users *mongo.Collection

func FindUser(username string) []bson.M {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filterCursor, err := database.Users.Find(ctx, bson.M{"Username": username})
	if err != nil {
			log.Fatal(err)
	}

	var usersFiltered []bson.M
	if err = filterCursor.All(ctx, &usersFiltered); err != nil {
			log.Fatal(err)
	}

	return usersFiltered
}