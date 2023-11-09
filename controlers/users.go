package controlers

import (
	"encoding/json"
	"fmt"
	"microsg/models"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if !isUserAuthenticated(w, r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if r.Method == "POST" {
		NewUser(w, r)
		return
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	var userdata models.User
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		fmt.Println(err)
	}
	user := models.User{
		Name:     userdata.Name,
		Email:    userdata.Email,
		Password: userdata.Password,
	}
	user.Create()

	// Write the user object as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func isUserAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}
	var data models.Post
	arr := json.NewDecoder(resp.Body).Decode(&data)
	if arr != nil {
		fmt.Println("data id:", data.Id, "data author:", data.Author.Name, "data text:", data.Text, "data title:", data.Title)
	}
	return resp.StatusCode == http.StatusOK
}
