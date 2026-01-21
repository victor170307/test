package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GeoResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// GetCoordinates converts "city, country" to Lat/Lon
func GetCoordinates(location string) (string, string, error) {
	// Clean the location string (remove underscores usually found in the API data)
	// e.g., "new_york-usa" -> "new york usa"
	cleanLoc := strings.ReplaceAll(location, "_", " ")
	cleanLoc = strings.ReplaceAll(cleanLoc, "-", " ")

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", cleanLoc)

	client := http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "GroupieTracker/1.0") // Nominatim requires a User-Agent

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var res []GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", err
	}

	if len(res) > 0 {
		return res[0].Lat, res[0].Lon, nil
	}
	return "", "", fmt.Errorf("no coordinates found")
}
