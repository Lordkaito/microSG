package controlers

import (
	"encoding/json"
	"fmt"
	"microsg/models"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isUserAuthenticated(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
		return
	}
	if r.Method == "POST" {
		NewPost(w, r)
		return
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	var postData models.Post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		fmt.Println(err)
	}
	post := models.Post{
		Id:     postData.Id,
		Title:  postData.Title,
		Text:   postData.Text,
		Author: postData.Author,
	}
	post.Create()

	// Write the user object as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
