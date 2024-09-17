package main

import (
	"backend/pkg/fetcher" // import the fetcher package locally
	"fmt"
	"log"
)

func main() {
	// Path to the TSV file (adjust this path to where your TSV file is located)
	tsvFilePath := "../../../data/title.basics.tsv"

	// Number of workers for concurrent processing
	numWorkers := 4

	// Call FetchFromCSV to process the TSV file
	movies, err := fetcher.FetchFromCSV(tsvFilePath, numWorkers)
	if err != nil {
		log.Fatalf("Error fetching data from TSV: %v", err)
	}

	// Print out the fetched movies for testing
	for _, movie := range movies {
		fmt.Printf("Title: %s, Year: %s, Length: %d, Rate Level: %s, Review: %.2f\n",
			movie.Title, movie.Year, movie.Length, movie.RateLevel, movie.Review)

		// Optional: Print additional fields if they exist
		if movie.Genre != nil {
			fmt.Printf("Genre: %s\n", *movie.Genre)
		}
		if movie.Stars != nil {
			fmt.Printf("Stars: %s\n", *movie.Stars)
		}
		fmt.Println("-----------")
	}
}
