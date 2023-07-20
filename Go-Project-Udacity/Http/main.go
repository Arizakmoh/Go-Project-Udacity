package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Customer struct
type Customer struct {
	ID      int
	Name    string
	Email   string
	Address string
}

// in memory database
var customers []Customer

// Handler functions

// Get All Customer Handler
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// Get Specific Customer Handler
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, c := range customers {
		if c.ID == id {
			json.NewEncoder(w).Encode(c)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func main() {
	// Seed with some initial data
	customers = append(customers, Customer{
		ID:    1,
		Name:  "ABDIRIZAK ABDULLAHI",
		Email: "abdirizak@example.com",
	})
	customers = append(customers, Customer{
		ID:    2,
		Name:  "MOHAMED WARSAME",
		Email: "moha@example.com",
	})

	for _, cust := range customers {
		fmt.Println(cust)
	}
}
