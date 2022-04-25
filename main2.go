package main

import "fmt"

func main() {
	var a, b = 10, 20
	a, b = b, a
	fmt.Println("a =", a, "\nb =", b)

}
