package main

import "fmt"

func loop(n int) {

	for i := 0; i < n; i++ {
		//loop for spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print("  ")
		}
		//loop for actual stars
		for j := 0; j <= i; j++ {
			fmt.Print("*   ")
		}
		fmt.Println()
	}
}
