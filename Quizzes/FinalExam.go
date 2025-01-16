package main

import (
	"fmt"
	"sync"
)

func main() {
	// array to be processed
	tab := [4][8]int{
		{0, 1, 5, 6, 2, 3, 2, 3},
		{4, 5, 6, 5, 6, 5, 6, 7},
		{4, 5, 2, 3, 6, 7, 5, 2},
		{8, 9, 5, 2, 5, 2, 1, 1}}

	var sum [4]int
	// Sequential processing of the array
	// Modify the code below
	// vvvvvvvvvvvvvvvvvvvvvvvvvvv

	var wg sync.WaitGroup

	// computes the sum of each row
	for i := 0; i < 4; i++ {
		sum[i] = 0
		for _, v := range tab[i] {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				sum[i] += v
			}(i)
			wg.Wait()

		}

	}
	fmt.Println("Expected: [22 44 34 33]")
	fmt.Println("Actual:  ", sum)
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	// End of sequential processing
	// Modify the code above

	// gets the max of row sums
	summax := 0
	for i := 0; i < 4; i++ {
		if sum[i] > summax {
			summax = sum[i]
		}
	}
	fmt.Print(summax)
}
