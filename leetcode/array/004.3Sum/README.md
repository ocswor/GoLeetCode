#[#3Sum](https://leetcode-cn.com/problems/3sum/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china)    

##题目
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
###Note:
The solution set must not contain duplicate triplets.

###Example:
```bigquery
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

##题目大意：
给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。


###解题思路

暴力解法 应该是三层for循环 组成三个不同的组合，判断三个数之和的等于0

不过这里可以降维，先将任意两个数之和保存起来，可以将时间复杂度降到O(n^2)


最后还要去重，组合的三个数字不能相同
这一题比较麻烦的一点在于，最后输出解的时候，要求输出不重复的解。数组中同一个数字可能出现多次，同一个数字也可能使用多次，
但是最后输出解的时候，不能重复。例如 [-1，-1，2] 和 [2, -1, -1]、[-1, 2, -1] 这 3 个解是重复的，
即使 -1 可能出现 100 次，每次使用的 -1 的数组下标都是不同的。