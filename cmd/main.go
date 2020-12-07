package main

import (
	"fmt"
	"go-user/internal/api"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", api.GetUsers).Methods("GET")            // Get Users
	router.HandleFunc("/add_user", api.AddUser).Methods("POST")         // Add New User
	router.HandleFunc("/update_user", api.UpdateUser).Methods("PUT")    // Update User
	router.HandleFunc("/delete_user", api.DeleteUser).Methods("DELETE") // Delete User

	http.Handle("/", router)

	fmt.Println("Connected to port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))
}
