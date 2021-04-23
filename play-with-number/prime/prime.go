package prime

import (
	"math"
)

func IsPrime(n int) bool {
	squareroot := math.Sqrt(float64(n))
	result := true

	//TODO: コマンドラインから取得した値が1の場合の処理を書く
	for i := 2; i <= int(squareroot); i++ {
		if n%i == 0 {
			//TODO: 結果の判定をする
			break
		}
	}
	return result
}
