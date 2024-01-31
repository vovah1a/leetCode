package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		min, isAMax := getMinHeight(height[i], height[j])
		sum := (j - i) * min
		if sum > max {
			max = sum
		}

		if isAMax {
			j--
		} else {
			i++
		}
	}
	return max
}

func getMinHeight(a, b int) (int, bool) {
	if a > b {
		return b, true
	}
	return a, false
}
