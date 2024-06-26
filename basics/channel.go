package main

import "fmt"

func channel() {
	n := 5
	c := make(chan int)

	// Goroutine to send numbers to the channel
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
