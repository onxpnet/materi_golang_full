package pegawai

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"esdm/db"
)

// PUT /pegawai?id=1
func Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	
	var p Pegawai
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec("UPDATE pegawai SET nama=$1, jabatan=$2 WHERE id=$3",
		p.Nama, p.Jabatan, id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	count, _ := res.RowsAffected()
	if count == 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Updated %d rows\n", count)
}


