package main

import "fmt"

func main() {
	a := []int{23, 54, 67, 64, 45, 4, 36, 43}
	b := 77
	found := false
	for i := 0; i < len(a); i++ {
		if b == a[i] {
			found = true

		}
	}
	if found == true {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

}
