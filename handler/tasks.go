package handler

import (
	"fmt"
	"log"
	"playground/newProject/models"
	"sort"
	"strconv"
)

type plusMinusSum struct {
	Plus  int
	Minus int
}

type numberAndSum struct {
	Number int
	Sum    int
}

type NumberSumPlusMinus struct {
	PlusNumber  int
	MinusNumber int
	PlusSum     int
	MinusSum    int
}

func (h *handler) Task1() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}
	var mp = make(map[string]int)
	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{
			Id: val.BranchID,
		})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		mp[branch.Id]++
	}
	for ind, val := range mp {
		slc = append(slc, []string{ind, strconv.Itoa(val)})
	}
	sort.Slice(slc, func(a, b int) bool {
		return slc[a][1] > slc[b][1]
	})
	for _, val := range slc {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{Id: val[0]})
		if err != nil {
			log.Fatal("Error read data:", err)
		}
		fmt.Println("Branch:", branch.Name, "-->  Count:", val[1])
	}
}

func (h *handler) Task2() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}
	var mp = make(map[string]int)
	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {

		branch, err := h.strg.Branch().GetBranch(models.IdRequest{Id: val.BranchID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		price, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		mp[branch.Id] += int(price.Price) * val.Quantity
	}
	for key, val := range mp {
		slc = append(slc, []string{key, strconv.Itoa(val)})
	}
	sort.Slice(slc, func(a, b int) bool {
		index_a, _ := strconv.Atoi(slc[a][1])
		index_b, _ := strconv.Atoi(slc[b][1])
		return index_a > index_b
	})
	for _, val := range slc {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{Id: val[0]})
		if err != nil {
			log.Fatal("Error read data:", err)
		}
		fmt.Println("Branch: ", branch.Name, "-->  Sum:", val[1])
	}
}

func (h *handler) Task3() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}

	var mp = make(map[string]int)
	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {

		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		mp[product.Id]++
	}

	for key, val := range mp {
		slc = append(slc, []string{key, strconv.Itoa(val)})
	}
	sort.Slice(slc, func(a, b int) bool {
		index_a, _ := strconv.Atoi(slc[a][1])
		index_b, _ := strconv.Atoi(slc[b][1])
		return index_a > index_b
	})
	for _, val := range slc {
		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val[0]})
		if err != nil {
			log.Fatal("Error read data:", err)
		}
		fmt.Println("Product: ", product.Name, "-->  Count:", val[1])
	}
}

func (h *handler) Task4() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data", err)
	}

	var mp = make(map[string]int)
	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {

		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data", err)
		}
		category, err := h.strg.Category().GetCategory(models.IdRequest{Id: product.CategoryId})
		if err != nil {
			log.Fatal("Error read data", err)
		}
		mp[category.Id]++
	}

	for key, val := range mp {
		slc = append(slc, []string{key, strconv.Itoa(val)})
	}
	sort.Slice(slc, func(a, b int) bool {
		index_a, _ := strconv.Atoi(slc[a][1])
		index_b, _ := strconv.Atoi(slc[b][1])
		return index_a > index_b
	})
	for _, val := range slc {
		category, err := h.strg.Category().GetCategory(models.IdRequest{Id: val[0]})
		if err != nil {
			log.Fatal("Error read data:", err)
		}
		fmt.Println("Category: ", category.Name, "-->  Count:", val[1])
	}
}

func (h *handler) Task5() {
	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}
	var res = make(map[string][]string)
	var res1 = make(map[string]int)
	result := make(map[string][]map[string]int)

	for _, v := range resp.BranchProductTransactions {
		res[v.BranchID] = append(res[v.BranchID], v.ProductID)
	}

	for branch_id, products := range res {

		res1 = map[string]int{}
		for _, v := range products {
			product, err := h.strg.Product().GetProduct(models.IdRequest{
				Id: v,
			})
			if err != nil {
				log.Fatal("Error read data: ", err)
			}
			category, err := h.strg.Category().GetCategory(models.IdRequest{
				Id: product.CategoryId,
			})

			if err != nil {
				log.Fatal("Error read data: ", err)
			}
			res1[category.Name]++
		}
		result[branch_id] = append(result[branch_id], res1)
	}
	for key, val := range result {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{
			Id: key,
		})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		fmt.Println("Branch:", branch.Name, val)
	}
}

func (h *handler) Task6() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}
	var mp = make(map[string]NumberSumPlusMinus)

	for _, val := range resp.BranchProductTransactions {

		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data", err)
		}
		plus_num := mp[val.BranchID].PlusNumber
		minus_num := mp[val.BranchID].MinusNumber
		plus_sum := mp[val.BranchID].PlusSum
		minus_sum := mp[val.BranchID].MinusSum
		if val.Type == "plus" {
			plus_num += 1
			plus_sum += val.Quantity * int(product.Price)
		} else {
			minus_num += 1
			minus_sum += val.Quantity * int(product.Price)
		}

		mp[val.BranchID] = NumberSumPlusMinus{
			PlusNumber:  plus_num,
			MinusNumber: minus_num,
			PlusSum:     plus_sum,
			MinusSum:    minus_sum,
		}
	}
	for key, val := range mp {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{Id: key})
		if err != nil {
			log.Fatal("Error read data:", err)
		}
		fmt.Println(branch.Name, "-->", " PlusNumber:", val.PlusNumber, "  MinusNumber:", val.MinusNumber, "  PlusSum:", val.PlusSum, "  MinusSum:", val.MinusSum)
	}
}

func (h *handler) Task7() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}

	var mp = make(map[string]int)
	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		if val.Type == "plus" {
			mp[val.CreatedAt] += val.Quantity
		}
	}
	for key, val := range mp {
		slc = append(slc, []string{key, strconv.Itoa(val)})
	}

	sort.Slice(slc, func(a, b int) bool {
		index_a, _ := strconv.Atoi(slc[a][1])
		index_b, _ := strconv.Atoi(slc[b][1])
		return index_a > index_b
	})

	for _, val := range slc {
		fmt.Println("Date: ", val[0], "--> Count:", val[1])
	}
}

func (h *handler) Task8() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}

	var mp = make(map[string][]string)
	var mp_plus = make(map[string]int)
	var mp_minus = make(map[string]int)

	var slc = [][]string{}

	for _, val := range resp.BranchProductTransactions {

		if val.Type == "plus" {
			mp_plus[val.ProductID]++
		} else {
			mp_minus[val.ProductID]++
		}
		mp[val.ProductID] = []string{strconv.Itoa(mp_plus[val.ProductID]), strconv.Itoa(mp_minus[val.ProductID])}
	}

	for ind, val := range mp {
		slc = append(slc, []string{ind, val[0], val[1]})
	}

	for _, val := range slc {
		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val[0]})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		fmt.Println("Product: ", product.Name, "--> Kirdi:", val[1], "--> Chiqdi:", val[2])
	}
}

func (h *handler) Task9() {

	resp, err := h.strg.BranchProduct().GetListBranchProduct()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}

	var mp = make(map[string]int)

	for _, val := range resp.BranchProducts {

		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		price := int(product.Price) * val.Quantity
		mp[val.BranchID] += price
	}

	for key, val := range mp {
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{Id: key})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		fmt.Println("Branch:", branch.Name, "--> Sum:", val)
	}
}

func (h *handler) Task10() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data: ", err)
	}

	var mp = make(map[string]int)

	for _, val := range resp.BranchProductTransactions {

		product, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		mp[val.UserID] += int(product.Price) * val.Quantity
	}

	for key, val := range mp {
		user, err := h.strg.User().GetUser(models.IdRequest{Id: key})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		fmt.Println("User:", user.Name, "--> Sum:", val)
	}
}

func (h *handler) Task11() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data", err)
	}
	var mp = make(map[[2]string]numberAndSum)

	for _, val := range resp.BranchProductTransactions {

		price, err := h.strg.Product().GetProduct(models.IdRequest{Id: val.ProductID})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		name_date := [2]string{val.UserID, val.CreatedAt}
		number := mp[name_date].Number
		sum := mp[name_date].Sum

		mp[name_date] = numberAndSum{
			Number: number + 1,
			Sum:    sum + (val.Quantity * int(price.Price)),
		}
	}
	for key, val := range mp {
		user, err := h.strg.User().GetUser(models.IdRequest{Id: key[0]})
		if err != nil {
			log.Fatal("Error read data: ", err)
		}
		fmt.Println(user.Name,"--> Day: ", key[1], "--> Count:", val.Number, "--> Sum:", val.Sum)
	}
}

func (h *handler) Task12() {

	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Fatal("Error read data", err)
	}

	var mp = make(map[string]plusMinusSum)

	for _, val := range resp.BranchProductTransactions {
		var plus, minus = 0, 0
		if val.Type == "plus" {
			plus = mp[val.UserID].Plus + 1
		} else {
			minus = mp[val.UserID].Minus + 1
		}

		mp[val.UserID] = plusMinusSum{
			Plus:  plus,
			Minus: minus,
		}
	}
	for key, val := range mp {
		user, err := h.strg.User().GetUser(models.IdRequest{Id: key})
		if err != nil {
			log.Fatal("Error read data", err)
		}
		fmt.Println(user.Name, ":>  Kiritgan:", val.Plus, " -->  Chiqargan:", val.Minus)
	}
}
