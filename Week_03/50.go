/*
50. Pow(x, n)
https://leetcode-cn.com/problems/powx-n/
*/
//2020/9/25

package main

func main() {

}

//暴力 超时
func myPow1(x float64, n int) float64 {
	call := func(x float64, n int) float64 {
		rel := 1.0
		for n > 0 {
			rel = rel * x
			n--
		}
		return rel
	}

	if n < 0 {
		return 1 / call(x, -n)
	}

	return call(x, n)
}

//分治
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
