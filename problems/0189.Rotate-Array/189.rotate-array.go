package problems

// Rotate1  additional memory solution
// time complexity:  O(n) -> copy is O(n)
// space complexity: O(n) -> copy is O(n)
func Rotate1(nums []int, k int)  {
    n := len(nums)
    k = k % n
    rotated := append(nums[n-k:], nums[:n-k]...)
    copy(nums, rotated)
}

func reverse(nums []int) {
	n := len(nums)
	for i:=0; i<n/2; i++{
		nums[i], nums[n-(i+1)] = nums[n-(i+1)], nums[i]
	}
}
// Rotate2 reverse array solution
// time complexity:  O(n) -> reverse is O(n)
// space complexity: O(1)
func Rotate2(nums []int, k int)  {
	n := len(nums)
	k = k % n
	reverse(nums[:n-k])
	reverse(nums[n-k:])
	reverse(nums)
}


// Rotate3 cyclic solution
// time complexity:  O(n)
// space complexity: O(1)
func Rotate3(nums []int, k int)  {
    n := len(nums)
    k = k % n
    count := 0
    for i:=0; count< n; i++{
        cur, prev := i, nums[i]
        for {
            next := (cur + k) % n
            temp := nums[next]
            nums[next] = prev
            prev = temp
            cur = next
            count++
            if cur == i{
                break
            }
        }
    }
}


