package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/moficodes/gophercises/urlshortener/urlshort"
)

var fileName string

func init() {
	const (
		fileValue = "urls.yaml"
		fileUsage = "File to get the questions from"
	)
	flag.StringVar(&fileName, "file", fileValue, fileUsage)
	flag.StringVar(&fileName, "f", fileValue, fileUsage+" (shorthand)")
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	log.Fatalln(http.ListenAndServe(":8080", yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
