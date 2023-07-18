package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// This struct represents the JSON structure of our Menu to be returned
// it also represented the model in the Database
type Menu struct {
	Id    string
	Name  string
	Price string
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	//GET
	r.HandleFunc("/menu", returnFullMenu)
	r.HandleFunc("/menu/{id}", returnItem)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func returnFullMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnFullMenu")

	db, err := gorm.Open(sqlite.Open("menu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var menu []Menu

	db.Find(&menu)
	fmt.Println(menu)

	json.NewEncoder(w).Encode(menu)
}

func returnItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	db, err := gorm.Open(sqlite.Open("menu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var menu []Menu

	db.Where("id = ?", key).Find(&menu)
	json.NewEncoder(w).Encode(menu)
}


func main() {
	fmt.Println("Webapp Rest API v1.0")

	// Add the call to our new initialMigration function
	initialMigration()

	handleRequests()
}

// our initial migration function
func initialMigration() {
	db, err := gorm.Open(sqlite.Open("menu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Menu{})

	//Initial population
	db.Create(&Menu{Id: "1", Name: "Pain au Chocolat", Price: "£1.99"})
	db.Create(&Menu{Id: "2", Name: "Croque monsieur", Price: "1.00"})
	db.Create(&Menu{Id: "3", Name: "Box of Macarons (6)", Price: "7.00"})
	db.Create(&Menu{Id: "4", Name: "Mille Feuille", Price: "1.50"})
}
