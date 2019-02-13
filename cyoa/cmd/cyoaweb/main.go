package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/moficodes/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3030, "the port to start the server on")
	filename := flag.String("file", "gopher.json", "the JSON file with the Story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Starting server on port : %d", *port)

	handler := cyoa.NewHandler(story, cyoa.WithPathFn(custom))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
