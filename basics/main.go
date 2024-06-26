package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Enter a positive value for n: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if n <= 0 || math.IsNaN(float64(n)) {
		fmt.Println("Invalid input: n must be a positive integer value.")
		return
	}
	fmt.Println("Executing loop():")
	loop(n)

	fmt.Println("\nExecuting channel():")
	channel()
}
