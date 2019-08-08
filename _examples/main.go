package main

import (
	"fmt"

	joke "github.com/rbo13/dad-joke"
)

func main() {

	res := joke.GetJSON()
	// res := joke.GetHTML()

	fmt.Print(res)
}
