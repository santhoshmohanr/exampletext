package main

import "fmt"

func main() {
	a := []int{2, 7, 9, 20, 59, 60, 88, 100}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			num := a[i] + a[j]
			if num%2 == 0 {
				fmt.Printf("(%d,%d)", a[i], a[j])
			}
		}
	}

}
