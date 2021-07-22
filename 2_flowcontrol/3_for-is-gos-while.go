package main

import "fmt"

func main() {
	sum := 1
	// for ; sum < 1000 ; {
	// ; を省略して while 的な使い方も可能。Go には while がないよ。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
