package problems

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	uniqueCount := 0

	for i:=1; i<len(nums); i++{
		if nums[uniqueCount] != nums[i]{
			nums[uniqueCount+1] = nums[i]
			uniqueCount++
		}
	}
	return uniqueCount+1
}