package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadFile(filePath string, ch chan Receipent) error {

	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Cannot Open File")
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println("Unable to Record records")
	}

	defer f.Close()

	for _, record := range records[1:] {
		ch <- Receipent{
			Name:  record[0],
			Email: record[1],
		}
	}
	return err

}
