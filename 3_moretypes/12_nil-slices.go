package main

import "fmt"

func main() {
	var s []int

	// slice のゼロ値はnil
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("s is nil!")
	}

	// array
	// a := [0]int{}
	// fmt.Println(a, len(a), cap(a))
	// if a == nil {
	// 	// len 0, cap 0は nil
	// 	fmt.Println("a is nil!")
	// }

	t := []int{1}
	u := t[0:0]
	fmt.Println(u, len(u), cap(u))
	if u == nil {
		// len 0, cap 1は nil ではない
		fmt.Println("u is nil!")
	}
}
