package model

// Users struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Addresses []Address `json:"address"`
}

type Address struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
}
