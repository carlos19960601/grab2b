package internal

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/urfave/cli/v2"
)

func Hot100Action(ctx *cli.Context) error {
	date := ctx.String("date")
	url := fmt.Sprintf("https://www.billboard.com/charts/hot-100/%s", date)

	c := colly.NewCollector()

	// 最新周榜
	// resp, err := client.R().Get("https://www.billboard.com/charts/hot-100")
	// 日期周榜
	// resp, err := client.R().Get("https://www.billboard.com/charts/hot-100/2020-01-01")
	// 年榜
	// resp, err := client.R().Get("https://www.billboard.com/charts/year-end/2012/hot-100-songs")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("  Body       :\n", resp)

	hot100 := Hot100{
		Chart: make([]*ChartElement, 100),
	}

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.OnHTML("li.chart-list__element", func(e *colly.HTMLElement) {
		rank, err := strconv.Atoi(e.ChildText("span.chart-element__rank__number"))
		if err != nil {
			log.Fatal(err)
		}

		trend := e.ChildText("span.chart-element__trend")
		song := e.ChildText("span.chart-element__information__song")
		artist := e.ChildText("span.chart-element__information__artist")
		delta := e.ChildText("span.chart-element__information__delta__text.text--default")
		last := e.ChildText("span.chart-element__meta.text--last")
		peak := e.ChildText("span.chart-element__meta.text--peak")
		week := e.ChildText("span.chart-element__meta.text--week")

		ele := ChartElement{
			Rank:   rank,
			Trend:  trend,
			Song:   song,
			Artist: artist,
			Delta:  delta,
			Last:   last,
			Peak:   peak,
			Week:   week,
		}

		hot100.Chart[ele.Rank-1] = &ele
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	c.Wait()

	log.Println("rank trend song artist delta last peak week")
	for _, ele := range hot100.Chart {
		log.Printf("%d %s %s %s %s %s %s %s \n", ele.Rank, ele.Trend, ele.Song, ele.Artist, ele.Delta, ele.Last, ele.Peak, ele.Week)
	}

	return nil
}
