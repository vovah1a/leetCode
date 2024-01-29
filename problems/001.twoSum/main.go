package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{i, hm[num]}
		}

		hm[target-num] = i
	}

	return nil
}
