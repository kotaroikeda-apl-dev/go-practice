# concurrency/basic

Go 言語の並行処理の基本（goroutine と channel）を学習するためのサンプルコード。

## 概要

goroutine の基本的な使い方と、channel を使用して goroutine 間でデータを送受信する基本的な例を示しています。

## 実行方法

### 個別に実行する場合

各ディレクトリで個別に実行できます：

```bash
# goroutineの例を実行
cd goroutine
go run main.go

# channelBasicの例を実行
cd channelBasic
go run main.go
```

各ディレクトリには個別の README.md があります。詳細は各ディレクトリの README.md を参照してください。

### 全ての例を実行する場合

```bash
./runConcurrencyBasic.sh
```

## ディレクトリ構成

- `runConcurrencyBasic.sh` - 全ての例を実行するシェルスクリプト
- `README.md` - このファイル
- [`goroutine/`](goroutine/README.md) - goroutine の基本的な例
- [`channelBasic/`](channelBasic/README.md) - バッファなしチャンネルの例
- [`channelBuffered/`](channelBuffered/README.md) - バッファ付きチャンネルの例
- [`channelRangeClose/`](channelRangeClose/README.md) - range と close を使った例
- [`channelSelect/`](channelSelect/README.md) - select 文を使った例
- [`mutex/`](mutex/README.md) - sync.Mutex を使った例

各ディレクトリには個別の README.md があります。詳細は各ディレクトリの README.md を参照してください。
