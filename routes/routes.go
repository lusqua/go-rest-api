package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lusqua/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/contato", controllers.GetContacts).Methods("GET")
	r.HandleFunc("/api/contato/{name}", controllers.GetContact).Methods("GET")
	r.HandleFunc("/api/contato", controllers.CreateContact).Methods("POST")
	r.HandleFunc("/api/contato/{name}", controllers.DeleteContact).Methods("DELETE")

	log.Println("API Inicializada")
	log.Fatal(http.ListenAndServe(":8000", r))
}