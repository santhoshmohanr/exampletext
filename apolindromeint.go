package main

import "fmt"

func main() {
	var num, tempnum, remainder int
	var reverse int = 0
	fmt.Println("enter the polindrome num:")
	fmt.Scan(&num)
	//num = 121
	tempnum = num

	for {
		//121
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10
		if num == 0 {
			break
		}
	}
	if tempnum == reverse {
		fmt.Printf("%d is a polindrome", tempnum)
	} else {
		fmt.Printf("%d is not a polindrome", tempnum)
	}

}
