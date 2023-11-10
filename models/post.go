package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"body"`
	Author User   `json:"userId"`
	// Likes []Like `json:"likes"`
}

func (post *Post) Create() error {
	connStr := "user=postgres password=isaicoloma2 dbname=microS sslmode=disable host=localhost port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS posts (id SERIAL PRIMARY KEY, title TEXT, content TEXT, author TEXT, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
