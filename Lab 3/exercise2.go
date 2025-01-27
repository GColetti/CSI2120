package main

import (
	"fmt"
	"math"
)

type NegError struct {
	num float64 // negative number
}

func (e *NegError) Error() string {
	return "Main NegError: "
}

type Div0Error struct {
}

func (e *Div0Error) Error() string {
	return "Main Div0Error: "
}

func rootDivN(num float64, n int) (res float64, err error) {
	if num < 0.0 {
		return num, &NegError{num}
	}
	if n == 0 {
		return num, &Div0Error{}
	}
	res = math.Sqrt(num) / float64(n)
	return
}

func main() {
	divs := []int{2, 10, 3, 0}
	nums := []float64{511.8, 0.65, -3.0, 2.123}

	for i, num := range nums {
		fmt.Printf("%d) sqrt(%f)/%d = ", i, num, divs[i])
		res, err := rootDivN(num, divs[i])
		if err == nil {
			fmt.Printf("%f\n", res)
		} else {
			switch err := err.(type) {
			case *NegError:
				fmt.Printf("Main NegError: Negative Number %f\n", err.num)
			case *Div0Error:
				fmt.Println("Main Div0Error: Division by 0 ")

			}
		}
	}
}
