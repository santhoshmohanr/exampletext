package main

import "fmt"

type name struct {
	firstname string
	lastname  string
}

func main() {
	name1 := name{
		firstname: "jeshna",
		lastname:  "mohan",
	}
	name2 := name{
		firstname: "devi",
		lastname:  "mohan",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")

	} else {
		fmt.Println("name1 and name2 are not equal")
	}

}
