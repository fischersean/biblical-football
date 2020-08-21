package utils

import (
	//"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func logRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"URL":    r.URL.EscapedPath(),
		"Method": r.Method,
		"Host":   r.Host,
	}).Info("Request Incoming")
}

func logResponse(w http.ResponseWriter) {
	log.WithFields(log.Fields{
		"Content-Type": w.Header()["Content-Type"],
	}).Info("Request Handled")
}

// Logged logs the details of the request and response
func Logged(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		h(w, r) // call original
		logResponse(w)
	})
}

// LoggedHandler logs the activity of a Handler interface
func LoggedHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		h.ServeHTTP(w, r) // call original
		logResponse(w)
	})
}
