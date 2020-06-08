package common

import (
	"net/http"
)
//Run start server
func Run(){
	http.ListenAndServe(":8080", nil)
}