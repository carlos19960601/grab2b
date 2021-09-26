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
	artistHeader = []string{"rank", "title", "last", "peak", "weeks"}
)

func Artist100Action(ctx *cli.Context) error {
	date := ctx.String("date")
	url := fmt.Sprintf("https://www.billboard.com/charts/artist-100/%s", date)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(artistHeader)
	table.SetBorder(false)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetAutoWrapText(false)

	headerColors := make([]tablewriter.Colors, len(artistHeader))
	for index := range artistHeader {
		headerColors[index] = tablewriter.Color(tablewriter.FgHiYellowColor)
	}
	table.SetHeaderColor(headerColors...)

	c := colly.NewCollector()

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	c.OnHTML("div.chart-list", func(list *colly.HTMLElement) {
		list.ForEach(".chart-list-item", func(index int, item *colly.HTMLElement) {
			rank := item.ChildText("div.chart-list-item__rank")
			title := item.ChildText("div.chart-list-item__text-wrapper span.chart-list-item__title-text")

			var last, peak, weeks string
			item.ForEach("div.chart-list-item__extra-info div.chart-list-item__stats-cell", func(index int, stat *colly.HTMLElement) {
				switch index {
				case 0:
					last = stat.ChildText(".chart-list-item__last-week")
				case 2:
					peak = stat.ChildText(".chart-list-item__weeks-at-one")
				case 3:
					weeks = stat.ChildText(".chart-list-item__weeks-on-chart")
				default:
				}
			})

			table.Append([]string{rank, title, last, peak, weeks})
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
