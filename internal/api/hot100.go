package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

var (
	header = []string{"rank", "trend", "song", "artist", "delta", "last", "peak", "weeks"}
)

func Hot100Action(ctx *cli.Context) error {
	date := ctx.String("date")
	url := fmt.Sprintf("https://www.billboard.com/charts/hot-100/%s", date)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetAutoWrapText(false)

	headerColors := make([]tablewriter.Colors, len(header))
	for index := range header {
		headerColors[index] = tablewriter.Color(tablewriter.FgHiYellowColor)
	}
	table.SetHeaderColor(headerColors...)

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

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.OnHTML("li.chart-list__element", func(e *colly.HTMLElement) {
		rank := e.ChildText("span.chart-element__rank__number")
		trend := e.ChildText("span.chart-element__trend")
		song := e.ChildText("span.chart-element__information__song")
		artist := e.ChildText("span.chart-element__information__artist")
		delta := e.ChildText("span.chart-element__information__delta__text.text--default")
		last := e.ChildText("span.chart-element__meta.text--last")
		peak := e.ChildText("span.chart-element__meta.text--peak")
		week := e.ChildText("span.chart-element__meta.text--week")

		table.Append([]string{rank, trend, song, artist, delta, last, peak, week})
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	c.Wait()

	table.Render()
	return nil
}
