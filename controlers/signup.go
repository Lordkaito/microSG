package controlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	User    string `json:"user,omitempty"`
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
		w.Write([]byte(err.Error()))
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

	defer resp.Body.Close()

	var response Response

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResponse)
}
