package ui

import (
	"fmt"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartApp(data []models.Artist) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	myWindow.Resize(fyne.NewSize(1000, 600))

	// --- RIGHT SIDE: Detail View ---
	detailsContainer := container.NewVBox(widget.NewLabel("Select an artist to see details"))

	// Function to update details
	updateDetails := func(a models.Artist) {
		detailsContainer.Objects = nil // Clear previous

		// Name
		title := canvas.NewText(a.Name, color.White)
		title.TextSize = 24
		title.TextStyle = fyne.TextStyle{Bold: true}

		// Info
		info := widget.NewLabel(fmt.Sprintf("Created: %d | First Album: %s", a.CreationDate, a.FirstAlbum))

		// Members
		membersLabel := widget.NewLabel("Members:\n" + strings.Join(a.Members, ", "))
		membersLabel.Wrapping = fyne.TextWrapWord

		// Geoloc Button Example
		geoBtn := widget.NewButton("View Concert Map Coordinates", func() {
			// This is where you call utils.GetCoordinates for their locations
			// And perhaps open a map URL or show a popup
			for loc := range a.Locations {
				lat, lon, _ := utils.GetCoordinates(loc)
				fmt.Printf("Location: %s -> %s, %s\n", loc, lat, lon)
			}
		})

		detailsContainer.Add(title)
		detailsContainer.Add(info)
		detailsContainer.Add(membersLabel)
		detailsContainer.Add(geoBtn)
		detailsContainer.Refresh()
	}

	// --- LEFT SIDE: Search & List ---

	// 1. Data Source for List
	filteredData := data
	list := widget.NewList(
		func() int { return len(filteredData) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(filteredData[i].Name)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		updateDetails(filteredData[id])
	}

	// 2. Search Bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search artist, member, location...")
	searchEntry.OnChanged = func(s string) {
		// Logic to filter the 'filteredData' slice based on 's'
		// If s is empty, reset filteredData to 'data'
		// Then call list.Refresh()

		// Note: For the "Suggestions" requirement, you might want to use
		// a PopUpMenu or a separate list that appears under the search bar.
	}

	// Layout
	leftSide := container.NewBorder(searchEntry, nil, nil, nil, list)
	split := container.NewHSplit(leftSide, container.NewVScroll(detailsContainer))
	split.SetOffset(0.3) // List takes 30% width

	myWindow.SetContent(split)
	myWindow.ShowAndRun()
}
