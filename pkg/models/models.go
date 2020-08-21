package models

type Game struct {
	SeasonID      int    `sql:"season_id"`
	WeekID        string `json:"Week,omitempty" sql:"week_id"`
	GameID        string `json:"id,omitempty"`
	Date          string `json:"Date,omitempty"`
	HomeTeam      string `json:"Home,omitempty"`
	VisitingTeam  string `json:"Visitor,omitempty"`
	HomeScore     int    `json:"HScore,omitempty"`
	VisitingScore int    `json:"VScore,omitempty"`
	OT            bool   `json:"OT,omitempty"`
	VerseData     BibleVerse
}

type Week struct {
	SeasonID   int      `json:"-" sql:"season_id"`
	Label      string   `json:"Label,omitempty" sql:"label"`
	Order      int      `json:"Order,omitempty" sql:"order_"`
	ValidBooks []string `json:"ValidBooks,omitempty"`
}

type Season struct {
	Year  int    `json:"Season,omitempty" sql:"season"`
	Weeks []Week `json:"Weeks,omitempty"`
	Games []Game `json:"Games,omitempty"`
}

type BibleVerse struct {
	Book    string `json:"Book"`
	Chapter int    `json:"Chapter"`
	Verse   int    `json:"Verse"`
	Text    string `json:"Text"`
}
