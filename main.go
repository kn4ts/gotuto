package main

import "fmt"

/*  メイン関数  */
func main() {

	var n int = 10		// int型変数の宣言&代入
	ch := make(chan int)	// int型チャネルの宣言

	fmt.Println("Hello")	// 標準出力関数はこんな感じ

	go countFunc(ch,n)	// countFunc関数を並列実行
	
	for {	// 無限ループは条件なしfor文
		recieved_val, ok := <- ch	// チャネルから値取り出し
		fmt.Println(recieved_val, ok)	// 受け取った値を表示

		if ok != true {
			break	// チャネルが閉じられていたら無限ループ抜ける
		}
	}
}

/*  カウント関数の定義  */
func countFunc (ch chan int, n int) {
	for i:=0 ; i<n ; i++ {	// for文はC言語と同じ感じ
		ch <- i
	}
	close(ch)	// チャネルを閉じて終了
}
