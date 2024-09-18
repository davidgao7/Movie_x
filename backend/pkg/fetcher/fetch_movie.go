package fetcher

import (
	"backend/pkg/model"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// FetchFromCSV reads a CSV file and processes each line concurrently.
func FetchFromCSV(filePath string, numWorkers int) ([]model.Movie, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Channel to send lines to workers
	linesChan := make(chan string, 100)
	moviesChan := make(chan model.Movie, 1000) // Channel to collect processed movies, Increase the buffer size to prevent blocking
	var wg sync.WaitGroup                      // WaitGroup to wait for all workers to finish

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range linesChan {
				movie, err := ProcessLine(line)
				if err == nil { // Only send valid movies
					moviesChan <- movie
				}
			}
		}()
	}

	// Read file line by line and send to workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	close(linesChan) // No more lines to process

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(moviesChan) // No more movies to collect
	}()

	// Collect all movies
	var movies []model.Movie
	for movie := range moviesChan {
		movies = append(movies, movie)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return movies, nil
}

// ProcessLine converts a line of CSV data into a Movie struct.
func ProcessLine(line string) (model.Movie, error) {
	// Split the line into fields (assuming comma-separated values)
	fields := strings.Split(line, ",")
	if len(fields) < 5 { // Adjust this based on the expected number of fields
		return model.Movie{}, fmt.Errorf("invalid line format: %s", line)
	}

	// Example parsing logic (adjust as necessary to match your CSV format)
	fmt.Println(fields)
	fmt.Println(len(fields))

	var movie model.Movie
	if len(fields) == 7 {
		movie = model.Movie{
			Title:     fields[0],
			Year:      fields[1],
			Length:    parseUint32(fields[2]),
			RateLevel: fields[3],
			Review:    parseFloat32(fields[4]),
			Genre:     &fields[5], // Assume this field is optional
			Stars:     &fields[6], // Assume this field is optional
		}
	} else if len(fields) == 5 {
		movie = model.Movie{
			Title:     fields[0],
			Year:      fields[1],
			Length:    parseUint32(fields[2]),
			RateLevel: fields[3],
			Review:    parseFloat32(fields[4]),
		}
	}

	return movie, nil
}

// Helper functions for parsing (replace with actual parsing logic)
func parseUint32(input string) uint32 {
	// Add proper error handling as needed
	var result uint32
	fmt.Sscanf(input, "%d", &result)
	return result
}

func parseFloat32(input string) float32 {
	// Add proper error handling as needed
	var result float32
	fmt.Sscanf(input, "%f", &result)
	return result
}
