package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/moficodes/gophercises/link"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to <span>NOT TODAY</span> another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", links)
}
