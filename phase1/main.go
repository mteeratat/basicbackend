package main

import (
	"fmt"
	"phase1/factFibo"
)

func main() {
	fmt.Println("Hello, World!")
	num := 7
	fmt.Printf("Factorial of %d = %d", num, factFibo.CalFactorials(num))
	fmt.Printf("\nFibonacci %d = %d", num, factFibo.CalFibonacci(num))
}
