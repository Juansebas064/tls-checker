package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"tls-checker/internal/model"
)

// Constants
var colors = map[string]tcell.Color{
	"primary": tcell.Color63,
	"text":    tcell.ColorWhite,
	"ok":      tcell.Color100,
	"warning": tcell.Color100,
	"error":   tcell.Color100,
}

type Application struct {
	// Global state for visited hosts and endpoints
	visitedHosts map[string]*model.Host

	// Application components
	tui              *tview.Application
	pages            *tview.Pages
	hostsSection     *tview.List
	endpointsSection *tview.List
	detailsSection   *tview.TextView
	searchSection    *tview.Form
	messagesSection  *tview.TextView

	// Flags
	hostField           *tview.InputField
	startNewCheck       *tview.Checkbox
	publishCheck        *tview.Checkbox
	fromCacheCheck      *tview.Checkbox
	maxAgeField         *tview.InputField
	ignoreMismatchCheck *tview.Checkbox
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
	// Host list
	app.hostsSection = tview.NewList().ShowSecondaryText(false)
	app.hostsSection.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		app.hostChanged(mainText)
	})

	// Endpoints section
	app.endpointsSection = tview.NewList().ShowSecondaryText(false)

	// Details section
	app.detailsSection = tview.NewTextView()
	app.detailsSection.SetText("Write a host address and hit send to begin")

	// Input for host
	app.hostField = tview.NewInputField().
		SetFieldWidth(25).
		SetPlaceholder("www.example.com").
		SetPlaceholderStyle(tcell.StyleDefault.Background(colors["primary"]).Dim(true))

	// Flags
	app.startNewCheck = tview.NewCheckbox().SetLabel("Start new").SetChecked(true)
	app.publishCheck = tview.NewCheckbox().SetLabel("Publish")
	app.fromCacheCheck = tview.NewCheckbox().SetLabel("From cache")
	app.maxAgeField = tview.NewInputField().
		SetLabel("Max age").
		SetFieldWidth(4).
		SetPlaceholder("0").
		SetPlaceholderStyle(tcell.StyleDefault.Background(colors["primary"]).Dim(true))
	app.ignoreMismatchCheck = tview.NewCheckbox().SetLabel("Ignore mismatch")

	// Search form
	app.searchSection = tview.NewForm()
	app.searchSection.AddFormItem(app.hostField).
		AddFormItem(app.startNewCheck).
		AddFormItem(app.publishCheck).
		AddFormItem(app.fromCacheCheck).
		AddFormItem(app.maxAgeField).
		AddFormItem(app.ignoreMismatchCheck).
		AddButton("Send", func() { app.searchHost(app.hostField.GetText()) }).
		SetHorizontal(true).
		SetFieldStyle(tcell.StyleDefault.Background(colors["primary"])).
		SetItemPadding(2).
		SetBorderPadding(0,0,0,0)

	// Messages section
	app.messagesSection = tview.NewTextView()
	app.messagesSection.SetText("API Ready")
}

// Shorthand for updating the TUI inside a goroutine via tview.Application.QueueUpdateDraw() function
func (app *Application) queueUpdateDraw(function func()) {
	go func() {
		app.tui.QueueUpdateDraw(function)
	}()
}

// Show messages to the user
func (app *Application) showMessage(message string, status string) {
	app.queueUpdateDraw(func() {
		app.messagesSection.Clear()
		app.messagesSection.SetTextColor(colors[status])
		app.messagesSection.SetText(message)
	})
}

func (app *Application) getAnalyzeHostQuery(checkRequestStatus bool) *model.AnalyzeHostQuery {
	var (
		getStatus = func(isChecked bool) string {
			if isChecked {
				return "on"
			} else {
				return "off"
			}
		}
		host           string = app.hostField.GetText()
		publish        string = getStatus(app.publishCheck.IsChecked())
		startNew       string
		fromCache      string = getStatus(app.fromCacheCheck.IsChecked())
		maxAge         string = app.maxAgeField.GetText()
		all            string = "done"
		ignoreMismatch string = getStatus(app.ignoreMismatchCheck.IsChecked())
	)

	if checkRequestStatus {
		startNew = "off"
	} else {
		startNew = getStatus(app.startNewCheck.IsChecked())
	}

	query := model.AnalyzeHostQuery{
		Host:           host,
		Publish:        publish,
		StartNew:       startNew,
		FromCache:      fromCache,
		MaxAge:         maxAge,
		All:            all,
		IgnoreMismatch: ignoreMismatch,
	}

	return &query
}
