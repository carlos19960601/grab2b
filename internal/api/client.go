package api

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
	"github.com/zengqiang96/grab2b/internal/proto"
)

type Client struct {
	collector *colly.Collector
}

func New() *Client {
	return &Client{
		collector: colly.NewCollector(),
	}
}

func (c *Client) Hot100(date string) ([]proto.SongEntity, error) {
	songs := make([]proto.SongEntity, 0, 100)
	url := fmt.Sprintf("https://www.billboard.com/charts/hot-100/%s", date)
	headerColors := make([]tablewriter.Colors, len(header))
	for index := range header {
		headerColors[index] = tablewriter.Color(tablewriter.FgHiYellowColor)
	}

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.collector.OnHTML("li.chart-list__element", func(e *colly.HTMLElement) {
		rank := e.ChildText("span.chart-element__rank__number")
		trend := e.ChildText("span.chart-element__trend")
		song := e.ChildText("span.chart-element__information__song")
		artist := e.ChildText("span.chart-element__information__artist")
		delta := e.ChildText("span.chart-element__information__delta__text.text--default")
		last := e.ChildText("span.chart-element__meta.text--last")
		peak := e.ChildText("span.chart-element__meta.text--peak")
		week := e.ChildText("span.chart-element__meta.text--week")

		songs = append(songs, proto.SongEntity{
			Rank:   rank,
			Trend:  trend,
			Song:   song,
			Artist: artist,
			Delta:  delta,
			Last:   last,
			Peak:   peak,
			Week:   week,
		})
	})

	err := c.collector.Visit(url)
	if err != nil {
		return nil, err
	}

	c.collector.Wait()
	return nil, nil
}

func (c *Client) Year100(year string) ([]proto.SongEntity, error) {
	songs := make([]proto.SongEntity, 0, 100)

	url := "https://www.billboard.com/charts/year-end/hot-100-songs"
	if len(year) > 0 {
		url = fmt.Sprintf("https://www.billboard.com/charts/year-end/%s/hot-100-songs", year)
	}

	headerColors := make([]tablewriter.Colors, len(yearHeader))
	for index := range yearHeader {
		headerColors[index] = tablewriter.Color(tablewriter.FgHiYellowColor)
	}

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.collector.OnHTML("div.chart-details__item-list", func(e *colly.HTMLElement) {
		e.ForEach("article.ye-chart-item", func(index int, item *colly.HTMLElement) {
			rank := item.ChildText("div.ye-chart-item__rank")
			song := item.ChildText("div.ye-chart-item__title")
			artist := item.ChildText("div.ye-chart-item__artist")

			songs = append(songs, proto.SongEntity{
				Rank:   rank,
				Song:   song,
				Artist: artist,
			})
		})
	})

	err := c.collector.Visit(url)
	if err != nil {
		return nil, err
	}

	c.collector.Wait()
	return songs, nil
}
