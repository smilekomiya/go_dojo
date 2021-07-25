package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	(*p).Y = 2 // ポインタのフィールドのアクセス
	p.X = 1e9  // 上記は面倒な表記なのでこのように表記可能
	fmt.Println(v)
}
