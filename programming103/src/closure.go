package main

import "fmt"

func main() {
	// main OMIT
	// first OMIT
	var closure func(string) = func(to string) {
		fmt.Printf("Hello, %s!\n", to)
	}
	// end first OMIT

	world := "World"

	another := func() {
		closure(world)
	}

	another()
	// end main OMIT
}
