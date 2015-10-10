package main

import "fmt"

func main() {
	// chan OMIT
	c := make(chan string)

	go func() {
		for _, v := range []string{"is", "this", "the", "real", "life?"} {
			c <- v
		}
		close(c)
	}()

	for v := range c {
		fmt.Print(v, " ")
	}
	// end chan OMIT

	fmt.Println()

}
