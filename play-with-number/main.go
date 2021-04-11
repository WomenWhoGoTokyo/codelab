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
		fmt.Print("数字を入力してください\n")
		fmt.Scan(&num)

	case "todayis":
		num, err = convertTodayToNum()
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Print("オプションを指定してください")
		return
	}
	PrimeNumdeterminer(num)
}

func PrimeNumdeterminer(num int) {
	result := prime.Prime(num)
	switch result {
	case true:
		fmt.Printf("%dは%s\n", num, "素数です")
	default:
		fmt.Printf("%dは%s\n", num, "素数ではありません")
	}
}

func convertTodayToNum() (int, error) {
	t := time.Now()
	var layout = "20060102"

	today, err := strconv.Atoi(t.Format(layout))
	if err != nil {
		return 0, err
	}
	return today, nil
}
