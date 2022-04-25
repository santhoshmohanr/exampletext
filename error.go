package main

import (
	"errors"
	"fmt"
)

func main() {
	a := []int{43, 56, 76, 53, 59}
	fmt.Println(a)
	b := 57
	check, err := check(a, b)
	fmt.Println("check", check)
	fmt.Println("err", err)
}

func check(a []int, b int) (bool, error) {
	found := false
	for i := 0; i < len(a); i++ {
		if b == a[i] {
			found = true
		}
	}
	if found == true {
		fmt.Println("found")
		return true, nil
	} else {
		//fmt.Println(errors.New("notfound"))
		return false, errors.New("not found")
	}
}
