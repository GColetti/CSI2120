package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func numbers(sz int) (res chan float64) {
	res = make(chan float64)
	go func() {
		defer close(res)
		num := 0.0
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		for i := 0; i < sz; i++ {
			num += math.Sqrt(math.Abs(rand.Float64()))
		}
		num /= float64(sz)
		res <- num
		return
	}()
	return
}

func main() {
	var nGo int
	rand.Seed(42)
	fmt.Print("Number of Go routines: ")
	fmt.Scanf("%d \n", &nGo)
	res := make([]chan float64, nGo)

	var wg sync.WaitGroup
	wg.Add(nGo)
	for i := 0; i < nGo; i++ {
		res[i] = numbers(1000)
		// Debug print statements
		// fmt.Println(<-res[i])
		wg.Done()

	}
	wg.Wait()

}

// 69
// 14
// 88
