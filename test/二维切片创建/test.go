package main

import "fmt"

func main() {
	n := 6
	p := make([][]int, 0, n)
	for i:= 0; i < n; i++ {
		t := make([]int, n)  // make 后面只有一个参数的时候，len = cap = n
		p = append(p, t)
	}
	fmt.Println(p)
	p[2][3] = 7
	fmt.Println(p)
}
