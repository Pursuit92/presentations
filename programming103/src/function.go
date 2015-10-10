package main

import "fmt"

func SayHello(to string) string {
	return fmt.Sprintf("Hello, %s!", to)
}

func main() {
	fmt.Println(SayHello("dg"))
}
