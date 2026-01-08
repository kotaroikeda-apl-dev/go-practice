# go-install-hello

`go install` の挙動を、手を動かしながら理解するための最小CLIサンプルです。

このサンプルで伝えたいことは次の3つです。

- `go install` はコマンドをビルドして、決まった場所に置く
- コマンドを叩いて実行できるかどうかは PATH で決まる
- `go install <module>@<version>` は、指定した版のコマンドを入れるために使う

## このリポジトリが役立つ人

- `go install` / `go get` / `@latest` / `PATH` 周りで混乱したことがある人
- 「なぜ `@version` が必要なのか」を、自分の言葉で説明できるようになりたい人

## 最短で試す

この3つを順に実行して、出力の違いを見てください。

```bash
cd goInstallHello
go run ./cmd/hello --version
mkdir -p ./bin
GOBIN=$(pwd)/bin go install ./cmd/hello
./bin/hello --version
```

もし `hello` を `hello --version` のようにコマンド名だけで実行したい場合は、インストール先が PATH に入っている必要があります。後述の「コマンドが使えないときの見方」を参照してください。

補足: 何度も試すなら、`GOBIN` をこのターミナルの間だけ固定しておくと楽です（新しいターミナルでは戻ります）。

```bash
cd goInstallHello
mkdir -p ./bin
export GOBIN="$(pwd)/bin"
go install ./cmd/hello
./bin/hello --version
```

## ディレクトリ構成

```
.
├── cmd/
│   └── hello/
│       └── main.go
└── go.mod
```

## 前提

- Go Modules を使います（このリポジトリは `go.mod` を持ちます）。
- `go install` の挙動は「`@version` を付けるかどうか」で大きく変わります。
  - このリポジトリをcloneして試すだけなら `go install ./cmd/hello` でOKです（ローカルのコードをビルドしてインストール）。
  - 「依存管理と切り離したツール導入」を試したい場合は、次の流れになります。
    - リポジトリを GitHub に push（例: `github.com/kotaroikeda-apl-dev/go-practice`）
    - タグを作る（例: `goInstallHello/v0.1.0`）
    - 任意の場所で `go install github.com/kotaroikeda-apl-dev/go-practice/goInstallHello/cmd/hello@v0.1.0` を実行する

## 実行

```bash
cd goInstallHello
go run ./cmd/hello
go run ./cmd/hello --name gopher
go run ./cmd/hello --version
```

`--version` は `runtime/debug.ReadBuildInfo()` を使って、ビルドされたバイナリに埋め込まれたモジュール情報を表示します。  
（`go install ...@vX.Y.Z` で入れた場合に、`version: vX.Y.Z` のように見えるのがポイントです。）

## ソースコードの説明

対象: `cmd/hello/main.go`

このコマンドは「どの `hello` を実行しているか」と「どのモジュール/バージョン由来でビルドされたか」を確認できるようにしています。

まず `--version` を付けると、あいさつではなくビルド情報を表示します。

- `hello --version` で表示するもの
  - Goのバージョン（例: `go: go1.24.4`）
  - mainパッケージの場所（例: `main package: github.com/.../cmd/hello`）
  - モジュールパスとバージョン（例: `module: ...`, `version: ...`）
  - 実行中バイナリの実体パス（例: `executable: /.../bin/hello`）

実装の要点は次の3つです。

- `flag` で `--version` を実装する（確認用オプション）
- `runtime/debug.ReadBuildInfo()` でビルド情報を読む（モジュール/バージョンの観察）
- `os.Executable()` で実行中バイナリの場所を出す（PATHトラブルの切り分け）

## ローカルインストール

この方法は、今いるモジュール（このリポジトリ）のコードをビルドしてインストールします。

```bash
cd goInstallHello
go install ./cmd/hello
```

インストール先は `GOBIN` → 未設定なら `GOPATH/bin` です。

```bash
go env GOBIN
go env GOPATH
```

補足: もし「書き込み権限がない」環境（企業端末・CI・サンドボックス等）なら、`GOBIN` を明示して回避できます。

```bash
cd goInstallHello
mkdir -p ./bin
GOBIN=$(pwd)/bin go install ./cmd/hello
./bin/hello --version
```

## 配布インストール

この方法は「依存管理と切り離したツール導入」です。

1) このリポジトリを GitHub に push（例: `github.com/kotaroikeda-apl-dev/go-practice`）  
2) タグを作る（例: `goInstallHello/v0.1.0`）  
3) 任意の場所で、次を実行:

```bash
go install github.com/kotaroikeda-apl-dev/go-practice/goInstallHello/cmd/hello@v0.1.0
hello --version
```

ポイント:

- `@v0.1.0` のようにバージョンを固定すると、誰がいつ実行しても同じソースに解決されやすくなります（再現性）。
- `@latest` は便利ですが、将来変わり得ます（「常に最新」を取りにいくため、再現性は落ちます）。

## コマンドが使えないときの見方

ここでは PATH 周りの切り分けをします。

よくある原因は、インストール先（`GOBIN` または `GOPATH/bin`）が `PATH` に入っていないことです。

まずは次の3つをそのまま実行して、状況を把握します:

```bash
go env GOBIN
go env GOPATH
which hello || true
```

次に `PATH` を確認します（zsh例）:

```bash
echo $PATH
```

読み方:

- `go env GOBIN` が空の場合
  - インストール先は通常 `$(go env GOPATH)/bin` です（多くの環境で `~/go/bin`）
  - そのディレクトリが `PATH` に無いと `hello: command not found` になります
- `go env GOBIN` に何かパスが出る場合
  - インストール先はそのディレクトリです
  - そのディレクトリが `PATH` に無いと `hello: command not found` になります

すぐ直す（このコマンドを実行したターミナルだけ）:

```bash
export PATH="$(go env GOBIN):$PATH"
hash -r
which hello
hello --version
```

補足: このREADMEの「最短で試す」では、プロジェクト配下の `./bin` を `GOBIN` に指定しています。
その場合は `hello` ではなく `./bin/hello` なら常に実行できます（PATHは不要です）。

## このサンプルから学べること

- `go install` は「ビルドしてどこかに置く」までで、実行できるかは PATH 次第だと切り分けできる
- `GOBIN` 未設定時は `GOPATH/bin` に入る、という既定動作を自分の環境で確認できる
- `go install ./...`（ローカル）と `go install <pkg>@<version>`（配布）の違いを、ビルド情報の表示で観察できる
- 再現性のために `@vX.Y.Z` が重要である理由（同じソースに解決させるため）を実験として説明できる


