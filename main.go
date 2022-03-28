package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lusqua/go-rest-api/database"
	"github.com/lusqua/go-rest-api/routes"
)

func init() {
	err := godotenv.Load("database.env")
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("Iniciando API")
	database.Client = database.Db_connect()

	database.Contacts = database.Client.Database("my_agenda").Collection("contacts")

	routes.HandleRequest()
}
