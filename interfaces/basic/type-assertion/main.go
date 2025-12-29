package main

import (
	"fmt"
)

// sample values to demonstrate type assertions
func main() {
	var any interface{} = "hello assertion"

	// 単一戻り値: 型が合わないと panic するので注意
	s := any.(string)
	fmt.Println("as string:", s)

	// 2値戻り値: 失敗しても panic しない
	i, ok := any.(int)
	fmt.Printf("as int: %v (ok=%v)\n", i, ok)

	// interface に入っている値を入れ替えて再利用
	any = 3.14
	f, ok := any.(float64)
	fmt.Printf("as float64: %.2f (ok=%v)\n", f, ok)
}


