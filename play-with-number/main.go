package main

import "fmt"

func main() {

	fmt.Print("数字を入力してください\n")

	var num int
	fmt.Scan(&num)

	fmt.Printf("この数字は%s\n", judge(num))

}

func judge(num int) (result string) {
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
