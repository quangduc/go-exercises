package main

import (
	"fmt"
	"strings"

	link "github.com/quangduc/htmllinkparser"
)

var exampleHmtl = `
	<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
	</body>
	</html>
`

func main() {
	r := strings.NewReader(exampleHmtl)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", links)
}
