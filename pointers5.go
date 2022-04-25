package main

import "fmt"

type family struct {
	name string
	age  int
	city string
}

func main() {
	//jeshna := family{"jeshna", 1, "trichy"}
	//swathi := family{"swathi", 24, "trichy"}

	families := []family{family{"jeshna", 1, "trichy"}, family{"swathi", 24, "trichy"}}
	for i := 0; i < len(families); i++ {
		fmt.Printf("Name: %v\n", families[i].name)
		fmt.Printf("age:%d\n", families[i].age)
		fmt.Printf("city:%s\n", families[i].city)
	}
}
