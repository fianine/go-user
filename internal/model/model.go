package model

// Users struct
type Users struct {
	ID        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

// Response struct
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
