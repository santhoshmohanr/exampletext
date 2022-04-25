package main

import "fmt"

func main() {
	a := []int{34, 54, 65}
	b := []int{54, 76, 89}
	find(45, a...)
	find(54, b...)

}
func find(a int, b ...int) {
	fmt.Printf("type of num is %T\n", a)
	found := false
	for i, v := range b {
		if v == a {
			fmt.Println(a, "found of index", i, "in", b)
			found = true
		}
	}
	if !found {
		fmt.Println(a, "not found of index", b)
	}

}
