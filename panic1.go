package main

import "fmt"

func main() {
	slicepanic()
	fmt.Println("normal returned ")

}
func slicepanic() {
	a := []int{34, 54, 65}
	fmt.Println(a[5])
	fmt.Println("last call can not be work", a)

}
