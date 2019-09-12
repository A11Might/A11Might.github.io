---
layout: post
title: 二维数组中的查找
tags: [codinginterview]
---

剑指offer第4题

#### 二维数组中的查找 [题解](https://github.com/A11Might/codingInterview/blob/master/code/offer04.java)

> 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序
>
> 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数

最简单的思路就是遍历矩阵的每一个元素(O(m ^ n))，来判断其中是否有目标元素，但这样没有用到题目中元素有序的条件

题意每一行的元素从左到右递增，每一列元素从上到下递增

从左下角(右上角同理)开始遍历矩阵，若当前元素小于目标元素，当前元素所在列中不可能有目标元素(当前元素所在列未遍历元素都小于当前元素，也都小于目标元素)，则向右移动索引继续寻找目标元素；若目标元素小于当前元素，当前元素所在行中不可能有目标元素(当前元素所在行未遍历元素都大于当前元素，也都大于目标元素)，则向上移动索引继续寻找目标元素

如上每次比较都pass掉一行或一列元素，大大加速遍历过程，时间复杂度为(O(m + n))

```java
public boolean Find(int target, int [][] array) {
        int row = array.length - 1, col = array[0].length - 1;
        // 从左下角向右上角遍历矩阵
        int curRow = array.length - 1, curCol = 0;
        while (curRow >= 0 && curCol <= col) {
            if (array[curRow][curCol] < target) {
                curCol++;
            } else if (target < array[curRow][curCol]) {
                curRow--;
            } else {
                return true;
            }
        }

        return false;
    }
```
