package ui

import (
	"tls-checker/internal/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Application struct {
	// Global state for visited hosts and endpoints
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

// Creates and returns a new *tview.Application with all its components initialized
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

// Initializes the main TUI
func (app *Application) Run() {
	// Run tview.Application
	if err := app.tui.Run(); err != nil {
		panic(err)
	}
}

// Creates the TUI components and state variables of the application
func (app *Application) initUIComponents() {
	// Create sections for the TUI
	app.hostsSection = tview.NewList()
	app.hostsSection.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		app.hostChanged(mainText)
	})

	app.endpointsSection = tview.NewList()

	app.detailsSection = tview.NewTextView()
	app.detailsSection.SetText("Write a host address and hit enter to begin")

	app.searchBarSection = tview.NewInputField()
	app.searchBarSection.SetFieldWidth(30).
		SetFieldBackgroundColor(tcell.Color63).
		SetPlaceholder("www.example.com")

	app.messagesSection = tview.NewTextView()
	app.messagesSection.SetText("API Ready")
}

// Shorthand for updating the TUI inside a goroutine via tview.Application.QueueUpdateDraw() function
func (app *Application) queueUpdateDraw(function func()) {
	go func() {
		app.tui.QueueUpdateDraw(function)
	}()
}
