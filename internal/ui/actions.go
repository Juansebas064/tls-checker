package ui

import (
	"fmt"
	"time"
	"tls-checker/internal/api"
	"tls-checker/internal/utils"
)

// Call api.analyzeHost(analyzeHostQuery) with the given host name and add it to []visitedHosts if it was not present. Call app.hostChanged(hostIndex) at the end.
func (app *Application) searchHost(host string) {
	go func() {
		// Evaluate if the host is already visited
		var indexSlice = app.hostsSection.FindItems(host, "", false, true)
		if len(indexSlice) > 0 {
			app.hostChanged(indexSlice[0])
		} else {
			// Make the request
			app.showMessage("Starting API request", utils.ColorText)
			res, err := api.AnalyzeHost(app.getAnalyzeHostQuery(false))

			if err != nil {
				app.showMessage(err.Error(), utils.ColorError)
				return
			}

			// Keep checking until status is ready
			// TODO: Move retrying logic to client.analyzeHost
			for i := 1; ; i++ {
				if res.Status == "READY" {
					break
				}

				time.Sleep(10 * time.Second)

				res, err = api.AnalyzeHost(app.getAnalyzeHostQuery(true))
				app.showMessage(fmt.Sprintf("Retrying for %d time, status: %s", i, res.Status), utils.ColorText)

				if err != nil {
					app.showMessage(err.Error(), utils.ColorError)
					return
				}
			}

			// Add host to visited hosts
			app.visitedHosts = append(app.visitedHosts, res)
			app.hostsSection.AddItem(host, "", 0, nil)
			app.showMessage("API request completed", utils.ColorText)
			app.hostChanged(len(app.visitedHosts) - 1)
		}
		app.tui.SetFocus(app.detailsSection)
	}()
}

// Change focus and the contents of endpointsSection when adding/navigating to another host in the list
func (app *Application) hostChanged(hostIndex int) {
	app.showMessage(fmt.Sprintf("Showing data from host \"%s\"", app.visitedHosts[hostIndex].Host), utils.ColorText)

	// TODO: Verify if host isn't already in focus
	// if hostAlreadyOnFocus := app.hostsSection.

	app.queueUpdateDraw(func() {

		// Set host on focus
		app.hostsSection.SetCurrentItem(hostIndex)

		// Show endpoints of the host
		app.endpointsSection.Clear()
		for _, endpoint := range app.visitedHosts[hostIndex].Endpoints {
			app.endpointsSection.AddItem(endpoint.IPAddress, "", 0, nil)
		}

		// Show details of the endpoint
		app.endpointChanged(0)
	})
}

// Change focus and the contents of detailsSection when adding/navigating to another endpoint in the list
func (app *Application) endpointChanged(endpointIndex int) {
	// Identify active host
	hostIndex := app.hostsSection.GetCurrentItem()
	endpoint := app.visitedHosts[hostIndex].Endpoints[endpointIndex]
	formatted := utils.FormatEndpoint(&endpoint)

	// Draw new endpoint details
	app.queueUpdateDraw(func() {
		app.endpointsSection.SetCurrentItem(endpointIndex)
		app.detailsSection.ScrollToBeginning()
		app.detailsSection.SetText(formatted)
	})
}
