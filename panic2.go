package main

import "fmt"

func main() {
	firstname := "jeshna"
	lastname := ""
	fullname(firstname, lastname)
	defer fmt.Println("defered call in main")
	fmt.Println("normally returned in main ")
}
func fullname(firstname string, lastname string) {
	defer recovername()
	if firstname == "" {
		panic("runtime error:firstname cannot be found")

	}
	if lastname == "" {
		panic("runtime error:lastname cannot be found")
	}
	fmt.Printf("%s %s\n", firstname, lastname)
	fmt.Println("normally retruned in fullname")
}
func recovername() {
	if r := recover(); r != nil {
		fmt.Println("recovered name", r)
	}
}
