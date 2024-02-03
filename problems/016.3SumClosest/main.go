package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSumClosest(nums, 3))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum < target:
				l++
				if delta > target-sum {
					delta = target - sum
					res = sum
				}
			case sum > target:
				r--
				if delta > sum-target {
					delta = sum - target
					res = sum
				}
			default:
				return sum
			}
		}
	}

	return res
}
