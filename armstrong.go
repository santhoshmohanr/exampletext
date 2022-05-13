package main

import "fmt"

func main() {

	var num, tempnum, remainder int
	var result int = 0
	fmt.Print("enter the number:")
	fmt.Scan(&num)
	tempnum = num

	for {
		remainder = tempnum % 10
		result += remainder * remainder * remainder
		tempnum /= 10
		if tempnum == 0 {
			break
		}
	}
	if result == num {
		fmt.Printf("%d is amstrong num ", num)
	} else {
		fmt.Printf("%d is not amstrong num", num)
	}

}
