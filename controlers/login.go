package controlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if r.Method == "POST" && token != "" {
		if isUserAuthenticated(token) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			loginUser(w, r)
		}
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/auth/login", r.Body)

	if err != nil {
		fmt.Println(err)
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
	fmt.Println("Response:", response)
	w.Write(jsonResponse)
}

func isUserAuthenticated(t string) bool {
	token := t

	type Response struct {
		Message string `json:"message"`
		User    string `json:"user,omitempty"`
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/auth/validate", nil)

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
	if response.Message != "Invalid token" {
		return true
	}
	return false
}
