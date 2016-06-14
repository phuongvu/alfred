package main

type Target struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title       string    `json:"title"`
	ProductType string    `json:"product_type"`
	Variants    []Variant `json:"variants"`
}

type Variant struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Grams int    `json:"grams"`
}

type Cart struct {
	TotalWeight string
	TotalPrice  string
	items       []Variant
}
