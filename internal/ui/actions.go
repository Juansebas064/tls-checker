package ui

import (
	"encoding/json"
	"fmt"
	"tls-checker/internal/api"
)

func (app *Application) searchHost() {
	// TODO: Block the execution of another coroutine until the previous one finishes
	host := app.searchBarSection.GetText()
	go func() {
		if _, exists := app.visitedHosts[host]; !exists {
			res, err := api.AnalyzeHost(host)
			if err != nil {
				panic(err)
			} else {
				// Add host to visited hosts
				app.visitedHosts[host] = res
				// Re-draw hosts
				app.tui.QueueUpdateDraw(func() {
					app.hostsSection.Clear()
					for key := range app.visitedHosts {
						app.hostsSection.AddItem(key, "", 0, nil)
					}
				})
			}
		}
		app.hostChanged(host)
	}()
}

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

func (app *Application) endpointChanged() {

}
