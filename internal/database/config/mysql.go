package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DB Struct
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// Connect to database
func ConnectSQL() (*DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Load env failed")
	}

	var env = fmt.Sprintf("%v:%v@/%v?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"))

	db, err := sql.Open("mysql", env)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, err
}
