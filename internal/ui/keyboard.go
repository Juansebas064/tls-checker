package ui

import "github.com/gdamore/tcell/v2"

// Sets the keybinds for navigating across the UI
func (app *Application) setKeyboardShortcuts() {
	// Main window keybinds
	app.tui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case rune(tcell.KeyCtrlH):
			app.tui.SetFocus(app.hostsSection)
		case rune(tcell.KeyCtrlE):
			app.tui.SetFocus(app.endpointsSection)
		case rune(tcell.KeyCtrlD):
			app.tui.SetFocus(app.detailsSection)
		case rune(tcell.KeyCtrlS):
			app.tui.SetFocus(app.searchSection)

		// Exit the app
		case rune(tcell.KeyCtrlQ):
			app.tui.Stop()
		}
		return event
	})

	// Section-related keybinds
	// Details
	app.detailsSection.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'n':
			// TODO: Next endpoint
		}
		return event
	})

	// Function assignment to components
	// Hosts
	app.hostsSection.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		app.hostChanged(mainText)
	})
	// Search
	app.searchSection.GetButton(0).SetSelectedFunc(func () {
		app.searchHost(app.hostField.GetText())
	})
	app.startNewCheck.SetChangedFunc(func (isChecked bool) {
		if isChecked {
			app.fromCacheCheck.SetChecked(false)
			app.maxAgeField.SetDisabled(true)
		}
	})
	app.fromCacheCheck.SetChangedFunc(func (isChecked bool) {
		if isChecked {
			app.startNewCheck.SetChecked(false)
			app.maxAgeField.SetDisabled(false)
		}
	})
}
