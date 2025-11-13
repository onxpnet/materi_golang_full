package pegawai

import (
	"encoding/json"
	"net/http"
	"esdm/db"
	"fmt"
)

// GET /pegawai
func Get(res http.ResponseWriter, req *http.Request) {

	rows, err := db.DB.Query("SELECT id, nama, jabatan FROM pegawai ORDER BY id")
	if err != nil {
		fmt.Println("Error Query")
		http.Error(res, err.Error(), 500)
		return
	}
	defer rows.Close()

	var list []Pegawai
	for rows.Next() {
		var p Pegawai
		rows.Scan(&p.ID, &p.Nama, &p.Jabatan)
		list = append(list, p)
	}
	
	if len(list) == 0 {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"message":"no data found"}`))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(list)
}


