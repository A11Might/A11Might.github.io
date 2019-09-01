---
layout: post
title: 二分查找
tags: [algorithm]
---

对于有序数组我们可以使用二分查找，时间复杂度为T(N) = T(n / 2) + O(1) = O(log N)，大大优于顺序查找时间复杂度O(N)。

### 1. 特殊情况

> 目标元素不存在
>
> 目标元素存在多个

### 2. 语义约定

例如：需要在目标元素后面插入新元素

- 即使失败，也应该给出新元素适当的插入位置

- 若允许重复元素，则每一组也需要按其插入次序排列

约定：在有序数组区间内[lo, hi)中， 返回不大于target的最后一个元素的下标

- 当有多个命中元素时，必须返回最靠后元素的下标

- 失败时，应返回小于target的最大者(含哨兵lo - 1)

### 3. 实现

```java
public static int binarySearch(int[] nums, int target) {
        int lo = 0, hi = nums.length;
        while (lo < hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (target < nums[mid]) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }
        return lo - 1;
    }
```

### 4. 正确性

#### a. 不变性：nums[0, lo) <= target < nums[hi, n)

- 初始时，lo = 0且hi = n，nums[0, lo) = nums[hi, n) = 空集，自然成立

- 数学归纳法：假设不变性一直保持至(a)，以下无非两种情况

![_config.yml]({{ site.baseurl }}/images/binarySearch1.png)

#### b. 单调性：显而易见

- 最终lo = hi，将有序数组分为：<= target 和 target <，返回lo - 1即为不大于target的最后一个元素的下标

![_config.yml]({{ site.baseurl }}/images/binarySearch2.png)

### 5. leetcode 34

在排序数组中查找元素的第一个和最后一个位置


> 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置
>
> 你的算法时间复杂度必须是 O(log n) 级别
>
> 如果数组中不存在目标值，返回 [-1, -1]

思路：二分查找目标元素 - 1的下标加一和目标元素的下标即为在排序数组中查找元素的第一个和最后一个位置

```java
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int start = binarySearch(nums, target);
        // 当查找失败时，返回的是左侧哨兵节点 -1，或右侧哨兵节点的左临 n - 1
        if (start == -1 || nums[start] != target) {
            return new int[] {-1, -1};
        }
        // 二分查找目标元素 - 1的下标加一和目标元素的下标即为在排序数组中查找元素的第一个和最后一个位置
        return new int[] {binarySearch(nums, target - 1) + 1, start};
    }

    // 二分查找返回不大于目标元素的最后一个元素
    private int binarySearch(int[] nums, int target) {
        int left = 0, right = nums.length;
        while (left < right) {
            int mid = left + ((right - left) >> 1);
            if (target < nums[mid]) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left - 1;
    }
}
```

### 6. 参考

- 数据结构 - 邓俊辉

- leetcode高赞解 - [vision57](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/discuss/14701)
