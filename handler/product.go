package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateProduct(name, categoryId string, price float64) {
	resp, err := h.strg.Product().CreateProduct(models.CreateProduct{
		Name:       name,
		CategoryId: categoryId,
		Price:      price,
	})
	if err != nil {
		fmt.Println("error from CreateProduct:", err.Error())
		return
	}
	fmt.Println("created new Product with id:", resp)
}
func (h *handler) UpdateProduct(id, name, categoryId string, price float64) {
	resp, err := h.strg.Product().UpdateProduct(models.Product{
		Id:         id,
		Name:       name,
		CategoryId: categoryId,
		Price:      price,
	})
	if err != nil {
		fmt.Println("error from UpdateProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetProduct(id string) models.Product {
	resp, err := h.strg.Product().GetProduct(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetProduct:", err.Error())
		return models.Product{}
	}
	return resp
}

func (h *handler) GetAllProduct(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Product().GetAllProduct(models.GetAllProductRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error from GetAllProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteProduct(id string) {
	resp, err := h.strg.Product().DeleteProduct(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}
