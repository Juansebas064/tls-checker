package ui

// Sets events when specific functions are triggered in the components
func (app *Application) setEvents() {
	// Hosts
	app.hostsSection.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		app.hostChanged(index)
	})
	app.hostsSection.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if len(app.visitedHosts[index].Endpoints) == 1 {
			app.tui.SetFocus(app.detailsSection)
		} else {
			app.tui.SetFocus(app.endpointsSection)
		}
	})
	//Endpoints
	app.endpointsSection.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		app.endpointChanged(index)
	})
	app.endpointsSection.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		app.tui.SetFocus(app.detailsSection)
	})
	// Search
	app.searchSection.GetButton(0).SetSelectedFunc(func() {
		app.searchHost(app.hostField.GetText())
	})
	app.startNewCheck.SetChangedFunc(func(isChecked bool) {
		// TODO: Fix focusing another element when changed
		if isChecked {
			app.fromCacheCheck.SetChecked(false)
			app.maxAgeField.SetDisabled(true)
		}
	})
	app.fromCacheCheck.SetChangedFunc(func(isChecked bool) {
		// TODO: Fix focusing another element when changed
		if isChecked {
			app.startNewCheck.SetChecked(false)
			app.maxAgeField.SetDisabled(false)
		}
	})
}