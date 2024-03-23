package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
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
// get HTTP requests to IMDb's servers go get github.com/PuerkitoBio/goquery
func fetch_from_web() {
	c := colly.NewCollector(
         colly.AllowedDomains("imdb.com", "www.imdb.com")
        )
    // create a info collector 
  infoCollector := c.Clone()

	movies := []Movie{}

	c.OnHTML(".mode-detail", func(e *colly.HTMLElement) {
		fmt.Println("e: ", e)
    profileUrl := e.ChildAttr("div.lister-item-image > a", "href")
    profileUrl = e.Request.AbsoluteURL(profileUrl)
    infoCollector.Visit(profileUrl)

    // go to next page
    c.OnHTML("a.lister-page-next", func(e *colly.HTMLElement){
      nextPage :=e.Request.AbsoluteURL(e.Attr("href"))
    c.Visit(nextPage)
    })

    infoCollector.OnHTML("#content-2-wide", func(e *colly.HTMLElement){
            // create movie object
            movie := Movie{}
            // TODO: get css elements
            movie.title = e.ChildText("h1.header > span.itemprop")
            fmt.Println( "title: ", movie.title)

            movie.year = e.ChildText("h1.header > a")
            fmt.Println("year: ", movie.year)]

            movie.length = e.ChildText("time")
            fmt.Println("length: ", movie.length)

            movie.rate_level = e.ChildText("span[itemprop='ratingValue']")
            fmt.Println("rate_level: ", movie.rate_level)

            movie.review = e.ChildText("span[itemprop='ratingCount']")
            fmt.Println("review: ", movie.review)

            movie.genre = e.ChildText("span[itemprop='genre']")
            fmt.Println("genre: ", movie.genre)

            movie.stars = e.ChildText("span[itemprop='actors']")
            fmt.Println("stars: ", movie.stars)

            // append movie to movies
            movies = append(movies, movie)
            fmt.Println("append movie successfully")
        })
	})

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    infoCollector.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    }
}}

func main() {
	// fetch movie info from IMDb
	fetch_from_web()
}
