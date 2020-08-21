package routes

import (
	"app/internal/database"
	"app/internal/utils"
	"net/http"
	"strconv"
)

// getWeekVerses returns the list of verses associated with a given week of a NFL season
//    /api/weekverses?week=WEEK&season=SEASON
func getWeekVerses(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	book := q["book"][0]
	week := q["week"][0]

	season := q["season"][0]
	seasonInt, err := strconv.ParseInt(season, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	v, err := database.GetWeekVerses(book, string(week), int(seasonInt))

	utils.ServeMarahallableData(w, r, v)

}

// getVerse returns the bible verses from a given book that correspond to the supplied params
//		/api/bibleverse?book=BOOK&verse=VERSE&chapter=CHAPTER
// This method is expecting a get request with book, verse, and chapter included
// If no verses are found, a json object with an empty Text field is returned
func getVerse(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	book := q["book"][0]

	chapter := q["chapter"][0]
	chapterInt, err := strconv.ParseInt(chapter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	verse := q["verse"][0]
	verseInt, err := strconv.ParseInt(verse, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	v, err := database.GetBibleVerse(book, int(chapterInt), int(verseInt))

	utils.ServeMarahallableData(w, r, v)
}

// getSupportedBooks simply returns a json object with a list of all query-able books of the bible
// No url params are expected or handled
func getSupportedBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetBibleBooks()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.ServeMarahallableData(w, r, books)
}

// getSeasonMeta returns all seasons and weeks data. No url params are expected or handled
func getSeasonsMeta(w http.ResponseWriter, r *http.Request) {

	seasons, err := database.GetSeasonMetaData()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.ServeMarahallableData(w, r, seasons)

}
