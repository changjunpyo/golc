# [509. Fibonacci Number(Easy)](https://leetcode.com/problems/fibonacci-number/)

## 문제

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,


Example:

```
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
```

Given n, calculate F(n).

Constraints:

- `0 <= n <= 30`

## 풀이

이것에 풀이에는 다양한 풀이가 있다. 
1. 재귀(recursive)를 이용하는 방법
2. 반복문을 이용하는 방법

1번의 경우, 중복되어 계산 하는 값에 대한 값을 저장해서 반환해주지 않으면 O(n)의 시간으로 해결하지 못한다.
즉, DP(dynamic programming)의 memoization을 이용하여 한번 계산된 값은 저장되어야 빠르게 해결 할 수 있다.
재귀의 경우는 낮아봐야 공간복잡도가 O(n)가 나온다. 

따라서 반복문이 좋다고 생각 될 수 있지만, 단점으로는 과정이 한눈에 바로 들어오지 않는, 즉 가독성이 조금 떨어지게 된다. 

풀이에 대한 
   - Time complexity: O(N) 
   - Space Complexity: O(1)

