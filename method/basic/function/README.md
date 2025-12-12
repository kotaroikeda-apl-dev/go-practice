# ポインタと関数

このディレクトリでは、メソッドではなく通常の関数として、ポインタと値の違いを学びます。

## ポインタと関数とは

メソッドではなく通常の関数として定義する場合でも、ポインタと値の違いは同じです。

## 値で受け取る関数とポインタで受け取る関数の違い

1. **値で受け取る関数**: `func SetValue(c Counter, value int)`

   - 引数として構造体のコピーが渡される
   - 関数内で引数を変更しても、元の値は変更されない
   - `SetValue(counter, 100)`を呼び出しても、元の`Counter`の値は変更されない

2. **ポインタで受け取る関数**: `func Scale(c *Counter, factor int)`

   - 引数としてポインタが渡される
   - 関数内でポインタが指す値を変更すると、元の値が変更される
   - `Scale(&counter, 3)`を呼び出すと、元の`Counter`の値が変更される

## この例で学ぶこと

1. **値で受け取る関数**: `func SetValue(c Counter, value int)`

   - 構造体を値で受け取る
   - 関数内で変更しても元の値は影響を受けない
   - `SetValue(counter, 100)`を呼び出しても、元の`Counter`の値は変更されない

2. **ポインタで受け取る関数**: `func Scale(c *Counter, factor int)`

   - 構造体のポインタを受け取る
   - 関数内で`c.Value`を変更すると、元の構造体の値が変更される

3. **関数の呼び出し**: `Scale(&counter, 3)`

   - ポインタで受け取る関数を呼び出す際は、`&counter`のようにアドレスを渡す必要がある
   - 値で受け取る関数は`counter`をそのまま渡す

4. **メソッドとの違い**
   - メソッド: `counter.Multiply(3)`のように、レシーバーに対して呼び出す
   - 関数: `Scale(&counter, 3)`のように、関数名と引数として渡す

## 実行方法

```bash
go run main.go
```

## 出力例

```
=== Pointers and Functions ===
初期状態: Counter{Value: 5}
SetValue(counter, 100)呼び出し後（変更なし）: Counter{Value: 5}
Scale(&counter, 3)呼び出し後（変更あり）: Counter{Value: 15}
Increment(&counter)呼び出し後（変更あり）: Counter{Value: 16}
```
