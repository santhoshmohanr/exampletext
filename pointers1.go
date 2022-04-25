package main

import "fmt"

func main() {
	name := "hello "
	b := &name
	fmt.Println(b)

	Add("world ", "go ", "world ", b)
	fmt.Println(name)
}
func Add(name1 string, name2 string, name3 string, b *string) {
	name := *b
	*b = name + name1 + name2 + name3

}
