package main

import (
	"fmt"

	// "github.com/swinton/golang-playground/life"
	"github.com/swinton/golang-playground/octocat"
	// "rsc.io/quote"
)

func main() {
	// fmt.Println(quote.Opt())
	// fmt.Println("The meaning of life:", life.Meaning())

	octocat, err := octocat.Get()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, based in %s, has %d followers. For more info, see: %s.\n", octocat.Name, octocat.Location, octocat.Followers, octocat.URL)
}
