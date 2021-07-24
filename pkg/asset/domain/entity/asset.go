package entity

// New function is used to create a new struct
func New(images []AssetImage, furnished, visitorParking, elevator, communalArea, gym bool, floorLevel, rentingPrice, area, rooms, bathRooms, parkings float64, country, province,
	city, description, category, state, pType, code, registrationNumber string) *Asset {
	return &Asset{
		Code:               code,
		RegistrationNumber: registrationNumber,
		Country:            country,
		Province:           province,
		City:               city,
		Description:        description,
		Category:           category,
		State:              state,
		RentingPrice:       rentingPrice,
		Area:               area,
		Rooms:              rooms,
		BathRooms:          bathRooms,
		Parkings:           parkings,
		Furnished:          furnished,
		Elevator:           elevator,
		VisitorParking:     visitorParking,
		FloorLevel:         floorLevel,
		CommunalArea:       communalArea,
		Gym:                gym,
		Type:               pType,
		Images:             images,
	}
}

type Asset struct {
	Code               string       `json:"code"`
	RegistrationNumber string       `json:"RegistrationNumber"`
	Country            string       `json:"country,omitempty"`
	Province           string       `json:"province,omitempty"`
	City               string       `json:"city,omitempty"`
	Description        string       `json:"description,omitempty"`
	Category           string       `json:"category,omitempty"`
	State              string       `json:"state,omitempty"`
	RentingPrice       float64      `json:"rentingPrice"`
	Area               float64      `json:"area"`
	Rooms              float64      `json:"rooms"`
	BathRooms          float64      `json:"bathrooms"`
	Parkings           float64      `json:"parkings"`
	Furnished          bool         `json:"furnished"`
	FloorLevel         float64      `json:"floorLevel"`
	Elevator           bool         `json:"elevator"`
	VisitorParking     bool         `json:"visitorParking"`
	CommunalArea       bool         `json:"communalArea"`
	Gym                bool         `json:"gym"`
	Type               string       `json:"type,omitempty"`
	Images             []AssetImage `json:"images,omitempty"`
}

type AssetImage struct {
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
}
