package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
}

func NewProductRepo(fn string) *productRepo {
	return &productRepo{fileName: fn}
}

// Createproduct method creates new product with given name, price, category_id and returns its
func (b *productRepo) CreateProduct(req models.CreateProduct) (string, error) {

	productes, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	productes = append(productes, models.Product{
		Id:         id,
		Name:       req.Name,
		Price:      req.Price,
		CategoryId: req.CategoryId,
	})

	err = b.write(productes)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *productRepo) UpdateProduct(req models.Product) (msg string, err error) {

	productes, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range productes {
		if v.Id == req.Id {
			productes[i] = req
			msg = "updated successfully"
			err = b.write(productes)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *productRepo) GetProduct(req models.IdRequest) (resp models.Product, err error) {
	productes, err := b.read()
	if err != nil {
		return models.Product{}, err
	}
	for _, v := range productes {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Product{}, errors.New("not found")
}
func (b *productRepo) GetAllProduct(req models.GetAllProductRequest) (resp models.GetAllProduct, err error) {
	productes, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(productes) {
		resp.Products = []models.Product{}
		resp.Count = len(productes)
		return resp, nil
	} else if end > len(productes) {
		return models.GetAllProduct{
			Products: productes[start:],
			Count:    len(productes),
		}, nil
	}

	return models.GetAllProduct{
		Products: productes[start:end],
		Count:    len(productes)}, nil
}
func (b *productRepo) DeleteProduct(req models.IdRequest) (string, error) {
	productes, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range productes {
		if v.Id == req.Id {
			if i == (len(productes) - 1) {
				productes = productes[:i]
			} else {
				productes = append(productes[:i], productes[i+1:]...)
			}
			err = b.write(productes)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")

}

func (u *productRepo) read() ([]models.Product, error) {
	var (
		productes []models.Product
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &productes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return productes, nil
}

func (u *productRepo) write(productes []models.Product) error {

	body, err := json.Marshal(productes)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
