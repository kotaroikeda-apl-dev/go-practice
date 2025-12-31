package main

// runBad は典型的な地雷。
// 「*MyError は nil だけど、error に入った瞬間に (type=*MyError, value=nil) になってしまう」例。
func runBad() error {
	var e *MyError = nil
	return e // ★ここで error に暗黙変換される
}

// runGood は正しい書き方。
// エラーがないなら interface 全体を nil にする（= (type=nil, value=nil)）。
func runGood() error {
	return nil
}

// runSafe は安全なパターン（分岐で nil か実体を返す）。
func runSafe(fail bool) error {
	if !fail {
		return nil
	}
	return &MyError{Msg: "something went wrong"}
}
