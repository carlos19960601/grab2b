package proto

import "time"

type Hot100 struct {
	WeekNo    int
	WeekStart time.Time
	WeekEnd   time.Time
	Chart     []*SongEntity
}

type YearSong struct {
	Rank   int
	Song   string
	Artist string
}

type SongEntity struct {
	Rank   string
	Trend  string // failing
	Song   string
	Artist string
	Delta  string
	Last   string
	Peak   string
	Week   string
}
