package recursion

var (
	ans  [][]int
	path []int
	used []bool
)

func Permute(nums []int) [][]int {
	ans, path = make([][]int, 0), make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	f(nums)
	return ans
}
func f(nums []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		ans = append(ans, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			f(nums)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
