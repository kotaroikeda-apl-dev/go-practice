package main

import "fmt"

// Greeter は挨拶を返す型を表すインターフェース
type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

// 値レシーバーなので Person 値でも満たす
func (p Person) Greet() string {
	return "Hello, " + p.Name
}

type Robot struct {
	ID int
}

// ポインタレシーバーなので *Robot だけが満たす
func (r *Robot) Greet() string {
	return fmt.Sprintf("Beep-%03d", r.ID)
}

func printGreeting(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	var g Greeter

	g = Person{Name: "Taro"} // implements を書かなくても満たす
	printGreeting(g)

	g = &Robot{ID: 42} // ポインタレシーバーなので & が必要
	printGreeting(g)

	// Robot 値は Greeter を満たさないためコンパイルエラーになる例
	// g = Robot{ID: 42}
}
