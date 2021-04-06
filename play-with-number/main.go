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

	switch game {

	case "prime":
		prime()
		fmt.Printf("この数字は%s\n", PrimeNumdeterminer(num))
	case "todayis":
		fmt.Printf("今日の日付は%s\n", TodaysDatedeterminer())
	default:
		fmt.Printf("オプションを指定してください")
	}
}
