package main

import "fmt"

// 空インターフェースはメソッドゼロ。どんな型でも受け取れる。
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i) // ゼロ値は <nil>

	i = 99 // int
	describe(i)

	i = "gopher" // string
	describe(i)

	i = true // bool
	describe(i)
}
