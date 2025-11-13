package main

import (
	"fmt"
	"net/http"
	"esdm/pegawai"
)

func main() {
	http.HandleFunc("/pegawai", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pegawai.Get(w, r)
		case http.MethodPost:
			pegawai.Create(w, r)
		case http.MethodPut:
			pegawai.Update(w, r)
		case http.MethodDelete:
			pegawai.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
