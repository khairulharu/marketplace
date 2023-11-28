package dto

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Category string `json:"category"`
}
