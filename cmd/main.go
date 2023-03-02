package main

import (
	"fmt"

	"github.com/sgtcodfish/gotest-one/cmd/app"
)

func main() {
	src := []string{
		"hello world!",
		"not i with dot",
		"'n ijsberg",
		"here comes O'Brian",
	}

	for _, s := range src {
		fmt.Println(app.Transformer(s))
	}
}
