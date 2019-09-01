---
layout: post
title: 字符串表达式求解
tags: [codinginterview]
---

给定一个字符串str，str表示一个公式，公式里可能有整数，加减乘除符号和左右括号，返回公式的计算结果

> 1、给定字符串一定是正确的公式，即不需要对str做公式的有效性检测
>
> 2、若是负数，就需要有括号括起来，但若负数作为公式的开头或括号部分的开头，则可以没有括号
>
> 3、不用考虑计算过程中发生溢出的情况

## 一、分析

#### 1、简化题目

先考虑没有括号的情况，即给定一个字符串str，str表示一个公式，公式里可能有 `正整数` 和加减乘除但 `没有左右括号` ，返回公式的计算结果

由简化后的题目，假设计算的是 `1 + 2 * 3 - 8 / 4`，乘除的运算优先级高于加减，先算乘除得 `1 + 6 - 2`，后算加减得 `5`

#### 2、计算乘除

从前往后遍历字符串中每个字符，在每次记录值和所遇符号时，需要判断前面符号是否为'*'和'/'来先行计算乘除，使用栈这种结构，可以很好的实现这种操作

遍历字符串，若当前字符为数字则开始累计其值，遇到符号时即为当前值累计结束，在将值和所遇字符压入栈中之前，判断前面符号是否为'*'和'/'来先行计算乘除，然后重新开始累计其值，直至遍历完整个字符串，最后栈中只有整数和加减符号

> 特殊的，给定字符串str表示的公式为 `- 1 + 3`
>
> 在遇到符号 '-' 时，会将值和 '-' 压入栈中，但前面并没有值，所以压入的是默认值0，相当于 `0 - 1 + 3`

```java
public static int value(char[] str, int index) {
        Deque<String> deque = new ArrayDeque<>();
        int curNum = 0;
        while (index < str.length) {
            char curChr = str[index++];
            if ('0' <= curChr && curChr <= '9') {
                curNum = curNum * 10 + (curChr - '0');
            } else { 
                addNum(deque, curNum); 
                deque.addLast(String.valueOf(curChr)); 
                curNum = 0;
            }
        }
        addNum(deque, curNum); // 最后一个值没有符号位终止累计，需另行加入

        return getNum(deque);
}

// 将当前得到的值放入栈中，并将之前的*或/先行运算掉
public static void addNum(Deque<String> deque, int num) {
    if (!deque.isEmpty()) {
        String top = deque.pollLast();
        if (top.equals("*") || top.equals("/")) {
            int preNum = Integer.valueOf(deque.pollLast());
            num = top.equals("*") ? (preNum * num) : (preNum / num);
        } else {
            deque.addLast(top);
        }
    }
    deque.addLast(String.valueOf(num));
}
```

#### 3、计算加减

计算完乘除后，原str中的字符已经转化为值和符号，计算加减使，直接从前往后计算即可

```java
public static int getNum(Deque<String> deque) {
        int res = Integer.parseInt(deque.pollFirst());
        while (!deque.isEmpty()) {
            String operator = deque.pollFirst(); 
            int num = Integer.parseInt(deque.pollFirst());
            if (operator.equals("+")) {
                res += num;
            } else {
                res -= num;
            }
        }

        return res;
}
```

#### 4、计算加减plus

经过value()处理后得到的队列队首一定是数字(若str是以符号 '-' 开头，则返回的队列队首为0；若str是以数字开头，则返回的队列队首为数字)，如 `-1 + 3`，队首为 '-1'，使用getNum()计算时没有问题，但若给定的队列的为 `- 1 + 3`，注意不是`-1 + 3`，队列的第一个元素为符号 '-'， getNum将无法处理，这时可以学习value()的操作，创造 `0 - 1 + 3`，具体如下：

> `-1 + 3`，队首为符号时，处理为`0 - 1 + 3`
>
> `1 + 3`， 队首为数字时，处理为`0 + 1 + 3`

```java
public static int getNumPlus(Deque<String> deque) {
        int res = 0;
        boolean add = true;
        while (!deque.isEmpty()) { 
            String cur = deque.pollFirst();
            if (cur.equals("+")) {
                add = true;
            } else if (cur.equals("-")) {
                add = false;
            } else {
                int num = Integer.valueOf(cur);
                res += add ? num : (-num);
            }
        }

        return res;
    }
```

## 二、题解

1、双端队列

![_config.yml]({{ site.baseurl }}/images/deque.png)

2、[题解](https://github.com/A11Might/SomePracticeCode/blob/master/highClass/GetValue.java)

再回过头来看题目，给定一个字符串str，str表示一个公式，公式里可能有整数，加减乘除符号和左右括号，返回公式的计算结果，将原公式看做一个大括号中的公式，中间可能嵌套着其他括号，在遇到嵌套的括号时，递归调用value()计算括号内公式的值，直接返回替换掉括号内的公式即可

```java
public static int getValue(String str) {
        char[] chrs = str.toCharArray();
        return value(chrs, 0)[0];
    }

    // 返回值为含两个值的数组，bra[0] 括号中元素运算后的值，bra[1] 运算到第几位
    private static int[] value(char[] str, int index) {
        Deque<String> deque = new ArrayDeque<>();
        int curNum = 0;
        while (index < str.length && str[index] != ')') {
            char curChr = str[index++];
            if ('0' <= curChr && curChr <= '9') {
                curNum = curNum * 10 + (curChr - '0');
            } else if (curChr != '(') {
                addNum(deque, curNum);
                deque.addLast(String.valueOf(curChr));
                curNum = 0;
            } else {
                int[] bra = value(str, index); // 递归计算括号内的表达式
                curNum = bra[0];
                index = bra[1] + 1; 
            }
        }
        addNum(deque, curNum);

        return new int[] {getNum(deque), index};
    }
```

## 三、参考

- 2018高级算法课 - 左神
