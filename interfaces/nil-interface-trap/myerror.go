package main

// MyError は error を満たす独自エラー型。
//
// ポイント:
// - Error() は pointer receiver（*MyError）で実装する
// - nil レシーバでも安全に動くようにしておくと、デモが分かりやすい（実務でも役に立つことがある）
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	// nil ポインタのメソッド呼び出し自体は可能。
	// ただし中でフィールド参照すると panic になるので、ここでガードする。
	if e == nil {
		return "MyError(<nil>)"
	}
	return "MyError(" + e.Msg + ")"
}
