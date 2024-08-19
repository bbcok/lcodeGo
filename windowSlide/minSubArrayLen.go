package windowSlide

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum, ans := 0, int(^uint(0)>>1)
	for r < len(nums) {
		sum += nums[r]
		for sum-nums[l] >= target {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}
		}
		r++
	}
	if ans == int(^uint(0)>>1) {
		return 0
	}
	return ans
}

// func main() {
// 	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
// 	fmt.Println(res)
// }
