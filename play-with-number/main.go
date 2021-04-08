package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/WomenWhoGoTokyo/codelab/play-with-number/prime"
)

var game string

const (
	defaultGame = ""
	usage       = "ゲームのメニューを選択"
)

func init() {
	flag.StringVar(&game, "game", defaultGame, usage)
}

func main() {

	flag.Parse()

	var (
		num int
		err error
	)

	switch game {

	case "prime":
		// 数字を標準入力で取得する（numに判定したい数字が入る）
		fmt.Print("数字を入力してください\n")
		fmt.Scan(&num)

	case "todayis":
		// 本日の日付を数字に変換する（numに判定したい数字が入る）
		num, err = convertTodayToNum()
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Print("オプションを指定してください")
		return
	}
	//数字を判定した結果を文字列で出力する（共通処理）
	PrimeNumdeterminer(num)
}

// 素数判定の結果を出力
func PrimeNumdeterminer(num int) {
	result := prime.Prime(num)
	switch result {
	case true:
		fmt.Printf("%dは%s\n", num, "素数です")
	default:
		fmt.Printf("%dは%s\n", num, "素数ではありません")
	}
}

// 本日の日付をintegerに変換
func convertTodayToNum() (int, error) {
	t := time.Now()
	var layout = "20060102"

	today, err := strconv.Atoi(t.Format(layout))
	if err != nil {
		return 0, err
	}
	return today, nil
}
