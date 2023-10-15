package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type categoryRepo struct {
	fileName string
}

func NewCategoryRepo(fn string) *categoryRepo {
	return &categoryRepo{fileName: fn}
}

//CreateCategory method creates new Category with given name and address and returns its id
func (b *categoryRepo) CreateCategory(req models.CreateCategory) (string, error) {

	Categoryes, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	Categoryes = append(Categoryes, models.Category{
		Id:      id,
		Name:    req.Name,
	})

	err = b.write(Categoryes)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *categoryRepo) UpdateCategory(req models.Category) (msg string, err error) {

	Categoryes, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range Categoryes {
		if v.Id == req.Id {
			Categoryes[i] = req
			msg = "updated successfully"
			err = b.write(Categoryes)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *categoryRepo) GetCategory(req models.IdRequest) (resp models.Category, err error) {
	Categoryes, err := b.read()
	if err != nil {
		return models.Category{}, err
	}
	for _, v := range Categoryes {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Category{}, errors.New("not found")
}
func (b *categoryRepo) GetAllCategory(req models.GetAllCategoryRequest) (resp models.GetAllCategory, err error) {
	Categoryes, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(Categoryes) {
		resp.Categories = []models.Category{}
		resp.Count = len(Categoryes)
		return resp, nil
	} else if end > len(Categoryes) {
		return models.GetAllCategory{
			Categories: Categoryes[start:],
			Count:    len(Categoryes),
		}, nil
	}

	return models.GetAllCategory{
		Categories: Categoryes[start:end],
		Count:    len(Categoryes)}, nil
}
func (b *categoryRepo) DeleteCategory(req models.IdRequest) (string, error) {
	Categoryes, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range Categoryes {
		if v.Id == req.Id {
			if i == (len(Categoryes) - 1) {
				Categoryes = Categoryes[:i]
			} else {
				Categoryes = append(Categoryes[:i], Categoryes[i+1:]...)
			}
			err = b.write(Categoryes)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")

}

func (u *categoryRepo) read() ([]models.Category, error) {
	var (
		Categoryes []models.Category
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Categoryes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return Categoryes, nil
}

func (u *categoryRepo) write(Categoryes []models.Category) error {

	body, err := json.Marshal(Categoryes)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
