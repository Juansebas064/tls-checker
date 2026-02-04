package ui

import "github.com/gdamore/tcell/v2"

func (app *Application) setKeyboardShortcuts() {
	app.tui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Rune() {
		case rune(tcell.KeyCtrlS):
			app.tui.SetFocus(app.searchBarSection)
		case rune(tcell.KeyCtrlH):
			app.tui.SetFocus(app.hostsSection)
		// Exit the app
		case rune(tcell.KeyCtrlQ):
			app.tui.Stop()
		}

		return event
	})
}