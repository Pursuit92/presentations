package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	line, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println("You said:", line)
}
