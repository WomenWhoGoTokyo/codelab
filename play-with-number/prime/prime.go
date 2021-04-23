package prime

import (
	"math"
)

func IsPrime(n int) bool {
	squareroot := math.Sqrt(float64(n))
	result := true

	for i := 2; i <= int(squareroot); i++ {
		//TODO: コマンドラインから取得した値が1の場合の条件分岐処理を書く
		if n%i == 0 {
			//TODO: 結果の判定をする
			break
		}
	}
	return result
}
