package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func prime() int {
	fmt.Print("数字を入力してください\n")

	fmt.Scan(&num)
	return num
}

func PrimeNumdeterminer(num int) (result string) {
	var divisor []int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisor = append(divisor, i)
		}
	}

	switch len(divisor) {
	case 1:
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
