package main

import "fmt"

// Exercise 2
func main() {
	lineWidth := 5
	s := "x"
	lineSymb := s
	formatStr := fmt.Sprintf("%%%ds\n", lineWidth)

	for ; len(lineSymb) < lineWidth; lineSymb += s {
		fmt.Printf(formatStr, lineSymb)
	}
	for ; len(lineSymb) > 0; lineSymb = lineSymb[:len(lineSymb)-1] {
		fmt.Printf(formatStr, lineSymb)
	}

}
