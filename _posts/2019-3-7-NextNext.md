---
layout: post
title: 关于链表的p.next.next操作
tags: [algorithm]
---

在进行链表操作的时候我们会使用到p.next.next，用于链表的反转或者双指针，用途完全不一样。若是使用不当，会出现空指针的报错。

> 单链表反转
>
>> 环形链表
>>
>> 链表的中间结点


### 一、单链表反转（206）

使用递归实现反转链表 

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        if(head == null) return null;
        ListNode p = head;
        if(p.next == null) return p;//递归头，返回反转后的链表表头
        ListNode first = reverseList(p.next);
        p.next.next = p;//<-----this
        p.next = null;
        return first;
    }
}
```

在上述实现中 `p.next.next = p;` 的作用是将当前节点 `p` 后继节点 `p.next` 的后继 `p.next.next` 变为当前节点 `p` 自己，如下图

![_config.yml]({{ site.baseurl }}/images/nextnext1.png)

### 二、双指针

通过使用具有不同速度的快、慢两个指针遍历链表

#### a、环形链表（141）

给定一个链表，判断链表中是否有环

```java
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        if(head == null || head.next == null) return false;
        ListNode slow = head;
        ListNode fast = head.next;
        while(slow != fast){
            if(fast == null || fast.next == null) return false;
            slow = slow.next;
            fast = fast.next.next;//<-----this
        }
        return true;
    }
}
```

#### b、链表的中间结点（876）

给定一个带有头结点 head 的非空单链表，返回链表的中间结点；如果有两个中间结点，则返回第二个中间结点

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode middleNode(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;
        while(fast != null && fast.next != null){
            slow = slow.next;
            fast = fast.next.next;//<-----this
        }
        return slow;
    }
}
```

#### c、边界条件

1、在使用快慢指针的时候需要注意边界条件，若使用不当则可能出现空指针报错。判断是否到达边界的条件是 `fast != null` 和 `fast.next != null` ,如下图

![_config.yml]({{ site.baseurl }}/images/nextnext2.png)

2、 `if(fast == null || fast.next == null)` 和 `while(fast != null && fast.next != null)` ？

前者判断当前节点是否到达链尾，若到达则停止操作，fast == null和fast.next == null都是判断是当前节点否到达链尾的标志，满足其一即可；

而后者为了防止循环后空指针，fast != null和fast.next != null必须同时满足才能继续执行循环
