package main

import "fmt"

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		oldFirst := first
		first = second
		second = oldFirst + second
		return oldFirst
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
