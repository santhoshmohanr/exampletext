package main

import "fmt"

func main() {
	primenum(172)

}
func primenum(num int) {
	found := true
	for i := 2; i <= num/2; i++ {
		if num%2 == 0 {
			fmt.Println("This is not a prime num", num)
			found = false
			break
		}
	}
	if found == true {
		fmt.Println("this is prime num", num)
	}

}
