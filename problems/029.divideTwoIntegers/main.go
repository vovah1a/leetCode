package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-1, 1))
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	signM, absM := analysis(dividend)
	signN, absN := analysis(divisor)
	res, _ := d(absM, absN, 1)

	if signM != signN {
		res = res - res - res
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func analysis(num int) (sign, abs int) {
	if num < 0 {
		return -1, num - num - num
	}
	return 1, num
}

func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}

		return res, remainder

	}
}
