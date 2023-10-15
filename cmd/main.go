package main

import (
	"playground/newProject/config"
	"playground/newProject/handler"
	"playground/newProject/models"
	"playground/newProject/storage/memory"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage(models.FileNames{
		BranchFile:                   "data/branches.json",
		UserFile:                     "data/users.json",
		CategoryFile:                 "data/categories.json",
		ProductFile:                  "data/products.json",
		BranchProductFile:            "data/branch_products.json",
		BranchProductTransactionFile: "data/branch_pr_transaction.json",
	})

	con := handler.NewHandler(strg, *cfg)

	//	Tasks answer key con.Task1, con.Task2...
	//	Tasks answer code is in handler/tasks.go

	con.Task11()

}
