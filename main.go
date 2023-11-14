package main

import (
	"fmt"
	"microsg/controlers"
	"net/http"
)


func main() {
	http.HandleFunc("/login", controlers.LoginHandler)
	http.HandleFunc("/signup", controlers.SignupHandler)

	port := ":8080"
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
