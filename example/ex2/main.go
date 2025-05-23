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
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
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
