package dto

type Merchant struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}

type Offer struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    int        `json:"category"`
	Merchants   []Merchant `json:"merchants"`
	ValidTo     string     `json:"valid_to"`
}

type OffersData struct {
	Offers []Offer `json:"offers"`
}
