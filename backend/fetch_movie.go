package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	// "io"
	// "net/http"
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

// implement multi threading for fast fetching
func fetch_from_csv(filePath string, numNumWorkers int) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create channels for worker communication
	linesChan := make(chan string, 100) // buffer size 100
	var wg sync.WaitGroup               // wait for all workers to finish

	// start workers goroutines
	for i := 0; i < numNumWorkers; i++ {
		wg.Add(1)
		go processLine(linesChan, &wg)
	}

	// read file line by line and send to workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	close(linesChan) // no more lines to read

	// wait for all workers to finish
	wg.Wait()

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func processLine(linesChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range linesChan {
		fmt.Println(line)
		os.Exit(0)

	}
}

func main() {
	// fetch movie info from IMDb
	fetch_from_csv("../data/title.basics.tsv", 10)
}
