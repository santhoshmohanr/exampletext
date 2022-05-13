package main

import "fmt"

func main() {
	a := [...]int{23, 45, 67, 86, 32, 21, 60}
	b := []int{}
	fmt.Printf("oriinal array:%d\n", a)
	fmt.Printf("reversed array:")
	length := len(a)
	fmt.Println(length)

	for i := length - 1; i >= 0; i-- {
		//fmt.Printf(" %d ", a[i])
		b = append(b, a[i])

	}

	fmt.Println(b)
}
