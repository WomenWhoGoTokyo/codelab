package main

import (
	"flag"
	"fmt"
)

var game string
var num int

func init() {
	flag.StringVar(&game, "game", defaultGame, usage)
}

const (
	defaultGame = ""
	usage       = "ゲームのメニューを選択"
)

func main() {

	flag.Parse()

	// TODO:timeを使って今日の日付を数字に変換する

	switch game {

	case "prime":
		prime()
		fmt.Printf("この数字は%s\n", PrimeNumdeterminer(num))
	default:
		fmt.Printf("オプションを指定してください")
	}
}
