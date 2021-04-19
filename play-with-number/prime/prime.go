package prime

import (
	"math"
)

func Prime(n int) bool {
	squareroot := math.Sqrt(float64(n))
	result := true

	for i := 2; i <= int(squareroot); i++ {
		if n%i == 0 {
			result = false
			break
		}
	}
	return result
}
