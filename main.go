package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	pages := flag.String("pages", "Lion,Barack_Obama", "an input string of the format 'start,end'")
	flag.Parse()

	reader := *pages
	fmt.Println(reader)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
