package main

import (
	"fmt"
)

func main() {
	var originalString string = "2020"
	var reverseString string = ""
	var length = len(originalString)
	for i := length - 1; i >= 0; i-- {
		reverseString = fmt.Sprintf("%v%v", reverseString, string(originalString[i]))
	}
	fmt.Println(reverseString)

	if originalString == reverseString {
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}
