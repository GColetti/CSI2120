package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// max value for prime
const maxi int64 = 50000000

// PRODUCER
// a producer of prime numbers
func producer(ch chan int64, quit chan bool, wg sync.WaitGroup) {
	var number int64
	defer wg.Done()
	for {
		number = getPrime(maxi)

		select {

		case ch <- number: // send number to ch
		case <-quit: // to stop number generation
			close(ch)
			close(quit)
			return
		}
	}
	wg.Wait()
}

// CONSUMER
// a consumer of prime numbers
// gets two primes a and b
// if the product a*b ends with suffix then stop
// here suffix is always 2 digits
func consumer(suffix int64, ch chan int64, done chan bool, quit chan bool) {
	for {
		a := <-ch
		b := <-ch

		if a*b%100 == suffix { // checks if last 2 digits
			fmt.Println(a * b)
			quit <- true // are equal to suffix
			break
		}
	}
	done <- true
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	ch := make(chan int64)  // to send/receive prime numbers
	done := make(chan bool) // when a secret key is found
	quit := make(chan bool) // used to close producer channel

	// produces primes until true
	// is sent to channel quit

	// Sync
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go producer(ch, quit, wg)
	}

	// consumes prime numbers received
	// from channel ch
	go consumer(11, ch, done, quit)
	<-done

}

// DO NOT modify the code below
// Prime number generator functions
// returns true if the number is prime
func isPrime(v int64) bool {

	sq := int64(math.Sqrt(float64(v))) + 1
	var i int64

	for i = 2; i < sq; i++ {

		if v%i == 0 {
			return false
		}
	}
	return true
}

// generate a prime number
func getPrime(maxValue int64) int64 {
	for {
		n := rand.Int63n(maxValue)
		if isPrime(n) {
			return n
		}
	}
}
