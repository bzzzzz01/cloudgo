package service

import (
	"net/http"
)

func unknown(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 unknown", 501)
}