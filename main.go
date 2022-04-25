package main

import "fmt"

func main() {
	Fibonacci(5)
}
func Fibonacci(n int) {
	var n1, n2, n3 = 0, 1, 0
	for i := 1; i <= n; i++ {
		fmt.Println(n1)
		n3 = n1 + n2
		n1 = n2
		n2 = n3
	}
}
