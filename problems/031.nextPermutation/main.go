package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	last := len(nums) - 1
	peakIndex, largeIndex := last, last

	for ; peakIndex > 0; peakIndex-- {
		if nums[peakIndex-1] < nums[peakIndex] {
			break
		}
	}

	if peakIndex != 0 {
		for ; largeIndex >= peakIndex; largeIndex-- {
			if nums[peakIndex-1] < nums[largeIndex] {
				nums[largeIndex], nums[peakIndex-1] = nums[peakIndex-1], nums[largeIndex]
				break
			}
		}
	}

	for i := last; i > peakIndex; i, peakIndex = i-1, peakIndex+1 {
		nums[i], nums[peakIndex] = nums[peakIndex], nums[i]
	}
}
