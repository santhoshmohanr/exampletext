package main

import "fmt"

func main() {
	a := []int{23, 45, 67, 89, 54, 34, 24, 68}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}

		}
	}
	fmt.Println(a)
}
