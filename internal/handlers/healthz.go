package handlers

import (
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{
		"status":"ok"
	}`))
}
