package main

// Create a Customer struct
type Customer struct {
	ID    int
	Name  string
	Email string
}

// Mock in-memory database
var customers []*Customer

// Init some seed data
func init() {
	customers = append(customers, &Customer{
		ID:    1,
		Name:  "Abdirizak Abdullahi",
		Email: "abdirizak@noemail.com"})

	customers = append(customers, &Customer{
		ID:    2,
		Name:  "Nimco warsame",
		Email: "nimco@noemail.com"})
}
