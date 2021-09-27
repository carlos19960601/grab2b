package show2b

import (
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
	view.RemoveAllPages()

	for index, pageName := range PageNames {
		songs := view.Songs[PageSize*index : PageSize*(index+1)]
		t := tview.NewTable()
		for hcol, header := range view.Header {
			t.SetCell(0, hcol, tview.NewTableCell(header))
		}

		for row, song := range songs {
			t.SetCell(row+1, 0, tview.NewTableCell(song.Rank))
			t.SetCell(row+1, 1, tview.NewTableCell(song.Song))
			t.SetCell(row+1, 2, tview.NewTableCell(song.Artist))
		}
		view.SongPages.AddPage(pageName, t, false, index == 0)
	}

	view.SongPages.ShowPage(PageNames[0])
}

func (view *SongView) RemoveAndInit() {
	for pageNo, pageName := range PageNames {
		view.SongPages.RemovePage(pageName)

		t := tview.NewTable()
		for index, h := range view.Header {
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
