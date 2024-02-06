package main

import "fmt"

func main() {
	nums := []int{3, 5, 1}
	target := 1
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r-2 {
		k := (l + r) / 2

		switch {
		case nums[l] < nums[k] && nums[k] >= target && target >= nums[l]:
			r = k
		case nums[r] > nums[k] && target <= nums[r] && target >= nums[k]:
			l = k
		case nums[l] > nums[k] && (target <= nums[k] || target >= nums[l]):
			r = k
		case nums[r] < nums[k] && (target <= nums[r] || target >= nums[k]):
			l = k
		default:
			return -1
		}
	}

	if target == nums[r] {
		return r
	}
	if target == nums[l] {
		return l
	}
	return -1
}
