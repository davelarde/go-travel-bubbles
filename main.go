package main

import(
	_ "fmt"
	"log"
	"encoding/json"
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

// init laundries var as alice Laundry struct
var laundries []Laundry
func laundryPage(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Happy Bubbles!")
}
// get all laundries
func getLaundries(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(laundries)

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
// mock data
laundries = append(laundries, Laundry{ID: "1", Laundrycode:"414141", Name: "Uphill Laundry", Location: &Location{Country: "USA", City:"San Francisco"}})

laundries = append(laundries, Laundry{ID: "2", Laundrycode:"000001", Name: "Laundry Nowhere", Location: &Location{Country: "Chile", City:"Arica"}})

laundries = append(laundries, Laundry{ID: "3", Laundrycode:"123456", Name: "Lavanderia Castillo", Location: &Location{Country: "Spain", City:"Madrid"}})

laundries = append(laundries, Laundry{ID: "4", Laundrycode:"245678", Name: "Ciao Laundry!", Location: &Location{Country: "Italy", City:"Aosta"}})

laundries = append(laundries, Laundry{ID: "5", Laundrycode:"908756", Name: "Aussie Mat!", Location: &Location{Country: "Australia", City:"Melbourne"}})

laundries = append(laundries, Laundry{ID: "6", Laundrycode:"345677", Name: "Vetements propes", Location: &Location{Country: "France", City:"Bordeaux"}})

laundries = append(laundries, Laundry{ID: "7", Laundrycode:"745678", Name: "kabarciklar", Location: &Location{Country: "Turkey", City:"Istanbul"}})

laundries = append(laundries, Laundry{ID: "8", Laundrycode:"345678", Name: "Easy laundry", Location: &Location{Country: "USA", City:"Philadelphia"}})

laundries = append(laundries, Laundry{ID: "9", Laundrycode:"202204", Name: "Lavanderia Rapida", Location: &Location{Country: "Brazil", City:"Salvador do Bahia"}})

laundries = append(laundries, Laundry{ID: "10", Laundrycode:"444444", Name: "Aguayo", Location: &Location{Country: "Bolivia", City:"La Paz"}})

	// Route handlers /end points
	r.HandleFunc("/", laundryPage)
	r.HandleFunc("/laundries", getLaundries).Methods("GET")
	r.HandleFunc("/laundries/{id}", getLaundry).Methods("GET")
	r.HandleFunc("/laundries", createLaundry).Methods("POST")
	r.HandleFunc("/laundries/{id}", updateLaundry).Methods("PUT")
	r.HandleFunc("/laundries/{id}", deleteLaundry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}