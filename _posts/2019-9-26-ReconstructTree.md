---
post: layout
title: 二叉树的重构
tag: [algorithm, leetcode, codinginterview]
---



## 1. 重构

任何一个二叉树都可以导出三个序列，分别是先序、中序和后序遍历序列，这三个序列长度相同，都是由树中所有节点，依照对应的遍历策略确定的次序排列而成。若已知树的遍历序列，如何忠实的还原出树的拓扑结构。

### a. [先序 | 后序] + 中序

*先序或后序遍历序列 + 中序遍历序列可以忠实的还原出树的拓扑结构*

数学归纳法证明可行

- 假设任意规模n < N的二叉树这个规律都成立

- 当n == N时   

    - 先序遍历序列首个节点是根节点r，紧跟其后的是左子树对应的遍历序列，以及右子树对应的遍历序列  

    - 中序遍历序列，左子树对应的遍历序列是前缀，右子树对应的遍历序列是后缀，根节点r在它们之间

    - 由先序遍历序列可知树的根节点，进而在中序遍历序列中对该节点进行定位，找到左子树对应的中序遍历子序列和右子树对应的中序遍历子序列，即可以知道左子树和右子树分别有哪些节点组成，反过来将先序遍历序列中左子树和右子树对应的子序列切分开，这样就将全树的重构问题化简为两棵子树的重构问题，这两课子树在规模上都符合归纳假设，严格小于N，因此根据归纳假设左子树和右子树都可以如此重构出来，假设成立

![Crepe](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/reconstruct1.jpg)

#### 剑指offer[7]/leetcode[105] [题解](https://github.com/A11Might/codingInterview/blob/master/code/offer07.java)

> 从前序与中序遍历序列构造二叉树
>
> 根据一棵树的前序遍历与中序遍历构造二叉树，你可以假设树中没有重复的元素

```java
public class Solution {
    public TreeNode reConstructBinaryTree(int [] pre,int [] in) {
        HashMap<Integer, Integer> inMap = new HashMap<>();
        for (int i = 0; i < in.length; i++) {
            inMap.put(in[i], i);
        }
        return reConstructBinaryTreeCore(inMap, pre, in, 0, pre.length - 1, 0, in.length - 1);
    }

    private TreeNode reConstructBinaryTreeCore(HashMap<Integer, Integer> inMap, int[] pre, int[] in, int preStart, int preEnd, int inStart, int inEnd) {
        if (preStart > preEnd || inStart > inEnd) {
            return null;
        }
        int value = pre[preStart];
        TreeNode node = new TreeNode(value);
        int indexInInorder = inMap.get(value);
        int offset = indexInInorder - inStart; // 左子树节点个数
        node.left = reConstructBinaryTreeCore(inMap, pre, in, preStart + 1, preStart + offset, inStart, indexInInorder - 1);
        node.right = reConstructBinaryTreeCore(inMap, pre, in, preStart + offset + 1, preEnd, indexInInorder + 1, inEnd);
        return node;
    }
}
```

{: .box-note}
**Note:** 想当然的将后序遍历序列中根节点的位置直接映射到前序遍历序列中，进而切分左右子树时错误的，应该根据子树的个数切分左右子树

### b. 先序 + 后序

左子树和右子树可能有空树，树的规模为0

若一棵树的左子树或右子树是空的，则无法通过先序和后序遍历序列判断根节点前或后是左子树还是右子树，所以不能进行准确的重构

### c. [先序 + 后序] × 真二叉树

*先序遍历序列 + 后序遍历序列可以重构出真二叉树*

非退化的真二叉树，左子树和右子树要么同时为空，要么同时非空，前一种情况显而易见，不妨假设它们都存在，其先序遍历序列首先出现的是根节点，其后是以左子树根节点引领的左子树遍历子序列，再接下来右子树根节点引领的右子树遍历子序列；而后序遍历序列必然是以根节点收尾，往前是以右子树根节点收尾的右子树遍历子序列，再往前以左子树树根收尾的左子树的遍历子序列

![Crepe](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/reconstruct2.jpg)

- 左子树的根节点l是先序遍历序列中的第二个节点，在任何给定的先序遍历序列中都可以快速的找到它，进而在后序遍历序列中对其定位，l节点在它所属的左子树的后序遍历子序列中垫后，这样可以明确的在后序遍历序列中界定左右子树的范围

- 对称的，在后序遍历序列中右子树的树根位置也是确定的，在先序遍历序列中对右子树的根节点r定位，同样可以确定左右子树的切分位置

- 通过递归的形式重构出一棵真二叉树原本的结构

#### leetcode[889] [题解](https://github.com/A11Might/leetcode/blob/master/codes/lc889.java)

> 根据前序和后序遍历构造二叉树
>
> 返回与给定的前序和后序遍历匹配的任何二叉树，pre 和 post 遍历中的值是不同的正整数

```java
class Solution {
    public TreeNode constructFromPrePost(int[] pre, int[] post) {
        Map<Integer, Integer> postMap = new HashMap<>();
        for (int i = 0; i < post.length; i++) {
            postMap.put(post[i], i);
        }
        return process(postMap, pre, post, 0, pre.length - 1, 0, post.length - 1);
    }

    private TreeNode process(Map<Integer, Integer> postMap, int[] pre, int[] post, int preStart, int preEnd, int postStart, int postEnd) {
        if (preStart > preEnd || postStart > postEnd) {
            return null;
        }
        TreeNode node = new TreeNode(pre[preStart]);
        if (preStart == preEnd || postStart == postEnd) {
            return node;
        }
        // 由于寻找左1需向后偏移一位，若当前[preStart, preEnd]只有一个元素，需特殊处理
        // 需在使用pre[preStart + 1]之前直接创建返回，防止数组越界错误
        // 这是和前序中序、中序后序的区别
        int nodeleftIndexInPostorder = postMap.get(pre[preStart + 1]);
        int offset = nodeleftIndexInPostorder - postStart + 1; // 左子树节点个数
        node.left = process(postMap, pre, post, preStart + 1, preStart + offset, postStart, nodeleftIndexInPostorder);
        node.right = process(postMap, pre, post, preStart + offset + 1, preEnd, nodeleftIndexInPostorder + 1, postEnd - 1);
        return node;
    }
}
```
## 2. 其他

- [leetcode树的重构总结](https://github.com/A11Might/leetcode/blob/master/category.md)

- 参考：邓俊辉 - 数据结构
