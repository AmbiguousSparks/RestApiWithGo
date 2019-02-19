package main

import (
	c "ApiUsers/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user", c.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", c.GetUser).Methods("GET")
	router.HandleFunc("/user", c.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", c.UpdateUser).Methods("POST")
	router.HandleFunc("/user/{id}", c.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", options).Methods("OPTIONS")
	log.Println("Running")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func options(w http.ResponseWriter, r *http.Request) {
	/*função de apoio, pois as vezes requests como delete, na vdd são enviadas
	primeiro como options, nesses casos o handle não irá fazer nada com o options
	mas com o delete ele funciona corretamente
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,DELETE,PUT")
}
