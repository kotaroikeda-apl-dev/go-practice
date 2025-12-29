package main

import (
	"fmt"
	"reflect"
)

// showInterfaceInternals は「interface は (type, value) のペア」を観察するための補助関数。
//
// - %T で動的型を見る
// - reflect で「中身が nil になり得る型か」を判定しつつ IsNil を安全に使う
// - panic しない実装にする（reflect.Value.IsNil は Kind によっては panic するため）
func showInterfaceInternals(label string, x any) {
	fmt.Printf("== %s ==\n", label)
	fmt.Printf("x == nil? %v\n", x == nil)
	fmt.Printf("%%T: %T\n", x)
	fmt.Printf("%%v: %v\n", x)

	if x == nil {
		fmt.Println("reflect: <nil interface> (type=nil, value=nil)")
		fmt.Println()
		return
	}

	t := reflect.TypeOf(x)
	fmt.Printf("reflect.TypeOf: %v\n", t)
	fmt.Printf("reflect.Kind:  %v\n", t.Kind())

	v := reflect.ValueOf(x)

	// IsNil が使える Kind だけに限定する（それ以外は呼ぶと panic）
	switch t.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Interface, reflect.Slice:
		fmt.Printf("reflect.Value.IsNil(): %v\n", v.IsNil())
	default:
		fmt.Println("reflect.Value.IsNil(): (not applicable)")
	}

	fmt.Println()
}


