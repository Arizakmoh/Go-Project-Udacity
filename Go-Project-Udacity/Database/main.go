package main

// Customer struct
type Customer struct {
	ID      int
	Name    string
	Email   string
	Address string
}

var customers []Customer

func main() {
	// Seed with some initial data
	customers = append(customers, Customer{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	})
}
