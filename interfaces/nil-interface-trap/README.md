# nil-interface-trap

`error` が **nil のはずなのに `err != nil` になる**（typed nil / nil interface）の典型パターンを、実行して観察するためのサンプルです。

## 使い方

```bash
cd interfaces/nil-interface-trap
go run ./...
```

※ `go run .` でも同じデモが動きます。

## 何が見える？

- `var err error = nil` は **(type=nil, value=nil)** なので `err == nil`
- `var e *MyError = nil; var err error = e` は **(type=\*MyError, value=nil)** なので `err != nil`
- 地雷 `runBad()` が「なぜ `err != nil` になるか」を出力で確認できます
