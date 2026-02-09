package ui

import (
	"encoding/json"
	"fmt"
	"time"
	"tls-checker/internal/api"
	"tls-checker/internal/model"
)

// Call api.analyzeHost(host) with the given host name and add it to []visitedHosts if it was not present. Call app.hostChanged(host) at the end.
func (app *Application) searchHost(host string) {
	go func() {
		if _, exists := app.visitedHosts[host]; !exists {
			var (
				res *model.Host
				err error
				i   int = 0
			)
			for {
				if res != nil {
					res, err = api.AnalyzeHost(app.getAnalyzeHostQuery(true))	// Check status of running req
					app.showMessage(fmt.Sprintf("Retrying for %d time, status: %s", i, res.Status), "text")
				} else {
					res, err = api.AnalyzeHost(app.getAnalyzeHostQuery(false))
					app.showMessage("Starting API request", "text")
				}

				if err != nil {
					app.showMessage(err.Error(), "error")
					break
				} else {
					if res.Status != "READY" {
						time.Sleep(5 * time.Second)
					} else {
						// Add host to visited hosts
						app.visitedHosts[host] = res
						app.hostsSection.AddItem(host, "", 0, nil)
						app.showMessage("API request completed", "text")
						app.hostChanged(host)
						break
					}
				}
				i++
			}
		}
	}()
}

// Change the contents of endpointsSection and detailsSection when adding/navigating to another host in the list
func (app *Application) hostChanged(host string) {
	app.tui.SetFocus(app.hostsSection)
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
