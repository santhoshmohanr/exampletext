package main

import "fmt"

func main() {
	a := []int{6, 3, 45, 15, 36, 98, 12, 10}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)

}
