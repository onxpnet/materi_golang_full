package pegawai

import (
	"fmt"
	"net/http"
	"strconv"
	"esdm/db"
)
// DELETE /pegawai?id=1
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	
	res, err := db.DB.Exec("DELETE FROM pegawai WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	count, _ := res.RowsAffected()
	if count == 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Deleted %d rows\n", count)
}


