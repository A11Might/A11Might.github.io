---
layout: post
title: 字符串和字符流
tags: [codinginterview]
---

剑指offer第50题

### 1、剑指offer[50][[题目一]](https://www.nowcoder.com/practice/1c82e8cf713b4bbeb2a5b31cf5b0417c?tpId=13&tqId=11187&tPage=2&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking)

第一个只出现一次的字符 [题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer50.java)

> 在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）

每个字符都对应一个特定的ascii码(标准ascii字符[0, 127])，用大小为128的整型数组的索引代表每个字符，统计每个字符的词频，再次遍历字符串，即可找到第一个词频为1的字符

有点类似[剑指offer[3]](https://www.nowcoder.com/practice/623a5ac0ea5b4e5f95552655361ae0a8?tpId=13&tqId=11203&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking)找出数组中任意一个重复的数字，使用大小为n(n为数组中最大的数字)的数组统计词频，不同的是本题使用的额外空间复杂度为O(1)(数组大小固定为128)，并且需要遍历两次字符串，第一次用于统计每个字符的词频，第二次用于找到第一个词频为1的字符(题3只需遍历一次数组，在统计词频时，出现词频为2的数字即可返回)

```java
public int FirstNotRepeatingChar(String str) {
    int n = str.length();
    if (str == null || n == 0) {
        return -1;
    }
    int[] times = new int[128];
    // 统计每个字符的词频
    for (int i = 0; i < n; i++) {
        times[(int) str.charAt(i)]++;
    }
    // 找到第一个词频为1的字符
    for (int i = 0; i < n; i++) {
        if (times[(int) str.charAt(i)] == 1) {
            return i;
        }
    }

    return -1;
}
```

### 2、剑指offer[50][[题目二]](https://www.nowcoder.com/practice/1c82e8cf713b4bbeb2a5b31cf5b0417c?tpId=13&tqId=11187&tPage=2&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking)

字符流中第一个不重复的字符 [题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer502.java)

> 请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"
>
> 如果当前字符流没有存在出现一次的字符，返回#字符

字符串和字符流的区别是，字符串是一次给定的，而字符流是一个一个字符读取的，若要用题目一的方法解决本题，可以在每次读取字符流中字符时累计词频，但在查找第一个只出现一次的字符时，由于需要遍历字符流到目前为止形成的字符串，所以额外空间复杂度O(n)(n为插入字符次数)

这O(n)的额外空间复杂度是可以省去的，操作如下

在统计词频时，将 `当前字符在字符流形成的字符串中的位置` 放入数组中索引为当前字符ascii码的位置(即原来字符的词频换成字符的index)，这样遍历词频数组则可得到字符出现的先后顺序，随后再语义规定，值为0代表当前位置索引对应的字符未出现过；值为-1代表当前位置索引对应的字符重复出现(在插入当前字符index时若该位置已有字符index，则置为-1标记为重复出现)，用于判断当前字符的词频

{: .box-note}
用于记录当前字符在字符流形成的字符串中的位置的全局变量index的初始值应为1，避免与语义约定的0冲突

```java
private int[] occurrence = new int[128];
private int index = 1;

//Insert one char from stringstream
public void Insert(char ch) {
    // 当前字符第一次出现，记录字符index
    if (occurrence[ch] == 0) {
        occurrence[ch] = index;
    // 当前字符重复出现，标记为-1
    } else if (occurrence[ch] > 0) {
        occurrence[ch] = -1;
    }
    index++;
}

//return the first appearence once char in current stringstream
public char FirstAppearingOnce() {
    char res = '#';
    int minIndex = Integer.MAX_VALUE;
    for (int i = 0; i < occurrence.length; i++) {
        if (occurrence[i] > 0 && occurrence[i] < minIndex) {
            res = (char) i;
            minIndex = occurrence[i];
        }
    }

    return res;
}
```
