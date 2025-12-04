package main

import (
	"fmt"
	"sync"
	"time"
)

// worker は Worker Pool パターンで使用するワーカー関数です
// この関数は goroutine として実行され、jobs チャネルからジョブを受け取って処理します
//
// 引数:
//   - id: ワーカーの識別子（1, 2, 3...）
//   - jobs: 処理するジョブを受け取るチャネル（受信専用 <-chan）
//   - results: 処理結果を送信するチャネル（送信専用 chan<-）
//   - wg: ワーカーの終了を待つための WaitGroup
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// defer で WaitGroup の Done() を呼び出す
	// この関数が終了する時（正常終了でもエラーでも）に必ず呼ばれる
	defer wg.Done()

	// jobs チャネルからジョブを受信し続けるループ
	// jobs チャネルが close されると、この range ループが自動的に終了する
	for j := range jobs {
		fmt.Printf("Worker %d: ジョブ %d を処理開始\n", id, j)

		// 処理時間を擬似的に再現するため、100ミリ秒待機
		// 実際の処理では、ここでデータベースアクセスやAPI呼び出しなどを行う
		time.Sleep(100 * time.Millisecond)

		// シンプルな処理：整数を 2 倍にする
		// 実際の処理では、ここでビジネスロジックを実行する
		result := j * 2

		fmt.Printf("Worker %d: ジョブ %d の処理完了（結果: %d）\n", id, j, result)

		// 処理結果を results チャネルに送信
		// バッファ付きチャネルなので、受信側が準備できていなくても送信できる
		results <- result
	}

	// jobs チャネルが close されて、全てのジョブを処理し終わった
	fmt.Printf("Worker %d: 全てのジョブを処理完了し、終了します\n", id)
}

func main() {
	// ========================================
	// 1. 設定とチャネルの準備
	// ========================================
	const numJobs = 10   // 処理するジョブの数
	const numWorkers = 3 // 並行して動作するワーカーの数

	// jobs チャネル: ワーカーにジョブを送るためのチャネル
	// バッファサイズを numJobs にすることで、全てのジョブを一度に送信できる
	// バッファがないと、ワーカーが受信するまで送信がブロックされる
	jobs := make(chan int, numJobs)

	// results チャネル: ワーカーから処理結果を受け取るためのチャネル
	// バッファサイズを numJobs にすることで、全ての結果を一度に受信できる
	results := make(chan int, numJobs)

	// WaitGroup: 複数のワーカーの終了を待つための同期プリミティブ
	var wg sync.WaitGroup

	// ========================================
	// 2. ワーカーを起動
	// ========================================
	fmt.Printf("=== %d 個のワーカーを起動します ===\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		// WaitGroup のカウンターを 1 増やす
		// これにより、このワーカーが終了するまで待つ必要があることを WaitGroup に伝える
		wg.Add(1)

		// ワーカーを goroutine として起動
		// 各ワーカーは jobs チャネルからジョブを受信し、results チャネルに結果を送信する
		// goroutine として起動するので、main 関数の処理は続行される（非ブロッキング）
		go worker(w, jobs, results, &wg)
		fmt.Printf("Worker %d: 起動しました\n", w)
	}

	// ========================================
	// 3. ジョブを投入
	// ========================================
	fmt.Printf("\n=== %d 個のジョブを投入します ===\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		// ジョブを jobs チャネルに送信
		// バッファ付きチャネルなので、バッファに空きがあればブロックしない
		// ワーカーが並行して処理しているので、ジョブはすぐに処理される
		jobs <- j
		fmt.Printf("ジョブ %d を投入しました\n", j)
	}

	// ========================================
	// 4. jobs チャネルを close
	// ========================================
	// jobs チャネルを close することで、ワーカー側の range ループが終了する
	// close しないと、ワーカーは jobs チャネルから新しいジョブを待ち続けてしまう
	// 全てのジョブを投入した後、close することで「これ以上ジョブはない」ことを伝える
	fmt.Println("\n=== 全てのジョブを投入完了。jobs チャネルを close します ===")
	close(jobs)

	// ========================================
	// 5. 処理結果を受信
	// ========================================
	// ワーカーが処理した結果を results チャネルから受信する
	// numJobs 個のジョブを投入したので、numJobs 個の結果が返ってくる
	fmt.Printf("\n=== 処理結果を受信します ===\n")
	for r := 1; r <= numJobs; r++ {
		// results チャネルから結果を受信
		// ワーカーが並行して処理しているので、結果は処理完了順に返ってくる（順序は保証されない）
		result := <-results
		fmt.Printf("結果 %d: %d\n", r, result)
	}

	// ========================================
	// 6. 全てのワーカーの終了を待つ
	// ========================================
	// WaitGroup で全てのワーカーの終了を待つ
	// 各ワーカーは jobs チャネルが close されると range ループが終了し、worker 関数が終了する
	// worker 関数が終了すると defer wg.Done() が実行され、WaitGroup のカウンターが減る
	// 全てのワーカーが終了すると、カウンターが 0 になり、wg.Wait() のブロックが解除される
	fmt.Println("\n=== 全てのワーカーの終了を待機中... ===")
	wg.Wait()

	fmt.Println("All workers finished")
}
