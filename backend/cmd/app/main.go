package main

import (
	// import the db package
	"backend/pkg/db"
	"backend/pkg/fetcher" // import the fetcher package locally
	"fmt"
	"log"
	"os"
)

func main() {
	// test redis connection
	redis_client, err := db.GetClient("redis://localhost:6379")
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer redis_client.Close()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
		return
	}
	fmt.Println("Current directory: ", dir) // Movie_x/backend

	// Path to the TSV file (adjust this path to where your TSV file is located)
	tsvFilePath := "../data/title.basics.tsv"

	// Number of workers for concurrent processing
	numWorkers := 4

	// Call FetchFromCSV to process the TSV file
	movies, err := fetcher.FetchFromCSV(tsvFilePath, numWorkers)
	if err != nil {
		log.Fatalf("Error fetching data from TSV: %v", err)
	}
	print(movies)
	os.Exit(0)

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
