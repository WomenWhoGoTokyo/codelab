package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/WomenWhoGoTokyo/codelab/play-with-number/prime"
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
		getNum()
		fmt.Printf("この数字は%s\n", PrimeNumdeterminer(num))
	case "todayis":
		fmt.Printf("今日の日付は%s\n", TodaysDatedeterminer())
	default:
		fmt.Printf("オプションを指定してください")
	}
}

func getNum() int {
	fmt.Print("数字を入力してください\n")

	fmt.Scan(&num)
	return num
}

func PrimeNumdeterminer(num int) (msg string) {
	result := prime.Prime(num)

	switch result {
	case true:
		return "素数です"
	default:
		return "素数ではありません"
	}
}

func TodaysDatedeterminer() (result string) {

	t := time.Now()
	const layout = "20060102"

	today, err := strconv.Atoi(t.Format(layout))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("本日は%v\n", today)

	return PrimeNumdeterminer(today)
}
