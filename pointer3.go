package main

import "fmt"

func main() {
	a := [...]int{23, 54, 67, 87}
	b := &a
	fmt.Println(b)
	modifie(&a)
	fmt.Println(a)
}
func modifie(b *[4]int) {
	(*b)[3] = 29

}
