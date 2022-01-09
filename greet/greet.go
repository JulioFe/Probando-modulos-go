package greet

import "fmt"

var emoji = ";)"

func English() {
	fmt.Printf("Hello! %v\n", emoji)
}

func Italian() {
	fmt.Printf("Ciao! %v\n", emoji)
}
