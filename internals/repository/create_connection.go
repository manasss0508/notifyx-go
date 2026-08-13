package repository

import (
	"database/sql"
	"os"

	"github.com/manasss0508/notifyx-go/internals/repository/queries"
)

func CreateConnection() *queries.Queries {
	// database connection url
	dbUrl := os.Getenv("DATABASE_URL")

	// making connection to database
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	// creating and returning queries instance
	return queries.New(conn)
}
