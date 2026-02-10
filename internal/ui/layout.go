package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

// Builds the layout of the app by its sections
func (app *Application) buildLayout() {

	// Frames for headers and shortcuts guide
	// Hosts
	hostsFrame := tview.NewFrame(app.hostsSection)
	hostsFrame.AddText("Hosts (Ctrl + h)", true, tview.AlignLeft, colors[colorPrimary]).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)
		
	// Endpoints
	endpointsFrame := tview.NewFrame(app.endpointsSection)
	endpointsFrame.AddText("Endpoints (Ctrl + e)", true, tview.AlignLeft, colors[colorPrimary]).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	// Details
	detailsFrame := tview.NewFrame(app.detailsSection)
	detailsFrame.AddText("Details (Ctrl + d)", true, tview.AlignCenter, colors[colorPrimary]).
	AddText(fmt.Sprintln("Previous (p)"), false, tview.AlignLeft, colors[colorText]).
	AddText(fmt.Sprintln("Next (n)"), false, tview.AlignRight, colors[colorText]).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	// Search
	searchBarFrame := tview.NewFrame(app.searchSection)
	searchBarFrame.AddText("Search (Ctrl + s)", true, tview.AlignCenter, colors[colorPrimary]).
		SetBorders(0, 0, 1, 0, 1, 1).
		SetBorder(true)

	// Arrange frames into the final layout
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
		AddItem(firstColumnAndDetailsLayout, 0, 20, false).
		AddItem(searchBarFrame, 0, 5, true).
		AddItem(app.messagesSection.SetTextAlign(tview.AlignCenter), 0, 1, false)

	// Main frame containing all the sections
	mainFrame := tview.NewFrame(infoAndSearchHostLayout).SetBorders(1, 0, 0, 0, 0, 0)
	mainFrame.AddText("TLS Checker - SSL Labs Api v2", true, tview.AlignCenter, colors[colorText])
	mainFrame.AddText(" Exit (Ctrl + q)", true, tview.AlignLeft, colors[colorText])

	// Set up final layout in main app
	app.pages = tview.NewPages().AddPage("main", mainFrame, true, true)
	app.tui = tview.NewApplication().SetRoot(app.pages, true)

}
