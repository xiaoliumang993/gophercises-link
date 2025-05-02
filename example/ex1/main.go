package main

import (
	"fmt"
	"github.com/xiaoliumang993/gophercises-link"
	"strings"
)

//var exampleHtml = `
//<html>
//<body>
//  <h1>Hello!</h1>
//  <a href="/other-page">A link to another page</a>
//</body>
//</html>
//`

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.ParseHTML(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------")
	fmt.Printf("%+v\n", links)
}
