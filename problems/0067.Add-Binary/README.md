# [67. Add Binary(Easy)](https://leetcode.com/problems/add-binary/)

## 문제

Given two binary strings a and b, return their sum as a binary string.

Example:

```
Input: a = "11", b = "1"
Output: "100"

Input: a = "1010", b = "1011"
Output: "10101"
```

Constraints:

- `1 <= a.length, b.length <= 104`
- `a and b consist only of '0' or '1' characters.`
- `Each string does not contain leading zeros except for the zero itself.`


## 풀이

1. bit by bit manipulation
LSB부터 각각 더해가면서 carry와 함꼐 String을 추가해주고 마지막 결과 값으로 나온 스트링을 reverse하여 리턴하여 준다.

   - Time complexity: O(max(N,M)) 
   - Space Complexity: O(max(N,M))

2. 내부적으로 스트링으로 구현 되어있는 덧셈 연산 사용 or 구.

 풀이.. 생략
