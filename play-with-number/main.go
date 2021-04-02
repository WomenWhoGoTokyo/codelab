package main

import "fmt"

func main() {

	fmt.Print("数字を入力してください\n")

	var num int
	fmt.Scan(&num)

	fmt.Printf("この数字は%s\n", PrimeNumdeterminer(num))

}
