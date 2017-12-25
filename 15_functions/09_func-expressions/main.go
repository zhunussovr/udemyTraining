package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hi there")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
