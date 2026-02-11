package ui

import (
	"tls-checker/internal/model"
	"tls-checker/internal/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Application struct {
	// Global state for visited hosts and endpoints
	visitedHosts []*model.Host

	// Application components
	tui              *tview.Application
	pages            *tview.Pages
	hostsSection     *tview.List
	endpointsSection *tview.List
	detailsSection   *tview.TextView
	summarySection   *tview.TextView
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
	app.visitedHosts = make([]*model.Host, 0)

	// Make the TUI components
	app.initUIComponents()

	// Build the layout
	app.buildLayout()

	// Set keyboard shortcuts
	app.setKeyboardShortcuts()

	// Set events for the TUI components
	app.setEvents()

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

	// Endpoints section
	app.endpointsSection = tview.NewList().ShowSecondaryText(false)

	// Details section
	app.detailsSection = tview.NewTextView()
	app.detailsSection.SetDynamicColors(true)
	app.detailsSection.SetText("Write a host address and hit send to begin")

	// Search section
	// Input for host
	app.hostField = tview.NewInputField().
		SetFieldWidth(25).
		SetPlaceholder("www.example.com").
		SetPlaceholderStyle(tcell.StyleDefault.Background(utils.ColorPrimary).Dim(true))

	// Flags
	app.startNewCheck = tview.NewCheckbox().SetLabel("Start new")
	app.publishCheck = tview.NewCheckbox().SetLabel("Publish")
	app.fromCacheCheck = tview.NewCheckbox().SetLabel("From cache").SetChecked(true)
	app.maxAgeField = tview.NewInputField().
		SetLabel("Max age").
		SetFieldWidth(4).
		SetPlaceholder("0").
		SetPlaceholderStyle(tcell.StyleDefault.Background(utils.ColorPrimary).Dim(true))
	app.ignoreMismatchCheck = tview.NewCheckbox().SetLabel("Ignore mismatch")

	// Search form
	app.searchSection = tview.NewForm()
	app.searchSection.AddFormItem(app.hostField).
		AddFormItem(app.startNewCheck).
		AddFormItem(app.publishCheck).
		AddFormItem(app.fromCacheCheck).
		AddFormItem(app.maxAgeField).
		AddFormItem(app.ignoreMismatchCheck).
		AddButton("Send", nil).
		SetHorizontal(true).
		SetFieldStyle(tcell.StyleDefault.Background(utils.ColorPrimary)).
		SetItemPadding(2).
		SetBorderPadding(0, 0, 0, 0)

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
func (app *Application) showMessage(message string, color tcell.Color) {
	app.queueUpdateDraw(func() {
		app.messagesSection.Clear()
		app.messagesSection.SetTextColor(color)
		app.messagesSection.SetText(message)
	})
}

// Build the request with the host address and flags
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
