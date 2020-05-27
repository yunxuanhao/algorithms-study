package leetcode

// https://leetcode-cn.com/problems/powx-n/

// func myPow(x float64, n int) float64 {
// 	var start float64
// 	start = 1.0
// 	if n > 0 {
// 		for i := 0; i < n; i++ {
// 			start *= x
// 		}
// 	} else if n < 0 {
// 		x = 1 / x
// 		for i := 0; i < -n; i++ {
// 			start *= x
// 		}
// 	}
// 	return start
// }

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
