package routes

import (
	//"flag"
	"app/internal/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	b := []byte("OK")
	w.Write(b)
}

// initAPI creates the routes for the app's api
func initAPI() {

	http.Handle("/", utils.Logged(healthCheck))
	http.Handle("/seasonsmeta", utils.Logged(getSeasonsMeta))
	http.Handle("/bibleverse", utils.Logged(getVerse))
	http.Handle("/weekverses", utils.Logged(getWeekVerses))
	http.Handle("/supportedbooks", utils.Logged(getSupportedBooks))

}

// initStaticFS assumes that we want "/" -> "index.html"
func initStaticFS(directory string) {
	return
	//log.Println("Creating FileServer @ " + directory)
	//fs := http.FileServer(http.Dir(directory))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// InitServer - Create routes and begin listening on provided port
func InitServer(port string) {

	//initStaticFS("./web/public")
	initAPI()

	log.Println("Listening on " + port + "...")
	log.Fatal(http.ListenAndServe(port, nil))
}
