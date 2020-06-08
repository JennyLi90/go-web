package common

import (
	"net/http"
)
//Run start server
func Run() error {
	return http.ListenAndServe(":8080", nil)
}