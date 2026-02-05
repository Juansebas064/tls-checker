package ui

import (
	"encoding/json"
	"fmt"
	"tls-checker/internal/api"
)

// Call api.analyzeHost(host) with the given host name and add it to []visitedHosts if it was not present. Call app.hostChanged(host) at the end.
func (app *Application) searchHost(host string) {
	// TODO: Block the execution of another coroutine until the previous one finishes
	go func() {
		if _, exists := app.visitedHosts[host]; !exists {
			app.searchBarSection.SetDisabled(true)
			res, err := api.AnalyzeHost(host)
			if err != nil {
				panic(err)
			} else {
				// Add host to visited hosts
				app.visitedHosts[host] = res
				app.hostsSection.AddItem(host, "", 0, nil)
			}
		}
		app.searchBarSection.SetDisabled(false)
		app.hostChanged(host)
	}()
}

// Change the contents of endpointsSection and detailsSection when adding/navigating to another host in the list
func (app *Application) hostChanged(host string) {
	app.tui.SetFocus(app.hostsSection)
	app.queueUpdateDraw(func() {

		// Set host on focus
		if indexSlice := app.hostsSection.FindItems(host, "", false, false); len(indexSlice) > 0 {
			app.messagesSection.Clear()
			app.messagesSection.SetText(fmt.Sprintf("Current host: %v", host))
			app.hostsSection.SetCurrentItem(indexSlice[0])
		}

		// Show endpoints of the host
		app.endpointsSection.Clear()
		for _, endpoint := range app.visitedHosts[host].Endpoints {
			app.endpointsSection.AddItem(endpoint.IpAddress, "", 0, nil)
		}

		// Show details of the host
		text, _ := json.MarshalIndent(app.visitedHosts[host], "", "\t")
		app.detailsSection.SetText(string(text))
	})
}

// TODO
func (app *Application) endpointChanged() {

}
