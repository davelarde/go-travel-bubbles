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

// laundry struct(models)
 type Laundry struct{
	ID string `json:id`
	Laundrycode string `json:laundrycode`
	Name string `json:name`
	Location *Location `json:author`
 }
//  location struct
type Location struct{
	Country string `json:country`
	City string `json:city`
}
func main(){
	// initializing mux router
	r := mux.NewRouter() 

	// Route handlers /end points
	r.HandleFunc("/api/laundries", getLaundries).Methods("GET")
	r.HandleFunc("/api/laundries/{id}", getLaundry).Methods("GET")
	r.HandleFunc("/api/laundries", createLaundry).Methods("POST")
	r.HandleFunc("/api/laundries/{id}", updateLaundry).Methods("PUT")
	r.HandleFunc("/api/laundries/{id}", deleteLaundry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}