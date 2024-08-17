package dbCon

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var db *sql.DB

// GetDB method returns a DB instance
func GetDB() (*sql.DB, error) {

	if db == nil {
		// Retrieve environment variables
		dbName := os.Getenv("DB_NAME")
		dbPassword := os.Getenv("DB_PASSWORD")

		connection, err := sql.Open("postgres", "dbname="+dbName+" host=localhost user=postgres password="+dbPassword+" sslmode=disable")
		if err != nil {
			panic(fmt.Sprintf("DB: %v", err))
		}
		db = connection
	}

	return db, nil
}
