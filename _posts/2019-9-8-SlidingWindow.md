---
layout: post
title: 滑动窗口
tags: algorithm
---

### 1、主动滑动窗口

不断向右移动滑动窗口的右边界，并在每次窗口扩大时，维持窗口内所有元素满足条件(即调整左边界)

#### 无重复字符的最长子串 [leetcode[3]](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) 

> 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度

```java
public int lengthOfLongestSubstring(String s) {
    if (s == null || s.length() == 0) {
        return 0;
    }
    int[] chrIndex = new int[128];
    Arrays.fill(chrIndex, -1);
    int n = s.length();
    int l = 0, r = -1; // 滑动窗口[l, r]，初始r为-1，表示窗口中无元素
    int res = 0;
    while (r + 1 < n) {
        r++; // 不断扩大窗口
        if (chrIndex[s.charAt(r)] != -1) { // "abba"
            // 更新窗口时，只更新了窗口的左右边界，更新后窗口外之前窗口的值未变化(即窗口外的值仍有注册)
            // 所以当前元素在chrIndex中有出现时，需要判断想要更新的成为的左边界是否在窗口内，来决定是否更新窗口
            //          a、在窗口外说明已过期，无须更新
            //          b、在窗口内则任有效，更新左边界
            l = Math.max(l, chrIndex[s.charAt(r)] + 1);
        }
        res = Math.max(res, r - l + 1);
        chrIndex[s.charAt(r)] = r;
    }
    
    return res;
}
```

### 2、被动滑动窗口

通过判断窗口中所有元素是否满足条件，来不断滑动窗口(即调整窗口左右边界)

#### 长度最小的子数组 [leetcode[209]](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

> 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组
>
> 如果不存在符合条件的连续子数组，返回 0

```java
public int minSubArrayLen(int s, int[] nums) {
    if (nums == null || nums.length == 0) {
        return 0;
    }
    int n = nums.length;
    int l = 0, r = -1; // 滑动窗口[l, r]，初始r为-1，表示窗口中无元素
    int sum = 0; // 当前窗口中元素的和
    int res = n + 1; // 当前和>=正整数s的最短连续子数组长度
    while (l < n) {
        // 窗口中元素和小于s，则移动右边界扩大窗口，增加窗口和
        if (r + 1 < n && sum < s) {
            sum += nums[++r];
        // 窗口中元素和大于等于s，则移动左边界缩小窗口，寻找下一个和>=s的窗口
        } else {
            sum -= nums[l++];
        }
        
        // 实时更新和>=正整数s的最短连续子数组长度
        if (sum >= s) {
            res = Math.min(res, r - l + 1);
        }
    }
    
    // 不存在符合条件的连续子数组，返回 0
    if (res == n + 1) {
        return 0;
    }
    return res;
}
```

### 3、参考

- 玩转算法面试 - 波波老师
