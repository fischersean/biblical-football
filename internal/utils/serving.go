package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ServeJSON sets common headers for serving JSON data
func ServeJSON(w http.ResponseWriter, r *http.Request, b []byte) (int, error) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
