package problems

func containsDuplicate(nums []int) bool {
	check := make(map[int]bool, 0)
	for _, val := range nums{
		if check[val] == true{
			return true
		}
		check[val] = true
	}
	return false
}