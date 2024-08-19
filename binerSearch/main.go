package main

import (
	"fmt"
)

func BinerSearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func searchRangeLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				res = mid
				break
			}
			right = mid - 1

		}
	}
	return res
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x - 1
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid < x {
			left = mid + 1
			res = mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	return res
}

func main() {
	// 测试代码
	fmt.Println(mySqrt(5))
	fmt.Println("avv")
}
