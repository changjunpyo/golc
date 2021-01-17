# [1. Two Sum(Easy)](https://leetcode.com/problems/two-sum/)

## 문제

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Constraints:

- `2 <= nums.length <= 103`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**

## 풀이

1. Using Map(hash table)

   - 이 문제는 array를 순환하며 target 값과 각각의 값을 뺀 값을 key로 하고, array의 index를 value로 하여 hash table을 만들어 주며 (ex map[target- nums[idx]] = idx) 각각의 값에 접근할 때마다 map을 확인하여 value가 존재하면 value와 현재 index를 반환하면 된다.

   - Time complexity: O(N) 
   - Space Complexity: O(N)

