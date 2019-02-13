package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/moficodes/gophercises/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url of the site")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(response.Body)

	for _, l := range links {
		fmt.Println(l)
	}
}

/*

	1. Get the webpage
	2. parse all the links from the page
	3. build proper urls with the links
	4. filter links with diff domain
	5. find all the pages with BFS search
	6. Print out XML

*/
