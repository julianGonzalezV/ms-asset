package entity

// New function is used to create a new struct
func New(images []AssetImage, furnished bool, rentingPrice, area, rooms, bathRooms, parkings float64, country, province,
	city, description, category, state, pType, registrationNumber string) *Asset {
	return &Asset{
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
	BathRooms          float64      `json:"bathRooms"`
	Parkings           float64      `json:"parkings"`
	Furnished          bool         `json:"furnished"`
	Type               string       `json:"type,omitempty"`
	Images             []AssetImage `json:"images,omitempty"`
}

type AssetImage struct {
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
}
