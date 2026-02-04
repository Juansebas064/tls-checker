package ui

import (
	"encoding/json"
	"tls-checker/internal/api"
)

func (app *Application) AnalyzeHost(host string) {

	go func() {
		if app.visitedHosts[host] == nil {
			res, err := api.AnalyzeHost(host)
			if err != nil {
				panic(err)
			} else {
				// Add host to visited hosts
				app.visitedHosts[host] = res
			}
		}
		// Add endpoints (if available) to endpoints available

		// Show data
		text, _ := json.MarshalIndent(app.visitedHosts[host], "", "\t")
		app.tui.QueueUpdateDraw(func() {
			// Update VisitedHosts section
			app.hostsSection.Clear()
			for key := range app.visitedHosts {
				app.hostsSection.AddItem(key, "", 0, nil)
			}
			app.tui.SetFocus(app.hostsSection)
			// Update details section
			app.detailsSection.SetText(string(text))
		})
	}()
}
