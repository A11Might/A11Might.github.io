---
layout: post
title: 快速排序
tags: [codinginterview]
---

快速排序也是使用分治策略，用partition函数将序列一分为二(O(n))后，将子序列递归排序((1 / n) * ∑[T(k) + T(n - k - 1)])，最后合并有序子序列(O(1))，T(n) = O(n) + (1 / n) * ∑[T(k) + T(n - k - 1)] = O(n * logn)。

## 一、快速排序

#### 1、快速排序的实现

```java
public static void quickSortCore(int[] arr, int lo, int hi) {
        if (lo < hi) {
            int mid = partition(arr, lo, hi);
            quickSortCore(arr, lo, mid - 1);
            quickSortCore(arr, mid + 1, hi);
        }
    }
```

#### 2、partition的实现

```java
public static int partition(int[] arr, int lo, int hi) {
        // 令arr[hi]为pivot
        // 将原数组分为小于pivot、pivot和大于等于pivot三部分
        swap(arr, lo + (int) Math.random() * (hi - lo + 1), hi);
        int small = lo - 1;
        while (lo < hi) {
            if (arr[lo] < arr[hi]) {
                swap(arr, ++small, lo++);
            } else {
                lo++;
            }
        }
        swap(arr, ++small, hi);

        return small;
    }
```

#### 3、[源码](https://github.com/A11Might/SomePracticeCode/blob/master/util/QuickSort.java)

## 二、剑指offer[[40]](https://www.nowcoder.com/practice/6a296eb82cf844ca8539b57c23e6e9bf?tpId=13&tqId=11182&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

最小的k个数 [题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer40.java)

> 输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4

快排的partition函数会将原序列分为左右两个子序列，左边序列都小于pivot，右边序列都大于或等于pivot，当pivot为数组的第k个元素时，数组中pivot及其之前的元素都小于右边序列，即为n个整数中最小的k个数。

```java
public ArrayList<Integer> GetLeastNumbers_Solution(int [] input, int k) {
        if (input == null || input.length == 0 || input.length < k || k == 0) {
            return new ArrayList<>();
        }
        // 在数组中寻找位置为k - 1的pivot
        int start = 0, end = input.length - 1;
        int index = partition(input, start, end);
        while (index != k - 1) {
            if (index < k - 1) {
                start = index + 1;
            } else {
                end = index - 1;
            }
            index = partition(input, start, end);
        }

        // 收集这k个数
        ArrayList<Integer> res = new ArrayList<>();
        for (int i = 0; i <= index; i++) {
            res.add(input[i]);
        }
        return res;
    }
```

上述解法并不是本题的最优解，但可作为快速排序partition的应用。最优解见[题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer40.java)。

## 三、相关

[数组中出现次数超过一半的数字](https://github.com/A11Might/CodingInterview/blob/master/code/offer39.java)
