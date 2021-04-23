package prime

import (
	"math"
)

func IsPrime(n int) bool {
	squareroot := math.Sqrt(float64(n))
	result := true

	for i := 2; i <= int(squareroot); i++ {
		if n%i == 0 {
			//TODO: 結果の判定をする
			break
		}
	}
	return result
}
