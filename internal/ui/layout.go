package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *Application) buildLayout() {
	hostsFrame := tview.NewFrame(app.hostsSection)
	hostsFrame.AddText("Visited Hosts (Ctrl + h)", true, tview.AlignLeft, tcell.Color63).
		// AddText("Ctrl + h", false, tview.AlignLeft, tcell.Color110).
		SetBorders(0, 0, 0, 0, 1, 1).
		SetBorder(true)

	// Build layout
	hostsAndEndpointLayout := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(hostsFrame, 0, 1, false).
		AddItem(app.endpointsSection, 0, 1, false)

	firstColumnAndDetailsLayout := tview.NewFlex().
		SetDirection(tview.FlexRowCSS).
		AddItem(hostsAndEndpointLayout, 0, 1, false).
		AddItem(app.detailsSection, 0, 3, false)

	infoAndSearchHostLayout := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(firstColumnAndDetailsLayout, 0, 5, false).
		AddItem(app.searchBarSection, 0, 1, true)

	mainFrame := tview.NewFrame(infoAndSearchHostLayout).SetBorders(0, 0, 0, 0, 0, 0)
	mainFrame.AddText("TLS Checker - SSL Labs Api v2", true, tview.AlignCenter, tcell.Color63)

	// Set up layout in main app
	app.pages = tview.NewPages().AddPage("main", mainFrame, true, true)
	app.tui = tview.NewApplication().SetRoot(app.pages, true)

}