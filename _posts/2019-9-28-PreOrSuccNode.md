---
layout: post
title: 前驱节点和后继节点
tags: [codinginterview]
---

在二叉树的中序遍历的序列中，当前节点的前一个节点叫作当前节点的前驱节点，当前节点的后一个节点叫做当前节点的后继节点。若给定一个二叉树和其中一个节点node，如何求出node的前驱节点或者后继节点。

#### 1. 剑指offer[8] 二叉树的下一个节点 [题解](https://github.com/A11Might/codingInterview/blob/master/code/offer08.java)


> 给定一个二叉树和其中的一个节点node，请找出中序遍历顺序的下一个结点并且返回。
> 
> 注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针next。

- 当节点node有右子树时，其后继节点为其右子树中最左的节点

- 当节点node无右子树时，判断其是否为其父节点的左子节点

    - 若是，其后继节点为其父节点

    - 若不是，继续向上寻找其祖父节点，判断其父亲节点是否是其祖父节点的左子节点
    
        - 若是，其后继节点为其祖父节点

        - 若不是，继续下上寻找其曾祖父节点...
        
        - 直至当前节点为空，节点node无后继节点

![Crepe](/img/post/succnode.jpg){: .center-block :}

{: .box-note}
**Note:** 将二叉树想象成只有三个节点(根节点、左子节点和右子节点)，左子节点的后继节点就是根节点，即当前node节点无右子树，其后继为某节点，该节点左子树的最后一个节点为node；根节点的后继节点就是右子节点，即当前node节点有右子树，其后继为右子树中最左的节点；右子节点无后继节点

```java
public TreeLinkNode GetNext(TreeLinkNode pNode) {
    if (pNode == null) {
        return null;
    }
    // 当节点node有右子树
    if (pNode.right != null) {
        pNode = pNode.right;
        while (pNode.left != null) {
            pNode = pNode.left;
        }
        return pNode;
    // 当节点node无右子树
    } else {
        TreeLiknNode parent = pNode.parent;
        while (parent != null && parent.left != pNode) {
            pNode = parent;
            parent = pNode.parent;
        }
        return parent;
    }
}
```

#### 2. 二叉树的上一个节点 [题解](https://github.com/A11Might/SomePracticeCode/blob/master/learningCode/PrecurorNode.java)

同理

![Crepe](/img/post/prenode.jpg){: .center-block :}

#### 3. 参考

- 左神 - 2018算法初级课
