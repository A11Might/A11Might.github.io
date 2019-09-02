---
layout: post
title: 区间[0, n - 1]和区间[1, n - 1]
tags: [codinginterview]
---

区间[0, n - 1]和区间[1, n - 1]只一数之差？

## 1、剑指offer[3][[题目一]](https://www.nowcoder.com/practice/623a5ac0ea5b4e5f95552655361ae0a8?tpId=13&tqId=11203&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)

数组中重复的数字 [题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer03.java)

> 在一个长度为n的数组里的所有数字都在[0, n - 1]的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任意一个重复的数字
> 
> 例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，那么对应的输出是第一个重复的数字2

这道题最简单的做法，是统计出现数字的词频，词频大于1的即为重复数字，但缺点很明显，需要O(n)(n为数字的最大值)的额外空间复杂度

思考，长度为n的数组里的所有数字都在[0, n - 1]的范围内，将每个数字放置在与其值相等的索引位置来排序数组，若没有重复数字，每个位置对应的值应为其索引的值，若有重复数字，某些索引位置就可能对应多个数字，即为重复的数字，具体如下

遍历数组，将当前遍历到的数字放置在数组中索引与其值相等的位置，若该位置已有正确对应的数字，则当前数字为重复的数字

```java
public boolean duplicate(int numbers[],int length,int [] duplication) {
    if (numbers == null || numbers.length == 0) {
        return false; 
    }
    for (int i = 0; i < length; i++) {
        while (i != numbers[i]) { // 使用while，swap后继续判断当前位置的数字
            // 找到重复的数字
            if (numbers[numbers[i]] == numbers[i]) {
                duplication[0] = numbers[i];
                return true;
            // 未找到重复的数字，将当前数字放置在数组中索引与其值相等的位置
            } else {
                swap(numbers, i, numbers[i]);
            }
        }
    }

    return false;
}
```

## 2、剑指offer[3][题目二]

不修改数组找到重复的数字 [题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer032.java)

> 在一个长度为n的数组里的所有数字都在[1, n - 1]的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入数组
> 
> 例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的输出是重复的数字2或者3

长度为n的数组里的所有数字都在[1, n - 1]的范围内，若将这n个数字放入区间[1, n - 1]内，由于区间所含的位置数比数字个数少一个，所以至少含有一个重复的数字

将区间[1, n - 1]从中间数字m处，拆分成两部分[1, m]和[m + 1, n - 1]，若[1, m]区间的数字的数目超过m，则[1, m]区间一定含有重复数字；否则[m + 1, n - 1]区间一定含有重复数字，然后继续把一定含有重复数字的区间一分为二，直至找到一个重复数字

{: .box-note}
这种方法只能找到一组重复数字(题目二中那个至少含有的一个重复的数字)，无法找出其他重复的数字(题目一中的重复数字)

```java
public boolean duplicate(int numbers[],int length,int [] duplication) {
    if (numbers == null || numbers.length == 0) {
        return false; 
    }
    // 在[1, n - 1]区间进行类似二分搜素的操作
    int start = 1, end = length - 1;
    while (start <= end) { // O(logn)
        int mid = start + ((end - start) >> 1);
        int count = countRange(numbers, start, mid); // O(n)
        if (start == end) {
            if (count > 1) {
                duplication[0] = start;
                return true;
            } else {
                break;
            }
        }
        if (count > (mid - start + 1)) {
            end = mid;
        } else {
            start = mid + 1;
        }
    }

    return false;
}

// 统计数组numbers中，数值在[start, end]之间的数字的个数
private int countRange(int[] numbers, int start, int end) {
    int count = 0;
    for (int num : numbers) {
        if (start <= num && num <= end) {
            count++;
        }
    }

    return count;
}
```
