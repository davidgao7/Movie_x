package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "io"
	// "net/http"
	"reflect"
	"strings"
)

// store movie info in a struct
type Movie struct {
	title      string // lower case can only be accessed within the same package
	year       string
	length     uint32
	rate_level string
	review     float32
	genre      *string
	stars      *string
}

// TODO: implement multi threading for fast fetching later
func fetch_from_csv() {
	file, err := os.Open("../data/title.basics.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	csvReader := csv.NewReader(nil)

	// Read the first line (column names)
	if scanner.Scan() {
		csvReader = csv.NewReader(strings.NewReader(scanner.Text()))
		csvReader.Comma = '\t'
		columnNames, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Column names:", columnNames)
		fmt.Println("type of columnNames: ", reflect.TypeOf(columnNames[0]))
	}

	// Read the rest of the file
	for scanner.Scan() {
		csvReader = csv.NewReader(strings.NewReader(scanner.Text()))
		csvReader.Comma = '\t'
		record, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("=====================")
		fmt.Println(record)
		fmt.Println("type of record: ", reflect.TypeOf(record[0]))
		fmt.Println("=====================")

		// TODO: deal with one line first
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	// fetch movie info from IMDb
	fetch_from_csv()
}
