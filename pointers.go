package main

import "fmt"

func main() {
	fmt.Println("=========")
	totalPrice := 0

	b := &totalPrice
	fmt.Println(b)

	CalculateBill(10, 30, b)
	fmt.Println("total price", totalPrice)
	//CalculateBill(30, 30, b)
	//fmt.Println("total price", &totalPrice)

}

func CalculateBill(price1 int, price2 int, totalPrice *int) {
	totalAmount := *totalPrice
	*totalPrice = totalAmount + price1 + price2
}
