package show2b

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/zengqiang96/grab2b/internal/proto"
)

const (
	PageSize = 25
)

var PageNames = []string{"page-1", "page-2", "page-3", "page-4"}

type SongView struct {
	Header    []string
	SongPages *tview.Pages
	Songs     []proto.SongEntity
}

func (view *SongView) Refresh() {
	pageCount := view.SongPages.GetPageCount()
	if pageCount == 0 {
		view.InitPages()
	}

	for pageCount > 0 {
		pageName, pageItem := view.SongPages.GetFrontPage()

		ps := 0
		for row, song := range view.Songs {
			table := pageItem.(*tview.Table)
			table.SetCell(row+1, 0, tview.NewTableCell(song.Rank))
			table.SetCell(row+1, 1, tview.NewTableCell(song.Song))
			table.SetCell(row+1, 2, tview.NewTableCell(song.Artist))
			if ps
		}

	}

	for pageNo, pageName := range PageNames {
	}

	t := tview.NewTable()
	for index, header := range year100Header {
		t.SetCell(0, index, tview.NewTableCell(header))
	}
	for row, song := range songs {
		p := row/25 + 1
		t.SetCell(row+1, 0, tview.NewTableCell(song.Rank))
		t.SetCell(row+1, 1, tview.NewTableCell(song.Song))
		t.SetCell(row+1, 2, tview.NewTableCell(song.Artist))

		if row%25 == 0 {
			ui.songsPanel.AddPage(fmt.Sprintf("Page-%d", p), t, false, true)
			t = tview.NewTable()
			for index, header := range year100Header {
				t.SetCell(0, index, tview.NewTableCell(header))
			}
		}
	}
}

func (view *SongView) RemoveAndInit(header []string) {
	for pageNo, pageName := range PageNames {
		view.SongPages.RemovePage(pageName)

		t := tview.NewTable()
		for index, h := range header {
			t.SetCell(0, index, tview.NewTableCell(h))
		}

		view.SongPages.AddPage(pageName, t, false, pageNo == 0)
	}
}

func (view *SongView) RemoveAllPages() {
	for _, pageName := range PageNames {
		view.SongPages.RemovePage(pageName)
	}
}

func (view *SongView) InitPages() {
	for pageNo, pageName := range PageNames {
		t := tview.NewTable()
		for index, header := range view.Header {
			t.SetCell(0, index, tview.NewTableCell(header))
		}

		view.SongPages.AddPage(pageName, t, false, pageNo == 0)
	}
}
