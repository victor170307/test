package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

// fetchArtists fetches the list of artists from the API
func fetchArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// FetchAllData fetches artists and merges them with their relations
func FetchAllData() ([]models.Artist, error) {
	// 1. Fetch Artists
	artists, err := fetchArtists()
	if err != nil {
		return nil, err
	}

	// 2. Fetch Relations
	resp, err := http.Get(relationsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relData models.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relData); err != nil {
		return nil, err
	}

	// 3. Merge Data (Link relations to artists by ID)
	// Note: The APIs usually align by ID index (0 to n)
	for i := range artists {
		for _, rel := range relData.Index {
			if rel.ID == artists[i].ID {
				artists[i].Locations = rel.DatesLocations
				break
			}
		}
	}

	return artists, nil
}
