package controlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	User    User   `json:"user,omitempty"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      string `json:"id"`
	Author  string `json:"author"`
	User_id string `json:"user_id"`
	Likes   []Like `json:"likes"`
}

type Like struct {
	Post_id string `json:"post_id"`
	User_id string `json:"user_id"`
	ID      string `json:"id"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
	Posts []Post `json:"posts"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/auth/signup", r.Body)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	defer resp.Body.Close()

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	fmt.Println("Response:", response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
