package models

type CreateUser struct {
	Name string `json:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type IdRequestUser struct {
	Id string
}

type GetAllUserRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllUser struct {
	Users []User
	Count int
}
