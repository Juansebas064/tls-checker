package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *Application) buildLayout() {
	// Frames for headers and shortcuts guide
	hostsFrame := tview.NewFrame(app.hostsSection)
	hostsFrame.AddText("Hosts [Ctrl + h]", true, tview.AlignLeft, tcell.Color63).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	endpointsFrame := tview.NewFrame(app.endpointsSection)
	endpointsFrame.AddText("Endpoints [Ctrl + e]", true, tview.AlignLeft, tcell.Color63).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	detailsFrame := tview.NewFrame(app.detailsSection)
	detailsFrame.AddText("Details [Ctrl + d]", true, tview.AlignCenter, tcell.Color63).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	searchBarFrame := tview.NewFrame(app.searchBarSection)
	searchBarFrame.AddText("Search [Ctrl + s]", true, tview.AlignCenter, tcell.Color63).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	// Build layout
	hostsAndEndpointLayout := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(hostsFrame, 0, 1, false).
		AddItem(endpointsFrame, 0, 1, false)

	firstColumnAndDetailsLayout := tview.NewFlex().
		SetDirection(tview.FlexRowCSS).
		AddItem(hostsAndEndpointLayout, 0, 1, false).
		AddItem(detailsFrame, 0, 3, false)

	infoAndSearchHostLayout := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(firstColumnAndDetailsLayout, 0, 10, false).
		AddItem(searchBarFrame, 0, 2, true).
		AddItem(app.messagesSection.SetTextAlign(tview.AlignCenter), 0, 1, false)

	// Set the final layout
	mainFrame := tview.NewFrame(infoAndSearchHostLayout).SetBorders(1, 0, 0, 0, 0, 0)
	mainFrame.AddText("TLS Checker - SSL Labs Api v2", true, tview.AlignCenter, tcell.Color63)

	// Set up final layout in main app
	app.pages = tview.NewPages().AddPage("main", mainFrame, true, true)
	app.tui = tview.NewApplication().SetRoot(app.pages, true)

}
