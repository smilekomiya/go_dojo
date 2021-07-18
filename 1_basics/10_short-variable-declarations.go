package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// 関数の外ではこのような暗黙的な宣言はできない（var を使う）
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
