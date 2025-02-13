package entity

type Request struct {
	ItemsNumber int `json:"items_number" validate:"required,min=1"`
}

type PackResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Count  int    `json:"count"`
}

type Response struct {
	ItemsNumber int            `json:"items_number"`
	TotalItems  int            `json:"total_items"`
	Packs       []PackResponse `json:"packs"`
}

// CalculateTotalItems calculates the total number of items in all packs
func (r *Response) CalculateTotalItems() {
	total := 0
	for _, pack := range r.Packs {
		total += pack.Amount * pack.Count
	}
	r.TotalItems = total
}
