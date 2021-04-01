package main

func primenumdeterminer(num int) (result string) {
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
