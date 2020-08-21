package main

import (
	"app/pkg/models"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: Should print number of rows affected to logs. I am currentyl ignoring the first return val for stmt.Exec()

func createTables(db *sql.DB) (err error) {

	// One off function to simplify the loop below
	var runCreate = func(stmt *sql.Stmt) (err error) {
		_, err = stmt.Exec()
		if err != nil {
			return err
		}

		return nil
	}
	seasonsStmtString := "CREATE TABLE seasons (id INTEGER PRIMARY KEY AUTOINCREMENT, season INTEGER)"
	weeksStmtString := "CREATE TABLE weeks (id INTEGER PRIMARY KEY AUTOINCREMENT, season_id INTEGER, label TEXT, order_ INTEGER)"
	gamesStmtString := `CREATE TABLE games 
								(id INTEGER PRIMARY KEY AUTOINCREMENT, 
								season_id INTEGER, 
								week_id INTEGER, 
								game_id TEXT,
								date TEXT, 
								home_team TEXT, 
								visiting_team TEXT,
								home_score INTEGER,
								visiting_score INTEGER, 
								OT BOOLEAN)`
	bibleStmtString := "CREATE TABLE bible (id INTEGER PRIMARY KEY AUTOINCREMENT, chapter INTEGER, book TEXT, verse INTEGER, vtext TEXT)"

	createMap := map[string]string{"seasons": seasonsStmtString, "weeks": weeksStmtString, "games": gamesStmtString, "bible": bibleStmtString}

	var stmt *sql.Stmt

	for key, value := range createMap {
		log.Println("Creating", key, "Table")

		stmt, err = db.Prepare(value)

		if err != nil {
			return err
		}

		if err := runCreate(stmt); err != nil {
			log.Println("Failed to create", key, " table")
			return err
		}

	}
	return nil
}

func insertGames(db *sql.DB, s models.Season) (err error) {

	stmtString := "INSERT INTO games (season_id, week_id, game_id, date, home_team, visiting_team, home_score, visiting_score, OT) VALUES "

	var inserts []interface{}
	for _, game := range s.Games {
		stmtString += "(?, ?, ?, ?, ?, ?, ?, ?, ?),"
		inserts = append(inserts, s.Year, game.WeekID, game.GameID, game.Date, game.HomeTeam, game.VisitingTeam, game.HomeScore, game.VisitingScore, game.OT)
	}

	stmt, err := db.Prepare(stmtString[0 : len(stmtString)-1])

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(inserts...); err != nil {
		return err
	}
	return nil
}

func insertWeeks(db *sql.DB, s models.Season) (err error) {

	stmtString := "INSERT INTO weeks (season_id, label, order_) VALUES "

	var inserts []interface{}
	for _, week := range s.Weeks {
		stmtString += "(?, ?, ?),"
		inserts = append(inserts, s.Year, week.Label, week.Order)
	}

	stmt, err := db.Prepare(stmtString[0 : len(stmtString)-1])

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(inserts...); err != nil {
		return err
	}
	return nil
}

func createSeasonRecord(db *sql.DB, s models.Season) (err error) {

	stmtString := "INSERT INTO seasons (season) VALUES(?)"
	stmt, err := db.Prepare(stmtString)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(s.Year); err != nil {
		return err
	}

	if err = insertGames(db, s); err != nil {
		return err
	}

	if err = insertWeeks(db, s); err != nil {
		return err
	}

	return nil
}

func createVerseRecord(db *sql.DB, v models.BibleVerse) (err error) {

	stmtString := "INSERT INTO bible (chapter, book, verse, vtext) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(stmtString)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(v.Chapter, v.Book, v.Verse, v.Text); err != nil {
		return err
	}

	return nil
}
func populateGameDataTables(db *sql.DB, data []models.Season) (err error) {
	for _, season := range data {
		if err = createSeasonRecord(db, season); err != nil {
			return err
		}
	}
	return nil
}

func populateBibleTables(db *sql.DB, data []models.BibleVerse) (err error) {
	for _, verse := range data {
		if err = createVerseRecord(db, verse); err != nil {
			return err
		}
	}
	return nil
}

func downloadFile(url string) (b []byte, err error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	b, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func popGameData(db *sql.DB) (err error) {

	// Download file from fdb-repo
	log.Println("Fetching raw game data from GitHub")
	fURL := "https://raw.githubusercontent.com/fischersean/fdb-scrape/master/exports/gamedata.json"
	b, err := downloadFile(fURL)

	//f, err := ioutil.ReadFile(gdataFile)
	if err != nil {
		return err
	}

	log.Println("Successfully downloaded game data")

	// Read in JSON encoded data
	var sdata []models.Season
	if err := json.Unmarshal(b, &sdata); err != nil {
		return err
	}

	log.Println("Successfully parsed game json")

	// Populate tables
	if err := populateGameDataTables(db, sdata); err != nil {
		return err
	}
	return nil
}

func popBibleData(db *sql.DB) (err error) {

	log.Println("Fetching raw bible data from GitHub")
	fURL := "https://raw.githubusercontent.com/fischersean/json-bible/master/exports/bible.json"
	b, err := downloadFile(fURL)

	//f, err := ioutil.ReadFile(gdataFile)
	if err != nil {
		return err
	}

	log.Println("Successfully downloaded bible data")

	// Read in JSON encoded data
	var bdata []models.BibleVerse
	if err := json.Unmarshal(b, &bdata); err != nil {
		return err
	}

	log.Println("Successfully parsed bible json")

	// Populate tables
	if err := populateBibleTables(db, bdata); err != nil {
		return err
	}
	return nil
}

func popDerivedTables(db *sql.DB) (err error) {

	stmtString := `CREATE TABLE valid_books AS
                        SELECT DISTINCT weeks.season_id, weeks.label, bible.book
                        FROM games 
                        LEFT JOIN bible 
                        ON (bible.chapter = games.home_score AND
                        bible.verse = games.visiting_score)
						LEFT JOIN weeks
						ON (weeks.season_id = games.season_id AND weeks.label = games.week_id)`

	stmt, err := db.Prepare(stmtString)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(); err != nil {
		return err
	}

	return nil
}

func main() {

	dbName := "./app.db"
	//gdataFile := "./resources/gamedata.json"
	log.Println("Creating new application database")

	if _, err := os.Stat(dbName); err == nil {
		log.Println("Found pre-exisiting database. Deleting and re-creating")
		os.Remove(dbName)
	}

	db, _ := sql.Open("sqlite3", dbName)
	// Create tables
	if err := createTables(db); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := popGameData(db); err != nil {
		panic(err.Error())
	}

	if err := popBibleData(db); err != nil {
		panic(err.Error())
	}

	if err := popDerivedTables(db); err != nil {
		panic(err.Error())
	}

	log.Println("Database population complete")

}
