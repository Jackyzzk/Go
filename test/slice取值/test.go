package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 6, 5}
	fmt.Println(test[len(test) - 1])
	test1 := []int{}
	fmt.Println(test1, len(test1) == 0, test1 == nil)
	var test2 []int
	fmt.Println(test2, len(test2) == 0, test2 == nil)
}
