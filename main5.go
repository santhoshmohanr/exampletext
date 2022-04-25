package main

import "fmt"

func main() {
	a := []int{34, 56, 76, 87, 89, 90, 99}
	count := 0

	for i := 0; i < len(a); i++ {
		count++

	}
	fmt.Print(count)
}
