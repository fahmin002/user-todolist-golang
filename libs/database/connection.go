package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//connStr := "user=postgres dbname=postgres password=mhfu123 sslmode=verify-full"
	connStr := "host=" + os.Getenv("DBHOST") +
		" user=" + os.Getenv("DBUSER") +
		" dbname=" + os.Getenv("DBNAME") +
		" password=" + os.Getenv("DBPASS") +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
