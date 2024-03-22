package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
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

// get HTTP requests to IMDb's servers go get github.com/PuerkitoBio/goquery
func fetch_from_web() {
	c := colly.NewCollector()

	movies := []Movie{}

	c.OnHTML(".movie", func(e *colly.HTMLElement) {
		fmt.Println("e: ", e)
		movie := Movie{
			title:      e.ChildText(".name"),
			year:       e.ChildText(".release-date"),
			length:     e.ChildText(".length"),
			rate_level: e.ChildText(".rate-level"),
			genre:      e.ChildText(".genre"),
			stars:      e.ChildText(".stars"),
		}
		movies = append(movies, movie)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://example.com/movies")

	for _, movie := range movies {
		fmt.Printf("%+v\n", movie)
	}
}

func main() {
	// fetch movie info from IMDb
	fetch_from_web()
}
