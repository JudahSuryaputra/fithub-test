package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Transaction struct {
	No                  int    `json:"no"`
	LocationTransaction string `json:"locationTransaction"`
	TotalAmount         int    `json:"totalAmount"`
}

type GrouppedTransaction struct {
	LocationTransaction string `json:"locationTransaction"`
	TotalAmount         int    `json:"totalAmount"`
}

func main() {
	var transactions []Transaction

	file, err := ioutil.ReadFile("transactions.json")
	if err != nil {
		log.Fatalf("cannot read json file")
	}

	err = json.Unmarshal([]byte(file), &transactions)
	if err != nil {
		log.Fatalf("cannot unmarshal json")
	}

	data := make(map[string]int)

	for _, transaction := range transactions {
		if _, isExist := data[transaction.LocationTransaction]; !isExist {
			data[transaction.LocationTransaction] = transaction.TotalAmount
		} else {
			data[transaction.LocationTransaction] = data[transaction.LocationTransaction] + transaction.TotalAmount
		}
	}

	// pls sort
	var grouppedTransactions []GrouppedTransaction

	for key, value := range data {
		transaction := GrouppedTransaction{
			LocationTransaction: key,
			TotalAmount:         value,
		}

		grouppedTransactions = append(grouppedTransactions, transaction)
	}

	sort.Slice(grouppedTransactions, func(i, j int) bool {
		return grouppedTransactions[i].TotalAmount > grouppedTransactions[j].TotalAmount
	})

	fmt.Println(grouppedTransactions)
	fmt.Printf("%+v\n", grouppedTransactions)
}
