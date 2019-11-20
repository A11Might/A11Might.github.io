---
layout: post
title: 股票问题
subtitle: 解决一系列股票问题最为通用的方法
tags: [leetcode, dynamic programming]
---

每一个股票问题都有一个最优解, 但是它们之间往往没有关联, 这使得我们很难找到一个统一的方法来解决它们. 在这里我介绍一种最一般化的解答, 当然细化到具体问题时需要根据具体题目作出相应的修改.

### I. 简单的例子

> 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格, 问你可以获得的最大利润是多少?

对于这个问题, 你可能会问"最大利润取决于我们从哪天开始, 我们被允许进行多少次交易". 当然, 这些因素会在具体的题目中给出. 但还有一个隐藏因素不知道你考虑到没有, 具体如下:

首先我们约定符号来简化我们的描述, 使用 prices 来表示长度为 n 的股票价格数组, i 表示第 i 天(i 在 [0, n) 之间), k 表示最大被允许的交易次数, T[i][k] 表示第 i 天最多交易 k 次最终可以获得的最大利润. 显然base case是: T[-1][k] = T[i][0] = 0, 代表没有股票或者没有交易次数来产生利润, 获得的最大利润为0(0 表示第一天, 所以 -1 表示没有股票). 现在如果我们可以以某种方式关联 T[i][k] 和它的子问题, 比如 T[i - 1][k], T[i][k - 1], T[i - 1][k - 1], ..., 我们就可以获得一个有效的递推关系来解决这个问题, 所以我们要如何获得呢?

最直接的方法就是观察第 i 天的操作. 我们有几种操作? **答案是三种: 买, 卖和休息**. 那我们应该采取哪种操作? 答案是: 不知道, 但是想知道话很简单. 我们可以尝试每一种操作, 然后选出获得利润最大的操作, 当然这是在没有限制条件的情况下. 然而我们有一个额外的限制条件, 不能同时参加多次交易, 即若想在第 i 天买入股票, 在购买前你的手上不能持有股票; 若想在第 i 天卖出股票, 在卖出前你的手上必须持有股票. 我们手中是否持有股票, 就是上文提到的隐藏因素, 它将影响第 i 天的操作从而影响获得的最大利润.

因此我们应该将 T[i][k] 的定义分为两个部分: T[i][k][0] 和 T[i][k][1], 前者表示在第 i 天最多交易 k 次执行完操作后不持有股票最终可以获得的最大利润, 后者表示在第 i 天最多交易 k 次执行完操作后持有股票最终可以获得的最大利润. 现在 base case 和递推关系可以写成如下形式:

- base case:

    T[-1][k][0] = 0, T[-1][k][1] = -Infinity
    T[i][0][0] = 0, T[i][0][1] = -Infinity

- recurrence relations:

    T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])
    T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k - 1][0] - prices[i])

对于 base case, T[-1][k][0] = T[i][0][0] = 0 意思和之前一样, 而 T[-1][k][1] = T[i][0][1] = -Infinity 强调的是在没有股票或者不允许交易的情况下是不可能持有股票的.

对于 T[i][k][0] 的递推关系: 我们最终不持有股票的情况下, 第 i 天只能进行休息或者卖出这两种操作. T[i - 1][k][0] 表示第 i 天休息获得的最大利润, 而 T[i - 1][k][1] + prices[i] 表示第 i 天卖出股票获得的最大利润. 注意: 这里的最大允许的交易次数是相同的, 这是因为一次交易包含买入和卖出两种操作, 仅仅执行买入操作时才会改变了最大允许的交易次数.

{: .box-note}
**Note:** 买入和卖出这一组操作是一次交易: 买入操作开启了一次新的交易, 需要消耗一次交易次数, 而卖出操作只是完成了上次已经进行了的交易, 不消耗交易次数.

对于 T[i][k][1] 的递推关系: 我们最终持有股票的情况下, 第 i 天只能进行休息或者买入这两种操作. T[i - 1][k][1] 表示表示第 i 天休息获得的最大利润, 而 T[i - 1][k - 1][0] - prices[i] 表示第 i 天买入股票获得的最大利润. 注意: 这里的最大允许的交易次数要减一, 因为在第 i 天进行买入操作时会使用一次交易次数, 解释同上.

这样我们可以简单的通过循环遍历 prices 数组并根据上述的递推关系来更新 T[i][k][0] 和 T[i][k][1] 的值, 进而得到在最后一天最终可以获得的最大利润. 最后的答案将会是 T[i][k][0] (最终不持有股票时我们会获得更大的利润).

### II. 具体示例的应用

按照允许的最大交易次数将 6 个股票问题分类(最后两个问题含有附加要求如冷冻期或者手续费), 我将把通用解法应用到它们每一题上.

#### case I: k = 1

对于这个例子, 我们每天有两个未知的变量: T[i][1][0] 和 T[i][1][1], 它们的递推关系为: 

```T[i][1][0] = max(T[i - 1][1][0], T[i - 1][1][1] + prices[i])```

```T[i][1][1] = max(T[i - 1][1][1], T[i - 1][0][0] - prices[i]) = max(T[i - 1][1][1], -pirces[i])```

由于只能交易一次, 所以在买入股票前不会有其它交易, 即买入前不可能产生利润(T[i - 1][0][0] = 0), 可以化简第二个等式.

通过上述的两个等式, 可以轻松的写出一个时间复杂度 O(n) 空间复杂度 O(n) 的解法. 但若你注意到在第 i 天获得的最大利润实际上仅仅依赖于第 i - 1 天的两个变量, 你可以将空间复杂度降到 O(1). 如下是空间最优解:

```java
public int maxProfit(int[] prices) {
    int T_i10 = 0, T_i11 = Integer.MIN_VALUE;
        
    for (int price : prices) {
        T_i10 = Math.max(T_i10, T_i11 + price);
        T_i11 = Math.max(T_i11, -price);
    }
        
    return T_i10;
}
```

现在让我们深入了解上面的解法. 仔细观察循环体的内部, 会发现 T_i11 只是代表 0 到第 i 天中所有股票价格的负值中的最大值, 换句话说就是, 所有股票价格中的最小值. 至于 T_i10 就是, 在*第 i - 1 天获得的最大利润*和*以0 到第 i - 1 天中所有股票价格中的最小值买入后在以第 i 天的价格卖出获得的最大利润*中选出较大值, 即为第 i 天可以获得的最大利润.

#### case II: k = 正无穷

当 k 趋于正无穷时 k 等于 k - 1, 这意味 T[i - 1][k - 1][0] = T[i - 1][k][0] 和 T[i - 1][k - 1][1] = T[i - 1][k][1]. 每天的两个未知变量: T[i][k][0] 和 T[i][k][1] 在 k 趋于正无穷时, 它们的递推关系为:

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])```

```T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k - 1][0] - prices[i]) = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])```

由 T[i - 1][k - 1][0] = T[i - 1][k][0] 可以化简第二个等式.

如下是时间复杂度 O(n) 空间复杂度 O(1) 的解法:

```java
public int maxProfit(int[] prices) {
    int T_ik0 = 0, T_ik1 = Integer.MIN_VALUE;
    
    for (int price : prices) {
        int T_ik0_old = T_ik0;
        T_ik0 = Math.max(T_ik0, T_ik1 + price);
        T_ik1 = Math.max(T_ik1, T_ik0_old - price);
    }
    
    return T_ik0;
}
```

这个解法其实是使用贪心策略来获取最大利润: 在每个区间低价处购买股票，其后立刻在这个区间高价处出售(只要今天的价格大于昨天的价格就在昨天买入今天卖出, 累加可得获得的最大利润), 这相当于在 prices 中寻找上升子序列, 以每个子序列开始的价格买入再以最后的价格卖出.

#### case III: k = 2

和例 I: k = 1 类似, 只不过现在我们每天有四个变量: T[i][1][0], T[i][1][1], T[i][2][0], T[i][2][1], 它们的递推关系为: 

```T[i][2][0] = max(T[i - 1][2][0], T[i - 1][2][1] + prices[i])```

```T[i][2][1] = max(T[i - 1][2][1], T[i - 1][1][0] - prices[i])```

```T[i][1][0] = max(T[i - 1][1][0], T[i - 1][1][1] + prices[i])```

```T[i][1][1] = max(T[i - 1][1][1], -prices[i])```

由 base case T[i][0][0] = 0, 可以化简最后一个等式. 时间复杂度O(n) 和 空间复杂度 O(1) 的解法如下: 

```java
public int maxProfit(int[] prices) {
    int T_i10 = 0, T_i11 = Integer.MIN_VALUE;
    int T_i20 = 0, T_i21 = Integer.MIN_VALUE;
        
    for (int price : prices) {
        T_i20 = Math.max(T_i20, T_i21 + price);
        T_i21 = Math.max(T_i21, T_i10 - price);
        T_i10 = Math.max(T_i10, T_i11 + price);
        T_i11 = Math.max(T_i11, -price);
    }
        
    return T_i20;
}
```

#### case IV: k = 任意整数

这是一个最一般化的例子, 我们需要更新每一天不同 k 值所对应的持有股票或者不持有股票最终可以获得的最大利润. 在这里我们可以进行一个小的优化, 当 k 超过临界值时, 最大利润不再依赖允许的交易次数, 但它还是会和总天数有关, 让我们来找出这个临界值.

一次赚钱的交易至少需要两天时间(一天买入股票, 另一天卖出股票, 保证买入价格要比卖出价格低). 如果总天数是 n (给定数组的长度), 那么赚钱的交易次数最多为 n / 2 (整除). 在那之后的交易是无法获利的, 这就意味着最大利润不会发生变化. 因此 k 的临界值为 n / 2. 如果给定的 k 不小于这个值, 即 k >= n / 2, 我们可以将 k 增大到正无穷, 这样问题就变成了例 II.

如下为时间复杂度O(kn), 空间复杂度O(k)的解答. 若不进行优化, 代码将会在较大的 k 值时, 发成TLE.

```java
public int maxProfit(int k, int[] prices) {
    if (k >= prices.length >>> 1) {
        int T_ik0 = 0, T_ik1 = Integer.MIN_VALUE;
    
        for (int price : prices) {
            int T_ik0_old = T_ik0;
            T_ik0 = Math.max(T_ik0, T_ik1 + price);
            T_ik1 = Math.max(T_ik1, T_ik0_old - price);
        }
        
        return T_ik0;
    }
        
    int[] T_ik0 = new int[k + 1];
    int[] T_ik1 = new int[k + 1];
    Arrays.fill(T_ik1, Integer.MIN_VALUE);
        
    for (int price : prices) {
        for (int j = k; j > 0; j--) {
            T_ik0[j] = Math.max(T_ik0[j], T_ik1[j] + price);
            T_ik1[j] = Math.max(T_ik1[j], T_ik0[j - 1] - price);
        }
    }
        
    return T_ik0[k];
}
```

#### case V: k = 正无穷但含冷冻期

这个例子与例 II 非常相似, 它们有相同的 k 值, 但是由于有冷冻期, 所以需要稍稍修改递推关系. 例 II 的递推关系为:

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])```

```T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])```

由于冷冻期的限制, 在第 i - 1 天卖出股票后, 我们不能在第 i 天买入股票. 因此在第二个等式中, 当我们想要在第 i 天买入股票时, 应该使用 T[i - 2][k][0] 代替 T[i - 1][k][0]. 其余部分同例 II 相同, 新的递推关系为: 

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])```

```T[i][k][1] = max(T[i - 1][k][1], T[i - 2][k][0] - prices[i])```

如下是时间复杂度 O(n) 空间复杂度 O(1) 的解答:

```java
public int maxProfit(int[] prices) {
    int T_ik0_pre = 0, T_ik0 = 0, T_ik1 = Integer.MIN_VALUE;
    
    for (int price : prices) {
        int T_ik0_old = T_ik0;
        T_ik0 = Math.max(T_ik0, T_ik1 + price);
        T_ik1 = Math.max(T_ik1, T_ik0_pre - price);
        T_ik0_pre = T_ik0_old;
    }
    
    return T_ik0;
}
```

#### case VI: k = 正无穷但含手续费

这个例子也与例 II 非常相似, 它们有相同的 k 值, 但是由于有手续费, 所以需要稍稍修改递推关系. 例 II 的递推关系为:

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])```

```T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])```

现在我们需要为每一笔交易支付手续费(使用符号 fee 表示), 即在第 i 天买入或卖出股票获得的利润中扣除这个费用, 因此新的递推关系为:

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])```

```T[i][k][1] = max(T[i - 1][k][1], T[i-1][k][0] - prices[i] - fee)```

或者

```T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i] - fee)```

```T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k][0] - prices[i])```

注意: 在扣除手续费时我们有两种选择. 这是因为(上文也说过了)每一笔交易都是一组买入和卖出这两个操作. 我们可以在买入股票时支付手续费(上面的两个等式), 也可以在卖出的时候支付手续费(下面的等式). 对应两个时间复杂度 O(n) 空间复杂度 O(1) 的解法, 分别如下: 

- 解法一: 在买入股票时支付手续费

```java
public int maxProfit(int[] prices, int fee) {
    int T_ik0 = 0, T_ik1 = Integer.MIN_VALUE;
    
    for (int price : prices) {
        int T_ik0_old = T_ik0;
        T_ik0 = Math.max(T_ik0, T_ik1 + price);
        T_ik1 = Math.max(T_ik1, T_ik0_old - price - fee);
    }
        
    return T_ik0;
}
```
- 解法二: 在卖出股票时支付手续费

```java
public int maxProfit(int[] prices, int fee) {
    int T_ik0 = 0, T_ik1 = Integer.MIN_VALUE;
    
    for (int price : prices) {
        int T_ik0_old = T_ik0;
        T_ik0 = Math.max(T_ik0, T_ik1 + price - fee);
        T_ik1 = Math.max(T_ik1, T_ik0_old - price);
    }
        
    return T_ik0;
}
```

### III. 总结

总的来说, 股票问题的通用解法可以使用一个三维的dp-table表示, 每个维度分别为天数 i, 最大允许的交易次数 k 和在最后一天是否持有股票. 我已经给出了最大利润的递推关系和终止条件, 也在最后给出了可以AC的题解.

Hope this helps and happy coding!

### IV. 其他

- 原文 - [fun4LeetCode的leetcode高赞解](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/most-consistent-ways-of-dealing-with-the-series-of-stock-problems)

- [LeetCode股票问题](https://github.com/A11Might/leetcode/blob/master/navigate/dp.md)
