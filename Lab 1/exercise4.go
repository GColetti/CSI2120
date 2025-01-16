package main

import "fmt"

// Define course structure
type Course struct {
	numStudents int
	professor   string
	average     float64
}

// Exercise 4
func main() {
	// Create a dynamic map m
	m := make(map[string]Course)

	// Add the courses CSI2120 and CSI2110 to the map
	m["CSI2120"] = Course{150, "Lang", 78.2}
	m["CSI2110"] = Course{175, "Moura", 82.0}

	for k, v := range m {
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.numStudents)
		fmt.Printf("Professor: %s\n", v.professor)
		fmt.Printf("Average: %f\n\n", v.average)
	}
}
