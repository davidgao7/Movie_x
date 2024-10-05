// dependencies
// go mod download github.com/gocolly/colly/v2
// go build will automatically install these dependencies
// install scraper
module backend

go 1.14

require (
	github.com/gocolly/colly/v2 v2.1.0
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
)
