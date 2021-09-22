package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

var (
	yearHeader = []string{"rank", "song", "artist"}
)

func Year100Action(ctx *cli.Context) error {
	date := ctx.String("year")

	url := "https://www.billboard.com/charts/year-end/hot-100-songs"
	if len(date) > 0 {
		url = fmt.Sprintf("https://www.billboard.com/charts/year-end/%s/hot-100-songs", date)

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(yearHeader)
	table.SetBorder(false)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetAutoWrapText(false)

	headerColors := make([]tablewriter.Colors, len(yearHeader))
	for index := range yearHeader {
		headerColors[index] = tablewriter.Color(tablewriter.FgHiYellowColor)
	}
	table.SetHeaderColor(headerColors...)

	c := colly.NewCollector()

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.OnHTML("div.chart-details__item-list", func(e *colly.HTMLElement) {
		e.ForEach("article.ye-chart-item", func(index int, item *colly.HTMLElement) {
			rank := item.ChildText("div.ye-chart-item__rank")
			song := item.ChildText("div.ye-chart-item__title")
			artist := item.ChildText("div.ye-chart-item__artist")

			table.Append([]string{rank, song, artist})
		})
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	c.Wait()

	table.Render()
	return nil
}
