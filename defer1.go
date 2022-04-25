package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	p := person{
		firstname: "jeshna",
		lastname:  "mohan",
	}
	defer p.fullname()
	fmt.Print("welcome ")
}
func (p person) fullname() {
	fmt.Printf("%s%s ", p.firstname, p.lastname)
}
