package main

import "fmt"

/*  Parameters: float value n
*  Function: Finds the corresponding ceiling and floor values of n
*  Returns: integer values x, y
 */
func limits(n float64) (x, y int) {
	x = int(n)
	if x > 0 {
		y = x + 1
	} else {
		y = x
		x = y - 1
	}

	return x, y
}

// Exercise 1
func main() {
	x := 32.6
	var l, u = limits(x)
	fmt.Println("Limits for", x, "are:", l, u)
	y := -3.2
	var a, b = limits(y)
	fmt.Println("Limits for", y, "are:", a, b)
}
