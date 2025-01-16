package main

import "fmt"

// Box
type Box struct {
	weight float64
	color  string
}

// methods
func (p *Box) SetColor(newColor string) {
	p.color = newColor
}

func (p Box) GetColor() string {
	return p.color
}

// interface
type color interface {
	SetColor(string)
	GetColor() string
}

func main() {

	var b = Box{32.4, "yellow"}

	var iii color

	// Modify this portion of code*********
	// Reference to the box struct
	iii = &b
	// End of the portion of code to modify ***

	// we want to print the color of the instance to which iii refers
	fmt.Printf("The color is: %s\n", iii.GetColor())
}

// The output that should be produced is:
// The color is: yellow
