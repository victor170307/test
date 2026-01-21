package main

import (
	"groupie-tracker/api"
	"groupie-tracker/ui"
	"log"
)

func main() {
	// 1. Fetch Data
	log.Println("Fetching data...")
	artists, err := api.FetchAllData()
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}

	// 2. Start GUI
	log.Println("Starting GUI...")
	ui.StartApp(artists)
}
