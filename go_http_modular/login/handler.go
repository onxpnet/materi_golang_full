package login

import (
	"fmt"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Login Page")
}
