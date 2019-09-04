--- 
layout: post
title: 从[1, n]的整数中1出现的次数
tsgs: [codinginterview]
---

剑指offer第43题

> 输入一个整数n，求从1到n这n个整数的十进制表示中1出现的次数
> 
> 例如输入12，从1到12这些整数中包含1的数字有1，10，11和12，1一共出现了5次

### 1、解题思路

通过计算n的十进制的每一位上1出现次数之和，来求出整数[1, n]中1出现的次数

记n的十进制上每一位的值为weight，将weight位置作为中间部分，将n划分为前半部分pre和后半部分succ

#### a、weight为个位


举例n为534，weight为个位4时，pre部分为53，succ部分没有

![Crepe](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/onetimes1.jpg){: .center-block :}

pre部分有53种情况，分别为[0, 52] (单独考虑53的情况)，succ部分只有1种情况(没有东西)，所以当前位置weight为1时至少有53种情况(53 * 1)

上述统计完[0, 529]中个位出现1的次数，接着统计[530, 534]中个位出现1的次数，即考虑当pre部分为53的情况

当pre部分为53时

- 若weight值等于0时，从[530, 530]中，个位也不会出现1，所以当weight为0时(即n为530)，个位上1出现次数为53次

- 若weight值等于1时，从[530, 531]中，531个位会出现1，所以当weight为1时(即n为531)，个位上1出现次数为54次(53 + 1)

- 若weight值大于1时(举例n为534)，从[530, 534]中，531个位会出现1，所以当weight为4时(即n为534)，个位上1出现次数也为54次(53 + 1)



#### b、weight为百位

n为534，当weight为百位3时，pre部分变为5，succ部分变为为4

![Crepe](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/onetimes2.jpg){: .center-block :}

pre部分有5中情况，分别为[0, 4] (单独考虑5的情况)，succ部分有10种情况，分别为[0, 9] (即base种情况，base为当前位置weight的权值)，所以当weight为1时至少有50种情况(5 * 10)，接着考虑当pre部分为5的情况

当pre部分为5时

- 若weight值等于0时，从[500, 504]中，百位不再会出现1，所以当weight为0时(即n为504)，个位上1出现次数为50次

- 若weight值等于1时，从[500, 514]中，百位会出现5次1(succ + 1, 即[0, succ])，分别为510、511、512、513、514，所以当weight为1时(即n为531)，个位上1出现次数为55次(50 + 5)

- 若weight值大于1时(举例534)，从[500, 534]中，百位会出现10次1(base，即[0, base - 1]，base为当前位置weight的权值)，分别为510、511、512、513、514、515、516、517、518、519，所以当weight为3时(即n为534)，个位上1出现次数也为60次(50 + 10)

#### c、weight为更高位

更高位的计算方式与十位相同

#### d、总结

![Crepe](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/onetimes3.jpg){: .center-block :}

把weight作为中间部分，将n分割为前半部分pre和后半部分succ，前半部分有pre + 1中情况，分别为[0, pre]；后半部分有base种情况，分别为[0, base - 1] (base为当前位置weight的权值)，所以当前位置weight上出现1的次数至少为pre * base

{: .box-note}
当前半部分等于pre时，weight位置出现1的次数由weight的决定，所以需要单独讨论

前半部分等于pre时

- 若weight值等于0，weight位不会再出现1，则weight上出现1的总次数为pre * base

- 若weight值等于1，weight位上会再出现1，出现次数为succ + 1次，则weight上出现1的总次数为pre * base + succ + 1

- 若weight值大于1，weight位上会再出现1，出现次数为base次，则weight上出现1的总次数为pre * base + base


### 2、[题解](https://github.com/A11Might/CodingInterview/blob/master/code/offer43.java)
```java
public int NumberOf1Between1AndN_Solution(int n) {
        if (n < 1) {
            return 0;
        }
        int pre = n, weight = 0, base = 1, count = 0;
        while (pre > 0) {
            weight = pre % 10;
            pre /= 10;
            // weight位置出现1的次数至少为pre * base
            count += pre * base; 
            // weight等于1时，weight位上会再出现succ + 1次1
            if (weight == 1) {
                count += (n % base) + 1;
            // weight大于1时，weight位上会再出现base次1
            } else if (weight > 1) {
                count += base;
            }
            // weight等于0时，weight位上不会再出现1

            // 计算下一个位置weight的权值
            base *= 10;
        }

        return count;
    }
```

### 3、时间复杂度

由分析思路或者代码都可以看出，while循环的次数就是n的位数logn(以10为底)(n能被10除logn(以10为底)次)，而循环体内执行的操作都是有限次的，所以时间复杂度为O(logn)

### 4、参考

- yi_afly：[从1到n整数中1出现的次数：O(logn)算法](https://blog.csdn.net/yi_Afly/article/details/52012593)
