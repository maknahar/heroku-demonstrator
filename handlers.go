package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GreetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()
	user := r.FormValue("name")
	if strings.TrimSpace(user) == "" {
		user = "there"
	}
	writeHTTPResponse([]byte("Hello "+user+"! Let's deploy on Heroku."), http.StatusOK, w, start, "text/plain")
}

func writeHTTPResponse(response []byte, statusCode int, w http.ResponseWriter, start time.Time, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(response)
}
