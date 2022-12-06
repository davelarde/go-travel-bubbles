package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv" 
	"github.com/gorilla/mux"
)  

func main(){
	// initializing mux router
	r := mux.NewRouter() 

	// Route handlers /end points
	r.HandleFunc("/api/laundries", getLaundries).Methods("GET")
	r.HandleFunc("/api/laundries/{id}", getLaundry).Methods("GET")
	r.HandleFunc("/api/laundries", createLaundry).Methods("POST")
	r.HandleFunc("/api/laundries/{id}", updateLaundry).Methods("PUT")
	r.HandleFunc("/api/laundries/{id}", deleteLaundry).Methods("DELETE")
}