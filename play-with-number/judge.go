package main

import "fmt"

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
