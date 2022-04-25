package main

import "fmt"

func main() {
	firstname := ""
	lastname := ""
	fullname(firstname, lastname)

}
func fullname(firstname string, lastname string) {
	if firstname == "" {
		fmt.Println("runtime error:firstname cannot be found")

	}
	if lastname == "" {
		fmt.Println("runtime error:lastname cannot be found")
	}
	fmt.Printf("%s %s\n", firstname, lastname)
}
