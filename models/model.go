package models

type (
	// Session Holds info about a session
	Session map[string]interface{}

	// JSON Holds a JSON object
	JSON map[string]interface{}
)

// Response ... type
type Response struct {
	Info  SearchInfo `json:"searchInformation"`
	Items []Job      `json:"items"`
}

// SearchInfo ... type
type SearchInfo struct {
	Num string `json:"totalResults"`
}

// Job ... type
type Job struct {
	Title string    `json:"title"`
	Link  string    `json:"link"`
	Image Thumbnail `json:"pagemap"`
}

// Thumbnail ... type
type Thumbnail struct {
	CseImage []ImageSrc `json:"cse_image"`
}

// ImageSrc ... type
type ImageSrc struct {
	Src string `json:"src"`
}
