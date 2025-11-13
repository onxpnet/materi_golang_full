package pegawai

import (
	"encoding/json"
	"net/http"
)

// POST /pegawai
func Create(w http.ResponseWriter, r *http.Request) {
	var p Pegawai
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = NextID
	NextID++
	PegawaiDB = append(PegawaiDB, p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}


