package main

type dog struct {
	name   string
	race   string
	female bool
}

func (d *dog) rename(newName string) {
	d.name = newName
	return
}

func main() {
	fido := dog{"Fido", "Poodle", false}
	// println(fido.name)
	fido.rename("Cocotte")
	// println(fido.name
}
