package main

import "fmt"

func main() {
	score := map[string]int{
		"team A": 20,
		"team B": 30,
	}
	fmt.Println(score)
	//b := &score
	//fmt.Println(b)
	add("teamc", 40, score)
	fmt.Println(score)
}
func add(key string, value int, b map[string]int) {
	//mapValue := *b
	b[key] = value
	//*b = mapValue
}

//	b := &a
//	Add("hello", "world", b)
//	fmt.Println(a)
//}
//
//func Add(name1 string, name2 string, b *map[string]int) {
//	a := *b
//	*b = a + name1 + name2
//}
