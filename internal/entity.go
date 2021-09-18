package internal

import "time"

type Hot100 struct {
	WeekNo    int
	WeekStart time.Time
	WeekEnd   time.Time
	Chart     []*ChartElement
}

type ChartElement struct {
	Rank   int
	Trend  string // failing
	Song   string
	Artist string
	Delta  string
	Last   string
	Peak   string
	Week   string
}
