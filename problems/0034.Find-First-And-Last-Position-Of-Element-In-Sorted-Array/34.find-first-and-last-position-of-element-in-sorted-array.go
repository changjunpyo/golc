package problems

func binarySearchLowerbound(target int, nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}
func binarySearchUpperbound(target int, nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if target >= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != target {
		return l - 1
	}
	return l
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := []int{-1, -1}

	low := binarySearchLowerbound(target, nums)
	up := binarySearchUpperbound(target, nums)
	if low != -1 {
		ans[0] = low
		ans[1] = up
	}

	return ans
}
