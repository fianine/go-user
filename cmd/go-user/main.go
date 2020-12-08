package main

import (
	"fmt"
	"go-user/internal/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Router
	router := router.Router()

	fmt.Println("Server is running...")

	log.Fatal(http.ListenAndServe(":3001", router))
}
