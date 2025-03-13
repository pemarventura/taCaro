package models

// RequestData represents the payload for processarQRCode endpoint.
// swagger:model
type RequestData struct {
	// URL to be fetched and processed.
	URL string `json:"url"`
}
