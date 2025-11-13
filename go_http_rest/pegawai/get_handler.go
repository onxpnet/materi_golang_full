package pegawai

import (
	"encoding/json"
	"net/http"
)

// GET /pegawai
func Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(PegawaiDB)
}


