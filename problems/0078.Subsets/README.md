# [78. Subsets](https://leetcode.com/problems/subsets/)

## 문제

Given an integer array `nums`, return *all possible subsets (the power set)*.

The solution set must not contain duplicate subsets.

Example:

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

- **Constraints:**
  - `1 <= nums.length <= 10`
  - `-10 <= nums[i] <= 10`
  - All the numbers of `nums` are **unique**.

## 풀이

- backtracking을 이용하여 모든 경우의 수를 다 찾아 탐색하여 array를 복사하여 추가하는 방식으로 풀이 할 수 있다.
- **Time Complexity** : **O(N*2^N)** , 각각의 모든 경우의 수를 탐색하게 되는데 2^N이 나오고 그 순간마다 array(최대 N개)를 복사하여 넣어주게 되기 때문.
- **Space Complexity**: **O(N*2^N)** subsets의 개수가 2^N개 이고 각각의 subset마다 N개 이하의 array 길이를 가지고 있다. 

