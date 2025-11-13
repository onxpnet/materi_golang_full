package pegawai

import (
	"encoding/json"
	"net/http"
	"esdm/db"
	"fmt"
)

// POST /pegawai
func Create(w http.ResponseWriter, r *http.Request) {
	var p Pegawai
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p.Nama, p.Jabatan)

	// Insert to DB and return new ID
	err := db.DB.QueryRow(
		"INSERT INTO pegawai (nama, jabatan) VALUES ($1, $2) RETURNING id",
		p.Nama, p.Jabatan,
	).Scan(&p.ID)
	if err != nil {
		fmt.Println("Insert ERROR:", err)
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}


