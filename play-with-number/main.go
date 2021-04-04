package main

import (
	"flag"
	"fmt"
)

var game string

func main() {

	flag.Parse()

	// TODO:timeを使って今日の日付を数字に変換する

	if game != "" {
		fmt.Print("数字を入力してください\n")

		var num int
		fmt.Scan(&num)

		fmt.Printf("この数字は%s\n", PrimeNumdeterminer(num))
	} else {
		fmt.Printf("オプションを指定してください")
	}
}

func init() {

	flag.StringVar(&game, "game", "", "ゲームのメニューを選択")
}
