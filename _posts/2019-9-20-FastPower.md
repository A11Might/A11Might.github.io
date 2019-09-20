---
layout: post
title: 数值的整数次方
tags: [codinginterview]
---

剑指offer第16题

## 1. 数值的整数次方 [题解](https://github.com/A11Might/codingInterview/blob/master/code/offer16.java)

> 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
>
> 保证base和exponent不同时为0

最简单的思路就是迭代的乘base exponent次(O(n)(n等于exponent次))，但O(n)的方法只能处理10 ^ 8级别的数据，leetcode上的说明是 `n是32位有符号整数，其数值范围是[−2 ^ 31, 2 ^ 31 − 1]`，所以这个方法是过不了oj的

#### a. 快幂算法

- a ^ n = (a ^ (n / 2)) * (a ^ (n / 2)) if n 为偶数

- a ^ n = (a ^ ((n - 1) / 2)) * (a ^ ((n - 1) / 2)) * a if n 为奇数

```java
public double Power(double base, int exponent) {
        if (exponent == 0) {
            return 1;
        }
        if (exponent == 1) {
            return base;
        }
        // 区别正数次幂和负数次幂
        boolean flag = exponent > 0 ? true : false;
        exponent = Math.abs(exponent);
        // a ^ (n / 2)
        double res = Power(base, exponent >> 1);
        res *= res;
        // n若为奇数，再乘以a
        if ((exponent & 1) == 1) {
            res *= base;
        }

        return flag ? res : 1 / res;
    }
```
#### b. 按位计算

n的二进制最后一位对应的值是a，倒数第二位对应的是a ^ 2，倒数第三位(a ^ 2) ^ 2，以此类推

设结果初始值为res = 1.0，若n的二进制当前位上数为1，则将result乘以该位对应值；若n的二进制当前位上数为0，则将result乘以1，最终结果即为答案

```java
public double Power(double base, int exponent) {
        if (exponent == 0) {
            return 1;
        }
        if (exponent == 1) {
            return base;
        }
        boolean flag = exponent > 0 ? true : false;
        exponent = Math.abs(exponent);
        double res = 1;
        while (exponent > 0) {
            // 累乘每一位上对应的值
            if ((exponent & 1) == 1) {
                res *= base;
            }
            base *= base;
            exponent >>= 1;
        }

        return flag ? res : 1 / res;
    }
```

## 2. 其他

- 快幂的迭代算法真难写，自己写了个[iterator](https://github.com/A11Might/codingInterview/blob/master/code/offer16.java)，感觉很奇怪

- 这个迭代方法[power](https://kingsfish.github.io/2018/12/21/%E5%89%91%E6%8C%87offer-11-%E6%95%B0%E5%80%BC%E7%9A%84%E6%95%B4%E6%95%B0%E6%AC%A1%E6%96%B9/)不对？但可以过牛客oj

- 参考：-出发- [数值的整数次方](数值的整数次方)
