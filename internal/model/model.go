package model

type Model interface{}

// Response struct
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Model
}
