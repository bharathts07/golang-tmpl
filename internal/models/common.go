package models

// Image holds information for a single image component in view
type Image struct {
	Title string  `json:"title"`
	URL   string  `json:"url"`
	Caption string  `json:"caption"`
}
