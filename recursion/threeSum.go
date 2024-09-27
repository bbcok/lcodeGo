package recursion

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else if nums[l]+nums[r]+nums[i] < 0 {
				l++
			} else {
				ans = append(ans, []int{nums[l], nums[r], nums[i]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return ans
}
