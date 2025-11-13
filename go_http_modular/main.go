package main

import (
	"fmt"
	"net/http"
	"esdm/home"
	"esdm/profile"
	"esdm/login"
)

func main() {
	http.HandleFunc("/", home.Handler)
	http.HandleFunc("/profile", profile.Handler)
	http.HandleFunc("/login", login.Handler)

	fmt.Println("Serverd at http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
