package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/moficodes/gophercises/link"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlSet struct {
	Urls []loc `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url of the site")
	maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	var toXML = urlSet{
		Urls: make([]loc, len(pages)),
	}
	for i, page := range pages {
		toXML.Urls[i] = loc{page}
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}

	fmt.Println()
}

type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]empty)
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}
			for _, link := range get(url) {
				nq[link] = empty{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	var ret []string
	links, _ := link.Parse(r)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

/*

	1. Get the webpage
	2. parse all the links from the page
	3. build proper urls with the links
	4. filter links with diff domain
	5. find all the pages with BFS search
	6. Print out XML

*/
