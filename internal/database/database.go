package database

import (
	"app/pkg/models"
	"database/sql"
	"strconv"
)

var (
	// ConnectionString is the hardcoded database string
	ConnectionString string

	//DB is the shared database connection
	DB *sql.DB
)

// _GetBibleVerses returns all verses from the given book from the database that match the contents
// of the chapter and verse slice. chapter and verse must be the same length
func GetWeekVerses(book string, week string, season int) (v []models.Game, err error) {

	stmtString := `SELECT games.season_id, games.week_id, games.game_id, games.date, 
                        games.home_team, games.visiting_team, games.home_score, 
                        games.visiting_score, games.OT, bible.chapter, bible.verse, bible.vtext
                        FROM games 
                        LEFT JOIN bible 
                        WHERE bible.chapter = games.home_score AND
                        bible.verse = games.visiting_score AND
                        games.week_id = ? and games.season_id = ? AND
                        book = ?`

	stmt, err := DB.Prepare(stmtString)

	result, err := stmt.Query(week, season, book)

	if err != nil {
		return v, err
	}

	for result.Next() {
		var tempv models.BibleVerse
		var tempg models.Game
		if err = result.Scan(
			&tempg.SeasonID,
			&tempg.WeekID,
			&tempg.GameID,
			&tempg.Date,
			&tempg.HomeTeam,
			&tempg.VisitingTeam,
			&tempg.HomeScore,
			&tempg.VisitingScore,
			&tempg.OT,
			&tempv.Chapter,
			&tempv.Verse,
			&tempv.Text,
		); err != nil {
			return v, err
		}
		tempv.Book = book
		tempg.VerseData = tempv

		// Append the book name to the id to force Vue to redraw the same game with a different book selection
		tempg.GameID = tempg.GameID + tempv.Book
		v = append(v, tempg)
	}

	return v, nil
}

// GetBibleBooks returns a list of unique book names contained within the application database
// This is primarily used for rendering dropdown lists
func GetBibleBooks() (books []string, err error) {
	stmtString := "SELECT DISTINCT book FROM bible"

	stmt, err := DB.Prepare(stmtString)

	if err != nil {
		return books, err
	}

	result, err := stmt.Query()
	defer result.Close()

	var tmpb string
	for result.Next() {
		if err = result.Scan(&tmpb); err != nil {
			return books, err
		}
		books = append(books, tmpb)
	}
	return books, nil
}

// GetBibleVerse retrieves the verse object corresponding to the given book (Long form), chapter, and verse index
// If no verse was found a object with an empty text value is returned
func GetBibleVerse(book string, chapter int, verse int) (v models.BibleVerse, err error) {

	v.Book = book
	v.Chapter = chapter
	v.Verse = verse

	stmtString := "SELECT vtext FROM bible WHERE book = ? AND chapter = ? and verse = ?"

	stmt, err := DB.Prepare(stmtString)

	if err != nil {
		return v, err
	}

	result, err := stmt.Query(book, chapter, verse)
	defer result.Close()
	if err != nil {
		return v, err
	}

	for result.Next() {
		if err = result.Scan(&v.Text); err != nil {
			return v, err
		}
	}

	return v, nil
}

// GetSeasonMetaData returns all season-week combinations, including which books are valid for each week.
// Does not include acutal game data.
// This is primarily used for rendering dropdown lists for each season
func GetSeasonMetaData() (seasons []models.Season, err error) {

	// Select query. Guaranteed to be in accending order
	stmtString := `SELECT season_id, label, book FROM valid_books WHERE book IS NOT NULL`

	stmt, err := DB.Prepare(stmtString)

	if err != nil {
		return seasons, err
	}

	result, err := stmt.Query()
	defer result.Close()

	// We will fist scan the result into a []models.Week and further encode it as seasons later
	//var weeksResults, seasonsResults, booksResults []string
	var curYear int
	var curWeek string

	// This feels like kind of a hacky way to avoid index out of range errors
	// the first time that the counter is incremented. I need to fix this later
	sIndex := -1
	wIndex := -1

	for result.Next() {

		var tmpweek, tmpseason, tmpbook string
		if err = result.Scan(&tmpseason, &tmpweek, &tmpbook); err != nil {
			return seasons, err
		}

		// Create a new season and append it to the season slice
		tmpseasonInt, err := strconv.Atoi(tmpseason)
		if err != nil {
			return seasons, err
		}

		if tmpseasonInt != curYear {
			curYear = tmpseasonInt
			seasons = append(seasons, models.Season{})
			wIndex = -1
			sIndex++
		}

		// Create a new week and append it to the current season
		// I think this breaks if there are two weeks with identical labels but different
		// labels right next to each other in the table. I believe this is highly unlikely to happen.
		// It would essentially mean that there were no valid books for an entire season
		if tmpweek != curWeek {
			curWeek = tmpweek
			seasons[sIndex].Weeks = append(seasons[sIndex].Weeks, models.Week{})
			wIndex++
		}
		seasons[sIndex].Year = tmpseasonInt
		seasons[sIndex].Weeks[wIndex].SeasonID = tmpseasonInt
		seasons[sIndex].Weeks[wIndex].Label = tmpweek
		seasons[sIndex].Weeks[wIndex].ValidBooks = append(seasons[sIndex].Weeks[wIndex].ValidBooks, tmpbook)
	}
	return seasons, nil
}
