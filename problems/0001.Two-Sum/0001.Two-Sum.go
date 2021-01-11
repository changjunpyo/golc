package problems


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i:= 0; i<len(nums); i++{
		val, exist := m[nums[i]]
		if exist{
			return []int{val, i}
		}
		m[target - nums[i]] = i
	}
	return []int{-1, -1}
}