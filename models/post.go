package models

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"body"`
	Author User   `json:"userId"`
	// Likes []Like `json:"likes"`
}
