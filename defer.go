package main

import "fmt"

func main() {
	a := []int{345, 245, 2577, 8678, 2356, 4566}

	largest(a)

}
func largest(a []int) {
	defer finished()
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	fmt.Println("lardest number", a, "is", max)
}
func finished() {
	fmt.Println("finished defer")
}
