package main

import "fmt"

func main() {
	// slice は length: 長さ, capacity: 容量 を持っている。
	// 長さはそれに含まれる要素の数
	// 容量はスライスの最初の要素から数えて、「元となる配列の」要素数
	// 十分な容量を持って提供されているスライスを再スライスすることによって、スライスの長さを伸ばすことができる

	// len 6 cap 6
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// len 0 cap 6
	s = s[:0]
	printSlice(s)

	// len 4 cap 6
	s = s[:4]
	printSlice(s)

	// drop first 2 values
	// len 2 cap 4
	s = s[2:]
	printSlice(s)

	// panic: runtime error: slice bounds out of range [:100] with capacity 4
	// exit status 2
	// s = s[:100]
	// printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
