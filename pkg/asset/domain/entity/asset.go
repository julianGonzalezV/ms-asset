package entity

// New function is used to create a new struct
func New(items []Item, price float64, businessId, sku, name, description, category, state, productType, image string) *Asset {
	return &Asset{
		BusinessId:  businessId,
		Sku:         sku,
		Image:       image,
		Name:        name,
		Description: description,
		Category:    category,
		State:       state,
		Price:       price,
		ProductType: productType,
		Items:       items,
	}
}

type Asset struct {
	internalCode string       `json:"sku"`
	Name         string       `json:"name"`
	Country      string       `json:"state,omitempty"`
	province     string       `json:"state,omitempty"`
	City         string       `json:"state,omitempty"`
	Description  string       `json:"description,omitempty"`
	Category     string       `json:"category,omitempty"`
	State        string       `json:"state,omitempty"`
	RentingPrice float64      `json:"rentingPrice"`
	Area         float64      `json:"area"`
	Rooms        float64      `json:"rooms"`
	BathRooms    float64      `json:"bathRooms"`
	Parkings     float64      `json:"bathRooms"`
	Furnished    bool         `json:"furnished"`
	ProductType  string       `json:"type,omitempty"`
	Images       []AssetImage `json:"images,omitempty"`
}

type AssetImage struct {
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
}
