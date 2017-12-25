/*
Problem 5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?


*/

package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	for i := 1; i <= 20; i++ {
		if x%i == 0 {
			fmt.Println(x)
		}
	}
	return x * factorial(x-1)

}

func main() {
	a := 20
	z := factorial(a)
	fmt.Println("The factorial of 20 is ", z)
}
