package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	pages := flag.String("pages", "Lion,Barack_Obama", "an input string of the format 'start,end'")
	flag.Parse()

	reader := *pages
	s := strings.Split(reader, ",")
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://en.wikipedia.org/" + "wiki/" + s[0])
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
