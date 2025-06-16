package model

// Product define la estructura de un producto
type Product struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Price          float64  `json:"price"`
	PaymentMethods []string `json:"payment_methods"`
	Seller         string   `json:"seller"`
	Rating         float64  `json:"rating"`
	Stock          int      `json:"stock"`
	Images         []string `json:"images"`
}
