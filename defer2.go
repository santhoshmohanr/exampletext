package main

import (
	"errors"
	"fmt"
)

func main() {
	flag, err := errfunc(7)
	fmt.Println(flag)
	fmt.Println(err)

	flag, err = errfunc(12)
	fmt.Println(flag)
	fmt.Println(err)
}
func errfunc(a int) (bool, error) {
	if a%7 == 0 {
		return true, nil
	} else {
		return false, errors.New("number not divide by 2")
	}
}
