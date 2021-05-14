package problems

func shuffle(nums []int, n int) []int {
	ret := make([]int, n*2)

	for i := 0; i < n; i++ {
		ret[2*i] = nums[i]
		ret[2*i+1] = nums[i+n]
	}
	return ret
}
