package main

import (
	"fmt"
	"microsg/controlers"
	"net/http"
)

func main() {
	http.HandleFunc("/createUser", controlers.UserHandler)
	http.HandleFunc("/posts", controlers.PostHandler)

	port := ":8080"
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
