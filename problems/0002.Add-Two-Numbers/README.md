# [2. Add Two Numbers(Medium)](https://leetcode.com/problems/add-two-numbers/)

## 문제

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

- **Constraints:**
  - The number of nodes in each linked list is in the range `[1, 100]`.
  - `0 <= Node.val <= 9`
  - It is guaranteed that the list represents a number that does not have leading zeros.

## 풀이

- 문제에서 이미 계산을 하기 쉽게 거꾸로 되어있는 순서로 되어 있다. 각각의 linkedList를 순회하며 더하고 새로운 linkedList를 만들어 노드에 추가 해주면 된다. 추가적으로 10이 넘어가는 경우 carry 값을 1, 새로운 node에 %10을 하여 값을 넣어주는 방식으로 진행하면 된다.
- Time Complexity : O(N) , 각각의 노드를 한번씩 순회하므로 
- Space Complexity: O(N) 

