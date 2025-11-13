package pegawai

import (
	"fmt"
	"net/http"
	"strconv"
)
// DELETE /pegawai?id=1
func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	for i, p := range PegawaiDB {
		if p.ID == id {
			PegawaiDB = append(PegawaiDB[:i], PegawaiDB[i+1:]...)
			fmt.Fprintf(w, "Pegawai ID %d deleted\n", id)
			return
		}
	}
	http.NotFound(w, r)
}


