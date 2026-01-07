package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "print build info and exit")
		name        = flag.String("name", "world", "name to greet")
	)
	flag.Parse()

	if *showVersion {
		fmt.Print(versionText())
		return
	}

	fmt.Printf("hello, %s!\n", *name)
	fmt.Println()
	fmt.Println("tip: try `hello -version` to see module version/build settings.")
}

func versionText() string {
	var b strings.Builder

	fmt.Fprintf(&b, "go: %s\n", runtime.Version())

	info, ok := debug.ReadBuildInfo()
	if !ok {
		b.WriteString("build info: (not available)\n")
		return b.String()
	}

	writeModuleInfo(&b, info)
	writeVCSInfo(&b, info)
	writeExecutablePath(&b)

	return b.String()
}

func writeModuleInfo(b *strings.Builder, info *debug.BuildInfo) {
	// モジュールモードでビルドされたバイナリでは、次の情報が観察に役立ちます:
	// - info.Path: mainパッケージのimport path
	// - info.Main.Path / Version / Sum: モジュール情報
	fmt.Fprintf(b, "main package: %s\n", info.Path)
	fmt.Fprintf(b, "module: %s\n", info.Main.Path)
	fmt.Fprintf(b, "version: %s\n", info.Main.Version)
	if info.Main.Sum != "" {
		fmt.Fprintf(b, "sum: %s\n", info.Main.Sum)
	}
}

func writeVCSInfo(b *strings.Builder, info *debug.BuildInfo) {
	// VCS関連の設定は、状況に応じて go コマンドがバイナリに埋め込みます。
	// よく出るキーは次の通りです:
	// - vcs.revision
	// - vcs.time
	// - vcs.modified
	//
	// ※実際に入るかどうかは、ビルド方法/場所/リポジトリ状態に依存します。
	settings := make(map[string]string)
	for _, s := range info.Settings {
		if strings.HasPrefix(s.Key, "vcs.") {
			settings[s.Key] = s.Value
		}
	}
	if len(settings) == 0 {
		return
	}

	b.WriteString("vcs:\n")
	keys := make([]string, 0, len(settings))
	for k := range settings {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(b, "  %s=%s\n", k, settings[k])
	}
}

func writeExecutablePath(b *strings.Builder) {
	// どの実体バイナリを実行しているかも出します（PATHで混乱した時の切り分けに有効）。
	if exe, err := os.Executable(); err == nil {
		fmt.Fprintf(b, "executable: %s\n", exe)
	}
}


