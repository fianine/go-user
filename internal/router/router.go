package router

import (
	"go-user/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

// Router
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", api.GetUsers).Methods("GET")            // Get Users
	router.HandleFunc("/add_user", api.AddUser).Methods("POST")         // Add New User
	router.HandleFunc("/update_user", api.UpdateUser).Methods("PUT")    // Update User
	router.HandleFunc("/delete_user", api.DeleteUser).Methods("DELETE") // Delete User

	http.Handle("/", router)

	return router
}
