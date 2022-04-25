package main

import "fmt"

func main() {
	fmt.Println("main function")
	chn := make(chan string)
	go hello(chn)
	fmt.Println(<-chn)
	fmt.Println("last function")

}
func hello(chn chan string) {
	fmt.Println("hello world")
	chn <- "golang"
}
