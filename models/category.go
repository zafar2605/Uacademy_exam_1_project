package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type IdRequestCategory struct {
	Id string
}

type GetAllCategoryRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllCategory struct {
	Count      int
	Categories []Category
}
