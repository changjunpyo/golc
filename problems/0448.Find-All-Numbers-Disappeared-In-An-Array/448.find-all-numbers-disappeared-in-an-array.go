package problems

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		j := nums[i] - 1

		if i == j || nums[i] == nums[j] {
			continue
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i--
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			ret = append(ret, i+1)
		}
	}
	return ret
}
