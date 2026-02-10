package ui

import (
	"encoding/json"
	"fmt"
	"time"
	"tls-checker/internal/api"
	// "tls-checker/internal/model"
)

// Call api.analyzeHost(analyzeHostQuery) with the given host name and add it to []visitedHosts if it was not present. Call app.hostChanged(host) at the end.
func (app *Application) searchHost(host string) {
	go func() {
		if _, exists := app.visitedHosts[host]; !exists {
			// Make the request
			app.showMessage("Starting API request", "text")
			res, err := api.AnalyzeHost(app.getAnalyzeHostQuery(false))

			if err != nil {
				app.showMessage(err.Error(), "error")
				return
			}

			// Keep checking until status is ready
			for i := 1; ; i++ {
				if res.Status == "READY" {
					break
				}

				time.Sleep(10 * time.Second)

				res, err = api.AnalyzeHost(app.getAnalyzeHostQuery(true))
				app.showMessage(fmt.Sprintf("Retrying for %d time, status: %s", i, res.Status), "text")

				if err != nil {
					app.showMessage(err.Error(), "error")
					return
				}
			}

			// Add host to visited hosts
			app.visitedHosts[host] = res
			app.hostsSection.AddItem(host, "", 0, nil)
			app.showMessage("API request completed", "text")
			app.hostChanged(host)
		} else {
			app.hostChanged(host)
		}
	}()
}

// Change the contents of endpointsSection and detailsSection when adding/navigating to another host in the list
func (app *Application) hostChanged(host string) {
	app.tui.SetFocus(app.hostsSection)

	// TODO: If host isn't already in focus

	// if hostAlreadyOnFocus := app.hostsSection.

	app.queueUpdateDraw(func() {

		// Set host on focus
		if indexSlice := app.hostsSection.FindItems(host, "", false, false); len(indexSlice) > 0 {
			app.hostsSection.SetCurrentItem(indexSlice[0])
		}

		// Show endpoints of the host
		// app.endpointsSection.Clear()
		// for _, endpoint := range app.visitedHosts[host].Endpoints {
		// 	app.endpointsSection.AddItem(endpoint.IpAddress, "", 0, nil)
		// }

		// Show details of the host
		text, _ := json.MarshalIndent(app.visitedHosts[host], "", "\t")
		app.detailsSection.SetText(string(text))
	})
}

// TODO
func (app *Application) endpointChanged() {

}
