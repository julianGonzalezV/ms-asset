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
	FloorLevel         float64             `json:"floorLevel"`
	Elevator           bool                `json:"elevator"`
	VisitorParking     bool                `json:"visitorParking"`
	CommunalArea       bool                `json:"communalArea"`
	Gym                bool                `json:"gym"`
	Type               string              `json:"type,omitempty"`
	Images             []AssetImageRequest `json:"images,omitempty"`
}

type AssetImageRequest struct {
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
}
