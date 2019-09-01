---
layout: post
title: 二叉树的遍历
tags: [codinginterview]
---

二叉树的遍历分为深度优先遍历和广度优先遍历。

## 一、深度优先遍历

### 1、递归方神通

使用递归时系统帮忙压栈，可以来到当前节点三次，分别在这三个位置打印当前节点，即为前序、中序和后序遍历

```java
public void preorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    // one
    preorderTraversal(root.left);
    // two
    preorderTraversal(root.right);
    // three
}
```

#### a、先序遍历

```java
public void preorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    System.out.print(root.val + " ");
    preorderTraversal(root.left);
    preorderTraversal(root.right);
}
```

#### b、中序遍历

```java
public void inorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    inorderTraversal(root.left);
    System.out.print(root.val + " ");
    inorderTraversal(root.right);
}
```

#### c、后序遍历

```java
public void postorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    postorderTraversal(root.left);
    postorderTraversal(root.right);
    System.out.print(root.val + " ");
}
```

### 2、迭代乃人工

使用栈结构，手动模拟系统压栈

#### a、先序遍历

```java
public void preorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    Deque<TreeNode> stack = new ArrayDeque<>();
    stack.push(root);
    while (!stack.isEmpty()) {
        TreeNode cur = stack.pop();
        System.out.print(cur.val + " ");
        // 栈为后进先出，先压入当前节点右孩子，再压入当前节点左孩子
        if (cur.right != null) {
            stack.push(cur.right);
        }
        if (cur.left != null) {
            stack.push(cur.left);
        }
    }
}
```

#### b、中序遍历

```java
public void inorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    Deque<TreeNode> stack = new ArrayDeque<>();
    TreeNode cur = root;
    while (!stack.isEmpty() || cur != null) {
        if (cur != null) {
            stack.push(cur);
            cur = cur.left;
        } else {
            cur = stack.pop();
            System.out.print(cur.val + " ");
            cur = cur.right;
        }
    }
}
```

#### c、后序遍历 use two stack

后序遍历序列的打印顺序为， `左` `右` `中` ，可以通过调整先序遍历中压入孩子顺序，将先序遍历顺序变为 `中` `右` `左` ，后再逆序打印即可

```java
public void postorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    Deque<TreeNode> stack = new ArrayDeque<>();
    Deque<Integer> help = new ArrayDeque<>();
    stack.push(root);
    while (!stack.isEmpty()) {
        TreeNode cur = stack.pop();
        help.push(cur.val);
        // 调整先序遍历当前左右子孩子入栈顺序
        if (cur.left != null) {
            stack.push(cur.left);
        }
        if (cur.right != null) {
            stack.push(cur.right);
        }
    }
    // 逆序打印
    while (!help.isEmpty()) {
        System.out.print(help.pop().val + " ");
    }
}
```

#### d、后序遍历 use on stack

使用指针标记当前节点的左孩子和右孩子是否遍历过，若当前节点左右子孩子中有未遍历过的则将其压入栈中继续dfs，若当前节点左右子孩子都遍历过则打印当前节点，即最后打印中间节点

```java
public void postorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    Deque<TreeNode> stack = new ArrayDeque<>();
    stack.push(root);
    TreeNode cur = null;
    TreeNode visited = root; // visited初始化为非root左右子孩子和null
    while (!stack.isEmpty()) {
        cur = stack.peek();
        // 只使用一个指针标记当前最后遍历的位置
        // 所以判断左孩子是否遍历过，需要同时判断当前左右孩子是否都未遍历过
        // (左孩子遍历过，说明左孩子遍历过；右孩子遍历过，也说明左孩子遍历过)
        if (cur.left != null && visited != cur.left && visited != cur.right) {
            stack.push(cur.left);
        // 判断右孩子是否遍历过，只需判断当前右孩子是否都未遍历过
        } else if (cur.right != null && cur.right != visited) {
            stack.push(cur.right);
        } else {
            System.out.print(stack.pop().val + " ");
            visited = cur;
        }
    }
}
```

### 3、迭代巅峰 Morris遍历

#### Morris序

二叉树中的节点含有空闲的空指针，在遍历节点时将特定空指针指向当前遍历的节点，用于之后返回当前节点位置，将之前迭代遍历中手动压栈所需O(h)(h为二叉树的高度)的额外空间复杂度降为O(1)，具体操作如下

> 将当前节点记为cur
>
> 1、如果cur无左孩子，cur向右移动
>
> 2、如果cur有左孩子，找到cur左孩子最右的节点，记为mostRight
>
>       a、如果mostRight.right指针指向null，让其指向cur，cur向左移动
>
>       b、如果mostRight.right指针指向cur，让其指回null，cur向右移动
>
> 3、如果cur为空，morris遍历结束

![Crepe](/img/post/morris.png){: .center-block :}

```java
public void preorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    TreeNode cur = root;
    TreeNode mostRight = null;
    while (cur != null) {
        mostRight = cur.left;
        // 当前节点有左孩子
        if (mostRight != null) {
            // 找到当前节点左孩子的最右节点
            while (mostRight.right != null && mostRight.right != cur) {
                mostRight = mostRight.right;
            }
            // 第一次来到当前节点
            if (mostRight.right == null) {
                mostRight.right = cur;
                cur = cur.left;
            // 第二次来到当前节点
            } else {
                mostRight.right = null;
                cur = cur.right;
            }
        // 当前节点无左孩子
        } else {
            System.out.print(cur.val + " ");
            cur = cur.right;
        }
    }
}
```

#### a、先序遍历

在Morris遍历中，第一次遇到当前节点时打印当前节点，则为先序遍历

```java
public void preorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    TreeNode cur = root;
    TreeNode mostRight = null;
    while (cur != null) {
        mostRight = cur.left;
        if (mostRight != null) {
            while (mostRight.right != null && mostRight.right != cur) {
                mostRight = mostRight.right;
            }
            if (mostRight.right == null) {
                System.out.print(cur.val + " "); // <-----
                mostRight.right = cur;
                cur = cur.left;
            } else {
                mostRight.right = null;
                cur = cur.right;
            }
        } else {
            System.out.print(cur.val + " "); // <-----
            cur = cur.right;
        }
    }
}
```

#### b、中序遍历

在Morris遍历中，第二次遇到当前节点时打印当前节点，则为中序遍历

{: .box-note}
特殊的，当前节点无左孩子时，遍历到当前节点时相当于一次遍历了节点两次

```java
public void inorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    TreeNode cur = root;
    TreeNode mostRight = null;
    while (cur != null) {
        mostRight = cur.left;
        if (mostRight != null) {
            while (mostRight.right != null && mostRight.right != cur) {
                mostRight = mostRight.right;
            }
            if (mostRight.right == null) {
                mostRight.right = cur;
                cur = cur.left;
            } else {
                mostRight.right = null;
                System.out.print(cur.val + " "); // <-----
                cur = cur.right;
            }
        } else {
            System.out.print(cur.val + " "); // <-----
            cur = cur.right;
        }
    }
    return res;
}
```

#### c、后序遍历

在Morris遍历中， `真第二次` 遇到当前节点时逆序打印当前节点左子树的右边界(当前节点无左孩子时，遍历到当前节点时相当于一次遍历了节点两次，但其没有左子树的右边界，无需打印)，函数退出前打印整棵树的右边界，则为后序遍历

![Crepe](/img/post/postmorris.png){: .center-block :}

```java
public List<Integer> postorderTraversal(TreeNode root) {
    if (root == null) {
        return;
    }
    TreeNode cur = root;
    TreeNode mostRight = null;
    while (cur != null) {
        mostRight = cur.left;
        if (mostRight != null) {
            while (mostRight.right != null && mostRight.right != cur) {
                mostRight = mostRight.right;
            }
            if (mostRight.right == null) {
                mostRight.right = cur;
                cur = cur.left;
            } else {
                mostRight.right = null;
                printEdge(cur.left); // <-----
                cur = cur.right;
            }
        } else {
            cur = cur.right;
        }
    }
    printEdge(root); // <-----
    return ans;
}

// 逆序打印当前节点右边界
private void printEdge(TreeNode node) {
    TreeNode tail = reverseEdge(node);
    TreeNode cur = tail;
    while (cur != null) {
        System.out.print(cur.val + " ");
        cur = cur.right;
    }
    reverseEdge(tail);
}

// 反转当前节点右边界(相当于链表反转)
private TreeNode reverseEdge(TreeNode from) {
    TreeNode pre = null;
    TreeNode cur = from;
    TreeNode succ = null;
    while (cur != null) {
        succ = cur.right;
        cur.right = pre;
        pre = cur;
        cur = succ;
    }
    return pre;
}
```

## 二、广度优先遍历

使用队列来实现层次遍历

```java
public static void bfs(TreeNode root) {
    Deque<TreeNode> queue = new ArrayDeque<>();
    queue.offer(root);
    while (!queue.isEmpty()) {
        TreeNode cur = queue.poll();
        System.out.print(cur.val + " ");
        if (cur.left != null) {
            queue.offer(cur.left);
        }
        if (cur.right != null) {
            queue.offer(cur.right);
        }
    }
}
```

## 三、其他

- [leetcode树的遍历总结](https://github.com/A11Might/leetcode/blob/master/category.md)

- 参考：2018高级算法课 - 左神
