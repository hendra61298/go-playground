package main

import (
	"fmt"
	"sort"
)

func largestPerimeterV1(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 3, 4, 5}))

}
