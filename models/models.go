package models

// Artist represents the structure of the artist API
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// We will manually add this data after fetching Relations
	Locations map[string][]string
	Dates     []string
}

// Relation represents the structure of the relation API
type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// SearchResult is used for the Search Bar suggestions
type SearchResult struct {
	Text     string
	Type     string // e.g., "artist", "member", "location"
	ArtistID int    // To link back to the specific artist
}
