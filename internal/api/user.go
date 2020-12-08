package api

import (
	"go-user/internal/database/config"
	"go-user/internal/model"
	"go-user/internal/service"
	"log"
	"net/http"
)

// Get Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	connection, err := config.ConnectSQL()
	defer connection.SQL.Close()

	var users []model.Model
	rows, err := connection.SQL.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.FirstName, &user.LastName)

		users = append(users, user)
	}

	response := model.Response{
		Status:  200,
		Message: "Success",
		Data:    users,
	}

	responseWithJson(w, response)
}

// Add New User
func AddUser(w http.ResponseWriter, r *http.Request) {
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

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}

// Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	response := model.Response{
		Status:  201,
		Message: "Success",
	}

	responseWithJson(w, response)
}

// User with address
func UserAddress(w http.ResponseWriter, r *http.Request) {
	connection, _ := config.ConnectSQL()
	defer connection.SQL.Close()

	userID := r.URL.Query().Get("userID")

	row := connection.SQL.QueryRow("SELECT * FROM users WHERE id = ?", userID)

	var user model.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		responseWithJson(w, model.Response{
			Status:  404,
			Message: "Not found",
			Data:    []model.Model{},
		})
		return
	}

	user.Addresses, _ = service.GetAddressUser(user.ID)

	responseWithJson(w, model.Response{
		Status:  200,
		Message: "ok",
		Data:    []model.Model{user},
	})
}
