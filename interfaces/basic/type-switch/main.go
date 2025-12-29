package main

import "fmt"

// type switch を使って interface{} に入った複数の型をまとめて処理する例
func main() {
	values := []interface{}{
		nil,
		"gopher",
		21,
		3.14,
		[]byte{1, 2, 3},
	}

	for _, v := range values {
		describe(v)
	}
}

func describe(v interface{}) {
	switch vv := v.(type) {
	case nil:
		fmt.Println("nil value")
	case string:
		fmt.Printf("string len=%d val=%q\n", len(vv), vv)
	case int:
		fmt.Printf("int doubled=%d\n", vv*2)
	case float64:
		fmt.Printf("float ceil-ish=%d\n", int(vv)+1)
	case []byte:
		fmt.Printf("bytes len=%d\n", len(vv))
	default:
		fmt.Printf("unknown type %T\n", vv)
	}
}



