package main

import "fmt"

func main() {
	// array OMIT
	a := [4]int{1, 2, 3, 4}
	// end array OMIT

	// iter OMIT
	for i := range a {
		fmt.Println(i)
	}
	// end iter OMIT

	// iter2 OMIT
	for i, v := range a {
		fmt.Printf("a[%d] = %d", i, v)
	}
	// end iter2 OMIT
}
