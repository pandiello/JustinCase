package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pandiello/JustinCase/model"
)

func main() {

	records, err := readFile("fake.csv")
	fatalIf(err)

	headerOffset := 1
	var genericTransactions []model.GenericTransaction
	for _, row := range records[headerOffset:] {
		genericTransactions = append(genericTransactions, mapRow(row))
	}

	connectionString := "sqlserver://JustinCase:Perro@localhost"
	db, err := sql.Open("sqlserver", connectionString)
	for _, gt := range genericTransactions {
		_, err := db.Exec(`INSERT INTO GenericTransaction (ID,Name,Description,Date)
		values (@p1,@p2,@p3,@p4)`, gt.ID, gt.Name, gt.Description, time.Now())
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// ReadFile reads a csv file using the given file path.
func readFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)
	return r.ReadAll()
}

func mapRow(row []string) model.GenericTransaction {
	return model.GenericTransaction{
		ID:          row[0],
		Name:        row[1],
		Description: row[2],
	}
}

func fatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
