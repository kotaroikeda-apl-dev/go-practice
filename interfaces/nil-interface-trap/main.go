package main

import (
	"fmt"
)

func main() {
	fmt.Println("nil interface の正体デモ（error が nil なのに nil じゃない理由）")
	fmt.Println()

	// 例1: var err error = nil → (nil, nil)
	var err1 error = nil
	showInterfaceInternals("例1: var err error = nil", err1)

	// 例2: var e *MyError = nil; var err error = e → (*MyError, nil)
	var e *MyError = nil
	var err2 error = e
	showInterfaceInternals("例2: var e *MyError=nil; var err error = e", err2)

	fmt.Printf("err2 != nil ? %v\n", err2 != nil)
	fmt.Println()

	// 典型例: return nil *MyError を error として返してしまう
	errBad := runBad()
	showInterfaceInternals("地雷: runBad() error", errBad)
	fmt.Printf("runBad: err == nil ? %v\n", errBad == nil)
	fmt.Println("runBad: fmt.Println(err) は Error() を呼ぶ（nil レシーバでも呼べる）:")
	fmt.Println(errBad)
	fmt.Println()

	// 正しい: return nil
	errGood := runGood()
	showInterfaceInternals("正: runGood() error", errGood)
	fmt.Printf("runGood: err == nil ? %v\n", errGood == nil)
	fmt.Println()

	// 安全なパターン
	errSafe1 := runSafe(false)
	errSafe2 := runSafe(true)
	showInterfaceInternals("安全: runSafe(false)", errSafe1)
	showInterfaceInternals("安全: runSafe(true)", errSafe2)
}
