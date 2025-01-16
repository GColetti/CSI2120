package main

import "fmt"

func main() {

	x := []int{3, 1, 4, 1, 5, 9, 2, 6}

	var y [8]int

	// Add channel
	ch := make(chan int, len(x))

	// parallel loop
	for i, v := range x {

		go func(i int, v int) {
			calcul(v, ch)

		}(i, v) // Call to goroutine
	}

	// Add synchronization, receive from channel
	for i := range x {
		y[i] = <-ch
	}

	// Print result and close channel
	fmt.Println(y)
	close(ch)

}

// You can add a channel to the list of parameters
func calcul(v int, ch chan int) {
	ch <- 2*v*v*v + v*v

}
