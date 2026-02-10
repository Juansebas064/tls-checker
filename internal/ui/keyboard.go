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
			// Next endpoint
			selectedItemIndex := app.endpointsSection.GetCurrentItem()
			if selectedItemIndex == (app.endpointsSection.GetItemCount() - 1) {
				app.endpointChanged(0)
			} else {
				app.endpointChanged(selectedItemIndex + 1)
			}
			app.tui.SetFocus(app.detailsSection)
		case 'p':
			// Previous endpoint
			selectedItemIndex := app.endpointsSection.GetCurrentItem()
			if selectedItemIndex == 0 {
				app.endpointChanged(app.endpointsSection.GetItemCount() - 1)
			} else {
				app.endpointChanged(selectedItemIndex - 1)
			}
			app.tui.SetFocus(app.detailsSection)
		}
		return event
	})
}
