package main

import "fmt"

// Messenger はメッセージ文字列を返す型のインターフェース
type Messenger interface {
	Message() string
}

type Greeting struct {
	To string
}

// 値レシーバー: Greeting 値でも満たす
func (g Greeting) Message() string {
	return "Hello, " + g.To
}

type Alert struct {
	Level string
	Code  int
}

// ポインタレシーバー: *Alert だけが満たす。nil も扱う。
func (a *Alert) Message() string {
	if a == nil {
		return "(nil alert)"
	}
	return fmt.Sprintf("[%s] code=%d", a.Level, a.Code)
}

func dump(label string, m Messenger) {
	if m == nil {
		fmt.Printf("%s: <nil>\n", label)
		return
	}
	fmt.Printf("%s: type=%T value=%v message=%q\n", label, m, m, m.Message())
}

func main() {
	var m Messenger // ゼロ値は <nil>（型も値も空）
	dump("zero", m)

	m = Greeting{To: "Gopher"} // 値レシーバーなのでそのまま代入
	dump("greeting", m)

	var nilAlert *Alert // 動的型は *Alert、値は nil（非 nil インターフェース）
	m = nilAlert
	dump("nil alert value", m)

	m = &Alert{Level: "warn", Code: 503} // 動的型と値を差し替え
	dump("alert", m)
}
