package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	max := findMax(nums)
	fmt.Println("Max value:", max)
}

func findMax(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
