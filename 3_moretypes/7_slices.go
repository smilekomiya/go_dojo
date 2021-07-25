package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 最初の要素は含み最後の要素を含まない半開区間
	var s []int = primes[1:4]
	fmt.Println(s)
}
