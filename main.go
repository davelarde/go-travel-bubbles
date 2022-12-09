package main

import(
	_ "fmt"
	"log"
	"encoding/json"
	"net/http"
	 "math/rand"
	 "strconv" 
	"github.com/gorilla/mux"
	"html/template"
	
)  
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
// laundry struct(models)
 type Laundry struct{
	ID string `json:id`
	Laundrycode string `json:laundrycode`
	Name string `json:name`
	Country string `json:country`
	City string `json:city`
	// Location *Location `json:location`
 }
//  location struct
// type Location struct{
// }

// init laundries var as alice Laundry struct
var laundries []Laundry
func laundryPage(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "index.html", nil)
}

func formPage(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "form.html", nil)
}

func aboutPage(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "about.html", nil)
}
// get all laundries
func getLaundries(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(laundries)

}
// get single laundry

func getLaundry(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)  //get the params
	//loop through all laundries and find it by id
	for _, item := range laundries {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
json.NewEncoder(w).Encode(&Laundry{})
}

func createLaundry(w http.ResponseWriter, r *http.Request){
	
	 if r.Method != "POST"{
			tpl.Execute(w, nil)
			return 
		}
		details := Laundry{
	    Laundrycode : r.PostFormValue("laundryCode"),
		Name : r.PostFormValue("laundryName"),
		Country: r.PostFormValue("country"),
		City: r.PostFormValue("city"),
		

	
	}
		// fmt.Fprintf(w, "Name = %s\n",r.PostFormValue("laundryCode"))


		// *Location : r.PostFormValue("country"),
		// Location :r.PostFormValue("city"),
	

		tpl.ExecuteTemplate(w, "welcome.html", details)
		w.Header().Set("Content-Type", "application.json")
		var laundry Laundry
		_ = json.NewDecoder(r.Body).Decode(&laundry)
		laundry.ID = strconv.Itoa(rand.Intn(10000)) //mock id, not really used in production, use it only for test 
		laundry.Name = details.Name
		laundry.Laundrycode = details.Laundrycode
		laundry.Country = details.Country
		laundry.City = details.City

		laundries = append(laundries, laundry)
		json.NewEncoder(w).Encode(laundry)
		
	}
	
		
	

// Update a laundry
func updateLaundry(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)
	for index, item := range laundries{   //this part comes from deleting a laudry 
		if item.ID == params["id"]{
			laundries = append(laundries[:index], laundries[index+1:]...)
			var laundry Laundry
	_ = json.NewDecoder(r.Body).Decode(&laundry)
	laundry.ID = params["id"] //this part comes from create a new laundry
	laundries = append(laundries, laundry)
	json.NewEncoder(w).Encode(laundry)
	return
	}
}
json.NewEncoder(w).Encode(laundries)

}
// delete a laundry
func deleteLaundry(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)
	for index, item := range laundries{
		if item.ID == params["id"]{
		laundries = append(laundries[:index], laundries[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(laundries)

}

func main(){
	// initializing mux router
	r := mux.NewRouter() 
// mock data
laundries = append(laundries, Laundry{ID: "1", Laundrycode:"414141", Name: "Uphill Laundry",Country: "USA", City:"San Francisco"})

laundries = append(laundries, Laundry{ID: "2", Laundrycode:"000001", Name: "Laundry Nowhere",Country: "Chile", City:"Arica"})

laundries = append(laundries, Laundry{ID: "3", Laundrycode:"123456", Name: "Lavanderia Castillo", Country: "Spain", City:"Madrid"})

laundries = append(laundries, Laundry{ID: "4", Laundrycode:"245678", Name: "Ciao Laundry!", Country: "Italy", City:"Aosta"})

laundries = append(laundries, Laundry{ID: "5", Laundrycode:"908756", Name: "Aussie Mat!", Country: "Australia", City:"Melbourne"})

laundries = append(laundries, Laundry{ID: "6", Laundrycode:"345677", Name: "Vetements propes", Country: "France", City:"Bordeaux"})

laundries = append(laundries, Laundry{ID: "7", Laundrycode:"745678", Name: "kabarciklar", Country: "Turkey", City:"Istanbul"})

laundries = append(laundries, Laundry{ID: "8", Laundrycode:"345678", Name: "Easy laundry", Country: "USA", City:"Philadelphia"})

laundries = append(laundries, Laundry{ID: "9", Laundrycode:"202204", Name: "Lavanderia Rapida", Country: "Brazil", City:"Salvador do Bahia"})

laundries = append(laundries, Laundry{ID: "10", Laundrycode:"444444", Name: "Aguayo", Country: "Bolivia", City:"La Paz"})

	// Route handlers /end points
    // fs:= http.FileServer(http.Dir("./assets/"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", laundryPage)
	r.HandleFunc("/form", formPage)
	r.HandleFunc("/about", aboutPage)
	// r.HandleFunc("/welcome", formHandler)
	r.HandleFunc("/laundries", getLaundries).Methods("GET")
	r.HandleFunc("/laundries/{id}", getLaundry).Methods("GET")
	r.HandleFunc("/welcome", createLaundry).Methods("POST")
	r.HandleFunc("/laundries/{id}", updateLaundry).Methods("PUT")
	r.HandleFunc("/laundries/{id}", deleteLaundry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}