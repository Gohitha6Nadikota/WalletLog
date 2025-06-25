package graph

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"backend/graph/model"
)

const dataFile = "expenses.json"

func loadExpenses() []*model.Expense {
	file, err := os.Open(dataFile)
	if err != nil {
		return []*model.Expense{}
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	var expenses []*model.Expense
	json.Unmarshal(bytes, &expenses)

	return expenses
}

