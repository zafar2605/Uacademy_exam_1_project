package models

type Product struct {
	Id         string  `json:"id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}

type CreateProduct struct {
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}

type IdRequestProduct struct {
	Id string
}

type GetAllProductRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllProduct struct {
	Count    int
	Products []Product
}
