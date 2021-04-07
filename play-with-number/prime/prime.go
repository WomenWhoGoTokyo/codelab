package prime

func Prime(n int) bool {
	var divisor []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisor = append(divisor, i)
		}
	}

	switch len(divisor) {
	case 1:
		return true
	default:
		return false
	}

}
