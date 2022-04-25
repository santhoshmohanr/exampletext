package main

import "fmt"

func main() {
	a := [...]int{10, 20, 30, 40}

	b := a[:]
	Test(b)
	fmt.Println(b)
}

func Test(b []int) {
	b[0] = 57
	//c := b[:]

	//c = append(c, 80)
	//return c
	//	fmt.Println(c)
	//b = append(b, 57)
}
