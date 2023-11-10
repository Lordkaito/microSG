package controlers

import (
	"encoding/json"
	"fmt"
	"microsg/models"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err == nil {
		fmt.Println(data, r.Header)
	}
	token := r.Header.Get("Authorization")
	if !isUserAuthenticated(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
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

func isUserAuthenticated(t string) bool {
	token := t
	type Response struct {
		Message string `json:"message"`
		User    string `json:"user,omitempty"`
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/auth/validate", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Message:", response.Message)
	if response.User != "" {
		fmt.Println("User:", response)
		fmt.Println("User:", response.User)
	}
	if response.Message != "Wrong token" {
		return true
	}
	return false
}
