package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strings"
)

func main() {
	// fetch movie info from IMDb
	fetch_from_web()
}

// store movie info in a struct
type Movie struct {
	title  string
	year   string
	length unit32
	genre  *string
	stars  *string
}

// get HTTP requests to IMDb's servers go get github.com/PuerkitoBio/goquery
func fetch_from_web() {

	// Create a new collector
	c := colly.NewCollector(
		// allow only domains: www.imdb.com
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
		// set the recursion depth for links to visit, goes forever if not set
		colly.MaxDepth(3),
		// enables asynchronous network requests
		colly.Async(true),
	)
	// visit only domains: www.imdb.com

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		log.Printf("visiting", link)
		// Visit link found on page
		// Only those links are visited which are in the same domain
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://www.imdb.com
	c.Visit(web_link)

}
