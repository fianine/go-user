package api

import (
	"encoding/json"
	"go-user/internal/database/config"
	"go-user/internal/model"
	"log"
	"net/http"
)

// Get Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	var getUsers []model.Users
	var response model.Response

	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	rows, err := connection.SQL.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			getUsers = append(getUsers, users)
		}
	}

	response.Status = true
	response.Message = "Success"
	response.Data = getUsers

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Add New User
func AddUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseMultipartForm(4096)
	if parseErr != nil {
		panic(parseErr)
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	_, err = connection.SQL.Exec("INSERT INTO users (first_name, last_name) values (?,?)",
		firstName,
		lastName,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = true
	response.Message = "Successfully added"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseMultipartForm(4096)
	if parseErr != nil {
		panic(parseErr)
	}

	id := r.FormValue("user_id")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	_, err = connection.SQL.Exec("UPDATE users SET first_name = ?, last_name = ? WHERE id = ?",
		firstName,
		lastName,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = true
	response.Message = "Successfully Updated"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	parseErr := r.ParseMultipartForm(4096)
	if parseErr != nil {
		panic(parseErr)
	}

	id := r.FormValue("user_id")

	_, err = connection.SQL.Exec("DELETE from users WHERE id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = true
	response.Message = "Successufully Deleted"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
