package storage

import "playground/newProject/models"

type StorageI interface {
	Branch() BranchesI
	BranchProduct() BranchesProductsI
	BranchProductTransaction() BranchProductTransactionI
	User() UsersI
	Category() CategoryI
	Product() ProductI
}
type CategoryI interface {
	//CreateProduct method creates new branch with given name and address and returns its id
	CreateCategory(models.CreateCategory) (string, error)
	UpdateCategory(models.Category) (string, error)
	GetCategory(models.IdRequest) (models.Category, error)
	GetAllCategory(models.GetAllCategoryRequest) (models.GetAllCategory, error)
	DeleteCategory(models.IdRequest) (string, error)
}

type ProductI interface {
	//CreateCategory method creates new branch with given name and address and returns its id
	CreateProduct(models.CreateProduct) (string, error)
	UpdateProduct(models.Product) (string, error)
	GetProduct(models.IdRequest) (models.Product, error)
	GetAllProduct(models.GetAllProductRequest) (models.GetAllProduct, error)
	DeleteProduct(models.IdRequest) (string, error)
}

type BranchesI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}
type UsersI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateUser(models.CreateUser) (string, error)
	UpdateUser(models.User) (string, error)
	GetUser(models.IdRequest) (models.User, error)
	GetAllUser(models.GetAllUserRequest) (models.GetAllUser, error)
	DeleteUser(models.IdRequest) (string, error)
	GetListUser() (resp models.GetAllUser, err error)
}
type BranchesProductsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateBranchProduct(models.CreateBranchProduct) (string, error)
	UpdateBranchProduct(models.BranchProduct) (string, error)
	GetBranchProduct(models.IdRequest) (models.BranchProduct, error)
	GetAllBranchProduct(models.GetAllBranchProductRequest) (models.GetAllBranchProduct, error)
	DeleteBranchProduct(models.IdRequest) (string, error)
	GetListBranchProduct() (resp models.GetAllBranchProduct, err error)
}
type BranchProductTransactionI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateBranchProductTransaction(models.CreateBranchProductTransaction) (string, error)
	UpdateBranchProductTransaction(models.BranchProductTransaction) (string, error)
	GetBranchProductTransaction(models.IdRequest) (models.BranchProductTransaction, error)
	GetAllBranchProductTransaction(models.GetAllBranchProductTransactionRequest) (models.GetAllBranchProductTransaction, error)
	DeleteBranchProductTransaction(models.IdRequest) (string, error)
	GetListBranchProductTransaction() (resp models.GetAllBranchProductTransaction, err error)
}
