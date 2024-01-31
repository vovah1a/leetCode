package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	fmt.Println(reverse(x))
}

var max = int(math.Pow(2, 31) - 1)
var min = int(max * (-1))

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	isNegative := x < 0
	if isNegative {
		x *= -1
	}
	y := x % 10
	x /= 10
	for x >= 10 {
		y, x = y*10+x%10, x/10
	}
	y = y*10 + x
	if y > max || y <= min {
		return 0
	}
	if isNegative {
		return y * (-1)
	}
	return y
}
