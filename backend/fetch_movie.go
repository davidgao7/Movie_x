package main

import (
	"fmt"
	"io"
	"net/http"
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

}

func main() {
	// fetch movie info from IMDb
	fetch_from_web()
}
