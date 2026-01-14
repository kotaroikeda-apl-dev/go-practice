# http.FileServer

`http.FileServer` は、ローカルディスクなどのファイルシステムの内容を HTTP リクエストに応じて配信するためのハンドラを提供します。

## 特徴

- 静的ファイル（HTML, CSS, JavaScript, 画像など）の配信に最適です。
- リクエストされたパスに対応するファイルを探し、存在すればその内容を返します。
- ディレクトリがリクエストされた場合、その中に `index.html` があればそれを返し、なければファイル一覧を表示します（設定により変更可能）。
- `index.html` への直接リクエスト（例: `/index.html`）は、自動的にディレクトリパス（例: `/`）へリダイレクトされます。

## サンプルコードの解説

```go
// 静的ファイルが置かれているディレクトリを指定します
fs := http.Dir("./static")

// FileServer ハンドラを作成します
fileServer := http.FileServer(fs)

// ハンドラを登録します
http.Handle("/", fileServer)
```

## 実行方法

1. このディレクトリに移動します。
2. サーバーを起動します。

```bash
go run main.go
```

3. ブラウザで `http://localhost:8080` にアクセスします。
   `./static/index.html` の内容が表示されます。
