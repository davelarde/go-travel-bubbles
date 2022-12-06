package main

import(
	_ "fmt"
	"log"
	_ "encoding/json"
	"net/http"
	_ "math/rand"
	_ "strconv" 
	"github.com/gorilla/mux"
	"io"
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
func laundryPage(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Happy Bubbles!")
}
// get all laundries
func getLaundries(w http.ResponseWriter, r *http.Request){


}

func getLaundry(w http.ResponseWriter, r *http.Request){
	

}

func createLaundry(w http.ResponseWriter, r *http.Request){
	

}

func updateLaundry(w http.ResponseWriter, r *http.Request){
	

}

func deleteLaundry(w http.ResponseWriter, r *http.Request){
	

}

func main(){
	// initializing mux router
	r := mux.NewRouter() 

	// Route handlers /end points
	r.HandleFunc("/", laundryPage)
	r.HandleFunc("/laundries", getLaundries).Methods("GET")
	r.HandleFunc("/laundries/{id}", getLaundry).Methods("GET")
	r.HandleFunc("/laundries", createLaundry).Methods("POST")
	r.HandleFunc("/laundries/{id}", updateLaundry).Methods("PUT")
	r.HandleFunc("/laundries/{id}", deleteLaundry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}