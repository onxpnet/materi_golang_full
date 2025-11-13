package pegawai

import (
	"fmt"
	"net/http"
	//"strconv"
)
// DELETE /pegawai?id=1
func Delete(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.Query().Get("id")
	//id, _ := strconv.Atoi(idStr)
	fmt.Fprintln(w, "")
}


