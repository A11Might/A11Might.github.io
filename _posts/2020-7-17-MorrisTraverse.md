---
layout: post
title: Morris 遍历
tags: [algorithm]

---

当我们想要遍历二叉树的时候，可以使用一个栈来存储树中节点进而进行遍历，但不管是迭代还是递归（系统栈），都会使用 O(n) 的额外空间复杂度，那可不可以将这些空间省去呢，答案是肯定的。

二叉树中的节点含有空闲的空指针，在遍历节点时我们可以将特定的空指针指向当前遍历到的节点，用于遍历完当前节点的左子树之后返回当前节点，这样就将使用栈遍历二叉树所需 O(n) 的额外空间复杂度降为 O(1)。

#### Morris 遍历

上述的方法就是 **Morris 遍历**，具体操作如下：

将当前节点记为 cur

```java
TreeNode cur = root;
```

- 如果 cur 没有左孩子，那么 cur 向右移动

  ```java
  if (cur.left == null) cur = cur.right;
  ```

- 如果 cur 有左孩子，找到 cur 左孩子最右的节点，记为 mostRight

  ```java
  TreeNode mostRight = cur.left;
  while (mostRight.right != null && mostRight.right != cur) {
      mostRight = mostRight.right;
  }
  ```

  - 如果 mostRight.right 指针指向 null，说明第一次来到当前节点 cur，左子树还没遍历过，进入左子树遍历：让其指向 cur，cur 向左移动

    ```java
    if (mostRight.right == null) {
        mostRight.right = cur;
        cur = cur.left;
    }
    ```

  - 如果 mostRight.right 指针指向 cur，说明第二次来到当前节点 cur，左子树已被遍历过：让其指回 null，cur 向右移动

    ```java
    if (mostRight.right == cur) {
        mostRight.right = null;
        cur = cur.right;
    }
    ```

- 知道 cur 为空，Morris 遍历结束

在 Morris 遍历中，在第二次遇到当前节点时打印当前节点，则为**中序遍历**。

> 特殊的，当前节点无左孩子时，遍历到当前节点时相当于一次遍历了节点两次

示例如下图：

![morris](https://github.com/A11Might/A11Might.github.io/blob/master/img/post/morristraverse.jpg)

#### 例题 [99. 恢复二叉搜索树](https://leetcode-cn.com/problems/recover-binary-search-tree/)

题目说给定一个二叉搜索树，树中的两个节点被错误地交换，要求在不改变其结构的情况下，恢复这棵树，并且只能使用常数空间复杂度。

对于一个二叉搜索树，我们知道它的中序遍历是有序的，如果树中的两个节点交换的话，也就是中序遍历序列中两个数交换位置，会出现两种情况：

- 交换的是相邻两个数，例如 `1 3 2 4 5 6`，则第一个逆序对，就是被交换的两个数，这里是 3 和 2；
- 交换的是不相邻的数，例如 `1 5 3 4 2 6`，则第一个逆序对的第一个数，和第二个逆序对的第二个数，就是被交换的两个数，这里是 5 和 2；

所以我们可以中序遍历二叉搜索树，找到被交换的两个数，将它们换回来即可。

另外题目要求使用常数的空间复杂度，可以使用 Morris 遍历。

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public void recoverTree(TreeNode root) {
        if (root == null) return;
        TreeNode first = null, second = null; // 记录错误的两个节点
        TreeNode pre = null, cur = root;
        while (cur != null) {
            if (cur.left == null) {
                if (pre != null && pre.val > cur.val) {
                    // 找到一个逆序对
                    if (first == null) {
                        // 是第一个逆序对
                        first = pre;
                        second = cur;
                    } else {
                        // 是第二个逆序对
                        second = cur;
                    }
                }
                pre = cur;
                cur = cur.right;
            } else {
                TreeNode mostRight = cur.left;
                while (mostRight.right != null && mostRight.right != cur) {
                    mostRight = mostRight.right;
                }
                if (mostRight.right == null) {
                    mostRight.right = cur;
                    cur = cur.left;
                } else {
                    mostRight.right = null;
                    if (pre != null && pre.val > cur.val) {
                        // 找到一个逆序对
                        if (first == null) {
                            // 是第一个逆序对
                            first = pre;
                            second = cur;
                        } else {
                            // 是第二个逆序对
                            second = cur;
                        }
                    }
                    pre = cur;
                    cur = cur.right;
                }
            }
        }
        // 交换两个错误节点的值
        int tmp = second.val;
        second.val = first.val;
        first.val = tmp;
    }
}
```

#### 参考

[LeetCode 99. Recover Binary Search Tree - yxc题解](https://www.acwing.com/solution/content/181/)
