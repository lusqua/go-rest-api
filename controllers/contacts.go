package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lusqua/go-rest-api/database"
	"github.com/lusqua/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
)


func GetContacts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := database.Contacts.Find(ctx, bson.M{})
	
	if err != nil {
			log.Fatal(err)
	}
	var contacts_list []bson.M
	if err = cursor.All(ctx, &contacts_list); err != nil {
			log.Fatal(err)
	}
	
	json.NewEncoder(w).Encode(contacts_list)
}


func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	contacts := models.FindContact(params["name"])

	if len(contacts) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}


func CreateContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contato

	_ = json.NewDecoder(r.Body).Decode(&newContact)

	contact := models.FindContact(newContact.Name)

	if len(contact) > 1 {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("Já existe um contato com este nome")
		return
	}
	
	insertResult, err := database.Contacts.InsertOne(context.TODO(), newContact)
	if err != nil {
			log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertResult)
}


func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	contact := models.FindContact(params["name"])

	if len(contact) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(contact) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Há mais de um contato com este nome")

		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()


	res, err := database.Contacts.DeleteOne(ctx, bson.M{"_id": contact[0]["_id"]})

	if err != nil {
	log.Fatal("DeleteOne() ERROR:", err)
	}

	json.NewEncoder(w).Encode(res)
}