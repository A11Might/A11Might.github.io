---
layout: post
title: 排列与组合
tag: [backtrack, leetcode]
---

排列组合相同又不同

## 1. 排列

#### leetcode[[46](https://leetcode-cn.com/problems/permutations/)] [题解](https://github.com/A11Might/leetcode/blob/master/src/lc046.java)

> 给定一个没有重复数字的序列，返回其所有可能的全排列

**示例**

```
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

假设给定的序列是"1, 2, 3", 全排列思路为从1, 2, 3中选出一个数字例如1固定在一号位置, 再从2, 3中选出一个数字例如2固定在二号位置, 再从3中选出一个数字例如3固定在三号位置, 可以得到一种排列; 若在二号位置上选择的是3, 则三号位置上固定是2, 又可以得到另一种排列; 以此类推

简单来说就是, 先在给定序列中选择一个数字放在当前位置, 再对剩余的数字进行全排列(减而治之), 具体如下:

```java
private void dfs(int[] nums, int index, List<List<Integer>> ans) {
    if (index == nums.length) {
        ans.add(getList(nums));
        return;
    }
    for (int i = index; i < nums.length; i++) {
        swap(nums, index, i); // 为第index位选择一个字符
        dfs(nums, index + 1, ans); // 全排列剩下的字符
        swap(nums, index, i); // 还原数组; 回溯: 恢复之前的状态, 重新做决定
    }
}
```

{: .box-note}
**Note:** 使用swap函数, 将从序列中选出的数字(选出要放置在当前位置的数字)与当前位置的数字交换位置, 这样的可以方便剩余数字的全排列(全排列数组[index + 1, length)即可), 最后得到的结果即为一种排列

## 2. 组合

#### leetcode[[77](https://leetcode-cn.com/problems/combinations/)] [题解](https://github.com/A11Might/leetcode/blob/master/src/lc077.java)

> 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合

**示例**

```
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

与排列类似, 组合也可以先在给定序列中选择一个数字放在当前位置, 再对剩余的数字进行组合(减而治之), 但这样会出现重复的组合: abc和acb是两种排列, 但它们是一种组合, 即若两组序列的组成字符相同, 则它们是一种组合, 解决如下:

在每次递归调用dfs时, 只考虑组合当前使用数字(i)之后的数字(当前数字之前的数字, 在之前的递归或在本次循环中已经考虑过了), 可以避免出现重复组合, 具体如下:

```java
private void dfs(int n, int index, int k, List<Integer> sublist, List<List<Integer>> ans) {
        if (sublist.size() == k) {
            ans.add(new ArrayList<>(sublist));
            return;
        }
        // 剪枝: i <= n - (k - sublist.size()) + 1
        for (int i = index; i <= n - (k - sublist.size()) + 1; i++) { 
            sublist.add(i); // 为当前位置选择一个字符
            dfs(n, i + 1, k, sublist, ans); // i以前的数字已经考虑过了, 组合i之后的数字
            sublist.remove(sublist.size() - 1); // 还原列表; 回溯: 恢复之前的状态, 重新做决定
        }
    }
```

{: .box-note}
**剪枝:** 还有k - sublist.size()个空位，所以[1, n]中至少要有k - sublist.size()个元素, i最多为n - (k - sublist.size()) + 1(再大就没有足够的空位来找到k个数的组合)

## 3. 其他

- [leetcode排列组合总结]()
