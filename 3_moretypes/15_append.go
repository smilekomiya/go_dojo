package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append する際、元の配列が変数群を追加する際に容量が小さい場合は、より大きいサイズの配列を割り当てなおし、戻り値となるスライスは新しい割当先を示すようになる
	// cap は要素数に応じて 0, 1, 2, ..., 2^n ごとに増える。（0, 1, 2, 4, ...）

	// append works on nil slice
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3)
	printSlice(s)

	s = append(s, 4)
	printSlice(s)

	s = append(s, 5)
	printSlice(s)

	s = append(s, 6)
	printSlice(s)

	s = append(s, 7)
	printSlice(s)

	s = append(s, 8)
	printSlice(s)

	s = append(s, 9)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
