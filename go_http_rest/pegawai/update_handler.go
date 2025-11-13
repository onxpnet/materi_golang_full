package pegawai

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// PUT /pegawai?id=1
func Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	var updated Pegawai
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, p := range PegawaiDB {
		if p.ID == id {
			PegawaiDB[i].Nama = updated.Nama
			PegawaiDB[i].Jabatan = updated.Jabatan
			json.NewEncoder(w).Encode(PegawaiDB[i])
			return
		}
	}
	http.NotFound(w, r)
}


