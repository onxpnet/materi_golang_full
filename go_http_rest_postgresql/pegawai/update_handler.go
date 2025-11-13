package pegawai

import (
	//"encoding/json"
	"net/http"
	//"strconv"
	"fmt"
)

// PUT /pegawai?id=1
func Update(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.Query().Get("id")
	//id, _ := strconv.Atoi(idStr)
	//var updated Pegawai
	//if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	fmt.Fprintln(w, "")
}


