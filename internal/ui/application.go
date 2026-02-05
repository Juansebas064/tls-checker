package ui

import (
	"tls-checker/internal/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
	app.hostsSection.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune){
		app.hostChanged(mainText)
	})

	app.endpointsSection = tview.NewList()

	app.detailsSection = tview.NewTextView()
	app.detailsSection.SetText("Write a host address and hit enter to begin")

	app.searchBarSection = tview.NewInputField().SetFieldWidth(0)
	app.searchBarSection.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			app.searchHost()
		}
	})

	app.messagesSection = tview.NewTextView()
	app.messagesSection.SetText("API Ready")
}

func (app *Application) queueUpdateDraw(function func()) {
	go func() {
		app.tui.QueueUpdateDraw(function)
	}()
}