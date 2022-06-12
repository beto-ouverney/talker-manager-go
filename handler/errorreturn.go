package handler

import "net/http"

func errorReturn(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte("{\"message\":\"" + message + "\"}"))
}
