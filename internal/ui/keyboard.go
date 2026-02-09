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
	app.startNewCheck.SetChangedFunc(func (isChecked bool) {
		if isChecked {
			app.fromCacheCheck.SetChecked(false)
		}
	})
	app.fromCacheCheck.SetChangedFunc(func (isChecked bool) {
		if isChecked {
			app.startNewCheck.SetChecked(false)
		}
	})
}
