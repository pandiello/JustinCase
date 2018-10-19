package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/pandiello/JustinCase/model"
)

func main() {

	records, err := ReadFile("caca.csv")
	fatalIf(err)

	headerOffset := 1
	var genericTransactions []model.GenericTransaction
	for _, row := range records[headerOffset:] {

		genericTransactions = append(genericTransactions, model.GenericTransaction{
			ID:          row[0],
			Name:        row[1],
			Description: row[2],
		})
	}

	fmt.Printf("%+v\n", genericTransactions)
}

// ReadFile reads a csv file using the given file path.
func ReadFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)
	return r.ReadAll()
}

func fatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
