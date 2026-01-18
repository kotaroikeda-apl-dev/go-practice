# DefaultServeMux

`DefaultServeMux` は、Go の `net/http` パッケージであらかじめ定義されているデフォルトの `ServeMux` インスタンスです。

## 概要

`net/http` パッケージ内では以下のように定義されています。

```go
var DefaultServeMux = &defaultServeMux
```

> DefaultServeMux is the default ServeMux used by Serve.

## 特徴

1.  **グローバルな共有**: アプリケーション全体で共有される単一のマルチプレクサです。
2.  **暗黙的な使用**: `http.Handle` や `http.HandleFunc` といったパッケージレベルの関数を呼び出すと、自動的に `DefaultServeMux` に登録されます。
3.  **nil ハンドラ**: `http.ListenAndServe(":8080", nil)` のようにハンドラに `nil` を渡すと、`DefaultServeMux` が使用されます。

## 注意点

大規模なアプリケーションやテストの容易性を重視する場合、または複数の異なるサーバー設定が必要な場合は、`http.NewServeMux()` を使用して独自の `ServeMux` を作成することが推奨されます。
`DefaultServeMux` はグローバル変数であるため、サードパーティのパッケージが意図せずルートを登録してしまうなどの副作用が生じる可能性があるためです。

