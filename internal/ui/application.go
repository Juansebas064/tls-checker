package ui

import (
	"github.com/rivo/tview"
	"tls-checker/internal/model"
)

type Application struct {
	visitedHosts map[string]*model.Host

	// Application components
	tui              *tview.Application
	pages            *tview.Pages
	hostsSection     *tview.List
	endpointsSection *tview.List
	detailsSection   *tview.TextView
	searchBarSection *tview.InputField
	messagesSection  *tview.TextView
}

func NewApplication() *Application {
	// Initialize variables
	app := Application{}
	app.tui = tview.NewApplication()
	app.visitedHosts = make(map[string]*model.Host)

	// Make the UI components
	app.initUIComponents()

	// Build the layout
	app.buildLayout()

	// Set keyboard shortcuts
	app.setKeyboardShortcuts()

	return &app
}

func (app *Application) Run() {
	// Run tview.Application
	if err := app.tui.Run(); err != nil {
		panic(err)
	}
}

func (app *Application) initUIComponents() {
	// Create sections for the UI
	app.hostsSection = tview.NewList()

	app.endpointsSection = tview.NewList()
	app.endpointsSection.SetBorder(true)

	app.detailsSection = tview.NewTextView()
	app.detailsSection.SetBorder(true)
	app.detailsSection.SetText("Write a host address and hit enter to begin")

	app.searchBarSection = tview.NewInputField().SetFieldWidth(0)
	app.searchBarSection.SetBorder(true).SetTitle("Search")
	// SetDoneFunc(func(key tcell.Key) {
	// 	switch key {
	// 	case tcell.KeyEnter:
	// 		host := app.searchBarSection.GetText()
	// 		AnalyzeHost(host)
	// 	}
	// })
}
