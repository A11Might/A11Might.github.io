---
layout: post
title: Manacher算法
tags: [algorithm]
---

Manacher算法，又叫“马拉车”算法，可以在时间复杂度为 O(n) 的情况下求解一个字符串的最长回文子串的长度，其本质是通过存储的之前遍历字符的回文信息，加速后序字符寻找回文子串的过程，即优化暴力解法(它们都是求每个字符的回文半径，Manacher加速了这个过程)

## 一、Manacher算法

字符串 str 长度为 n

a. 有一个辅助数组 pArray 表示回文半径数组，用于存储每一个字符的回文半径

b. 有两个辅助变量，一个 R 表示最右回文半径的右边界，一个 C 表示最早到达最右回文右边界时所对应的回文中心

#### 1. 字符串预处理

遍历字符串的每个字符，以每个字符为中心向外扩，得到每一个字符的回文半径，例如：12321

但这种做法只能用于奇回文，对偶回文进行操作时会得不到正确答案，例如：偶回文 1221

解决方法：在每个字符前后加上一个相同的字符(任意字符)例如#，则原字符串变为 #1#2#2#1# ，它的最长回文子串的一半即为原字符串，这种方法可以同时用于奇偶回文

![_config.yml]({{ site.baseurl }}/images/manacher01.JPG)

#### 2. manacherString()代码实现

```java
private static char[] manacherString(String str) {
		char[] charArr = str.toCharArray();
		char[] res = new char[str.length() * 2 + 1]; // 相当于在每个字符后面加一个'#',再在第一个字符前加一个'#'
		int index = 0; // 当前字符位置
		for (int i = 0; i < res.length; i++) {
			res[i] = (i & 1) == 0 ? '#' : charArr[index++]; // i为res数组中的位置，若为偶数则放置'#'，奇数则放置当前字符(
		}
		return res;
	}
```

#### 3. 主函数

a. 当前字符所在位置 i 在最右回文半径的右边界 R 的上或右边

无法加速，从当前字符所在位置向外暴力扩

b. 当前字符所在位置 i 在最右回文半径的右边界 R 的左边

> i' 为 i 关于最早到达最右回文右边界时所对应的回文中心 C 的对称位置
>
> L 为 R 关于最早到达最右回文右边界时所对应的回文中心 C 的对称位置

- i' 位置字符的回文在 L 和 R 之间时，i 位置字符为中心的回文长度等于 i' 位置字符回文长度

![_config.yml]({{ site.baseurl }}/images/manacher02.JPG)

*你可能要问，为啥不能更长了能？*

L 到 R 为 C 位置的回文串，所以 X == X' 并且 Y == Y'，若 X == Y ，则 X' == Y'，i' 位置的回文串可以更长，说明之前所得 i' 位置的回文串长度有错，不可能所以不能更长了

- i' 位置字符的回文在 L 和 R 之外时，i 位置字符为中心的回文长度等于 i' 位置到 L 位置的距离

![_config.yml]({{ site.baseurl }}/images/manacher03.JPG)

*你可能又要问，为啥不能更长了能？*

L 到 R 为 C 位置的回文串，所以 X == X' 并且 Y != Y'，L' 到 R' 为 i' 位置的回文串，所以 X' == Y'，由上可得 X == X' == Y' != Y，即 X != Y，所以不能更长了

- i' 位置字符的回文在 L 上时，i 位置字符为中心的回文长度至少等于 i' 位置到 L 位置的距离，再往外无法判断，需继续向外扩

![_config.yml]({{ site.baseurl }}/images/manacher04.JPG)

*你不要问了，它可能更长！*

#### 4. maxLcpsLength()代码实现

```java
private static int maxLcpsLength(String str) {
		if (str == null || str.length() == 0)
			return 0;
		char[] charArr = manacherString(str); // 字符串预处理
		int[] pArr = new int[charArr.length]; // 回文半径数组
		int C = -1; // 最早到达最右回文右边界时所对应的回文中心
		int R = -1; // 最右回文半径的右边界
		int max = Integer.MIN_VALUE; // 最大回文半径
		for (int i = 0; i != charArr.length; i++) {
            // 两种情况：
            // a. i不在R内，需暴力扩，i位置的回文子串至少是它自己即1
            // b. i在R内，则i位置的回文长度至少为pArr[2 * C - i](i' 位置字符的回文在L和R之间)
            //    和R - i(i' 位置字符的回文在L和R之外)中较小值，即使i'位置字符的回文在L上也是在
            //    上述值上继续增加的，只需额外再向外扩一次即可判断具体是哪种情况
            pArr[i] = R > i ? Math.min(pArr[2 * C - i], R - i) : 1; // i' 位置为c - (i - c)即2 * c - i
            // 暴力扩时，pArr[i]初始值为1，依次向外扩增加值
            // 在之前加速的基础上，继续外扩
			while (i + pArr[i] < charArr.length && i - pArr[i] > -1) {
				if (charArr[i + pArr[i]] == charArr[i - pArr[i]]) {
					pArr[i]++;
				} else {
					break;
				}
			}
			// 实时更新R 和 C
			if (i + pArr[i] > R) {
				R = i + pArr[i];
				C = i;
			}
			// 实时更新max
			max = Math.max(max, pArr[i]);
		}
		return max - 1;
	}
```

#### 5. [Manacher代码](https://github.com/A11Might/SomePracticeCode/blob/master/highClass/Manacher.java)

## 二、实战

#### 实例1

> 给你一个原始串，只能在后面添加字符形成大串，要求形成回文串。返回最短的大串

当 R 扩到字符串最后一个字符时停止，将 R 关于 C 的对称位置 L 前的所有字符逆序，加到字符串后面，即可组成最短回文字符串

[code](https://github.com/A11Might/SomePracticeCode/blob/master/highClass/ManacherShorestEnd.java)

## 三、参考

- 2018高级算法课 - 左神
