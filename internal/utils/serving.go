package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func setHeaders(w http.ResponseWriter) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

}

// ServeJSON sets common headers for serving JSON data
func ServeJSON(w http.ResponseWriter, r *http.Request, b []byte) (int, error) {

	setHeaders(w)
	w.Header().Set("Content-Type", "application/json")

	wlen, err := w.Write(b)

	if err != nil {
		log.Print(err.Error())
		return wlen, err
	}

	return wlen, nil

}

// ServeMarshallableData serves to w the interface. Will throw an error if data cannot be marhalled
func ServeMarahallableData(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ServeJSON(w, r, b)
}
