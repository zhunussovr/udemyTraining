package main

import "fmt"

func main() {

	var s, b, i float64

	fmt.Println("In this program you should enter one small and one big number.")
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&s)
	fmt.Print("Please enter a big number: ")
	fmt.Scan(&b)

	i = b / s

	fmt.Println("The result is: ", i)

}
