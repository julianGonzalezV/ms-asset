package request

type AssetRequest struct {
	Code               string              `json:"code"`
	RegistrationNumber string              `json:"registrationNumber"`
	Country            string              `json:"country,omitempty"`
	Province           string              `json:"province,omitempty"`
	City               string              `json:"city,omitempty"`
	Description        string              `json:"description,omitempty"`
	Category           string              `json:"category,omitempty"`
	State              string              `json:"state,omitempty"`
	RentingPrice       float64             `json:"rentingPrice"`
	Area               float64             `json:"area"`
	Rooms              float64             `json:"rooms"`
	BathRooms          float64             `json:"bathrooms"`
	Parkings           float64             `json:"parkings"`
	Furnished          bool                `json:"furnished"`
	Type               string              `json:"type,omitempty"`
	Images             []AssetImageRequest `json:"images,omitempty"`
}

type AssetImageRequest struct {
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
}
