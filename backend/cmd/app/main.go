package main

import (
	// import the db package
	// import the fetcher package locally
	// Import the db package
	"backend/internal/db"
	"backend/pkg/fetcher"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	// // test redis connection
	// // Define Redis connection URL
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// redisURL := os.Getenv("REDIS_URL")
	//
	// // Get Redis client
	// redisClient, _ := db.GetRedisClient(redisURL)
	//
	// fmt.Println(redisClient)
	//
	// // TEST: set a key-value pair
	// ctx := context.Background()
	//
	// err = redisClient.Set(ctx, "testKey", "testValue", 0).Err()
	// if err != nil {
	// 	log.Fatalf("Failed to set key: %v", err)
	// }
	// // TEST: get the value of the key
	// val, err := redisClient.Get(ctx, "testKey").Result()
	// if err != nil {
	// 	log.Fatalf("Failed to get key: %v", err)
	// }
	// fmt.Println("testKey:", val)
	//
	// defer redisClient.Close()
	//
	// fmt.Println("===============Connected to Redis===============")
	//

	// Get PostgreSQL client
	postgreURL := os.Getenv("XATA_POSTGRE_SQL_ENDPOINT")
	dbClient, err := db.GetPostgresClient(postgreURL)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	fmt.Println("===============Connected to PostgreSQL===============")

	// initialize the database and tables
	fmt.Println("===============Initializing the database===============")

	fmt.Println("===============check if table exists===============")
	var exists int
	db_name := os.Getenv("DB_NAME")
	query_check_table := fmt.Sprintf("SELECT COUNT(id) as movies_num FROM %s.movies", pq.QuoteIdentifier(db_name))
	row := dbClient.QueryRow(query_check_table)
	err = row.Scan(&exists)
	fmt.Println("error", err)
	fmt.Println("===============check if table exists ends===============")

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P01" {
			// 42P01 is the PostgreSQL error code for "relation does not exist"
			fmt.Println("Table does not exist, creating table")
			_, err = dbClient.Exec(`
            CREATE TABLE IF NOT EXISTS movies (
                id SERIAL PRIMARY KEY,
                title VARCHAR(255) NOT NULL,
                year VARCHAR(4) NOT NULL,
                length INT,
                rate_level VARCHAR(50),
                review REAL,
                genre VARCHAR(255),
                stars VARCHAR(255)
            );
        `)
			if err != nil {
				log.Fatalf("Failed to create table: %v", err)
			}
		} else {
			log.Fatalf("Failed to check if table exists: %v", err)
		}
		fmt.Println("===========Table created=================")
	} else {
		fmt.Println("Table exists, skipping table creation")
	}

	fmt.Println("===============Initializing the database success!===============")
	os.Exit(0)

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

		// populate the movie into the database
	}
	defer dbClient.Close()
}
