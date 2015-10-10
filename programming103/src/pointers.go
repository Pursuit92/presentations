package main

import "fmt"

// do OMIT
func doStuff() *int {
	n := 5
	return &n
}

// end do OMIT

func main() {
	// with OMIT
	fmt.Println(*doStuff())
	// end with OMIT
}
