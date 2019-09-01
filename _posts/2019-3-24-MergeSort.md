---
layout: post
title: 小和问题
tags: [algorithm]
---

在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。给你一个数组，求它的小和。

>例[1,3,4,2,5]
>
>1左边比1小的数， 没有；
>
>3左边比3小的数， 1；
>
>4左边比4小的数， 1、 3；
>
>2左边比2小的数， 1；
>
>5左边比5小的数， 1、 3、 4、 2；
>
>所以小和为1+1+3+1+1+3+4+2=16

#### 初解

要解决小和问题，用笨方法很简单，两边遍历比较每个元素与其前面序列中元素的大小关系，维护一个变量res记录小和，如下：

```java
public static int smallSumSure1(int[] arr) {
        if (arr == null || arr.length < 2) {
			return 0;
		}
		int res = 0;
		for (int i = 1; i < arr.length; i++) {
			for (int j = 0; j < i; j++) {
				res += arr[j] < arr[i] ? arr[j] : 0;
			}
		}
		return res;
    }
```

但这样比较费时，每个数都需要依次比较，时间复杂度高达 O(n^2) 。有没有更好的方法呢？

#### 终解

先来看看归并排序，原理是将待排序列分为两个子序列，分别排序后，将两个有序子序列合并。在两个子序列S、M合并过程中，当S中元素A小于M中元素B时，因为M序列已经有序，所以不用比较，A一定是S、M两个子序列剩余元素中最小的元素，这省去了A与B后其他元素比较的操作。

![_config.yml]({{ site.baseurl }}/images/xche.jpg)

那小和问题怎么省去多余的比较操作？将原本题目反过来想：当前元素右侧序列每有一个比其大的元素，就有一个左边比当前数小的数即当前元素，维护一个变量res将其累加，如下：

```java
public static int smallSum2(int[] arr) {
		if (arr == null || arr.length < 2) {
			return 0;
		}
		return mergeSort(arr, 0, arr.length - 1);
	}

	public static int mergeSort(int[] arr, int l, int r) {
		if (l == r) {
			return 0;
		}
		int mid = l + ((r - l) >> 1);
		return mergeSort(arr, l, mid) + mergeSort(arr, mid + 1, r) + merge(arr, l, mid, r);
	}

	public static int merge(int[] arr, int l, int m, int r) {
		int[] help = new int[r - l + 1];
		int i = 0;
		int p1 = l;
		int p2 = m + 1;
		int res = 0;
		while (p1 <= m && p2 <= r) {
			res += arr[p1] < arr[p2] ? (r - p2 + 1) * arr[p1] : 0;  //<--- 这里
			help[i++] = arr[p1] < arr[p2] ? arr[p1++] : arr[p2++];
		}
		while (p1 <= m) {
			help[i++] = arr[p1++];
		}
		while (p2 <= r) {
			help[i++] = arr[p2++];
		}
		for (i = 0; i < help.length; i++) {
			arr[l + i] = help[i];
		}
		return res;
	}
```

#### 对数器验证

>对数器原理
>
>0， 有一个你想要测的方法a，
>
>1， 实现一个绝对正确但是复杂度不好的方法b，
>
>2， 实现一个随机样本产生器
>
>3， 实现比对的方法
>
>4， 把方法a和方法b比对很多次来验证方法a是否正确
>
>5， 如果有一个样本使得比对出错， 打印样本分析是哪个方法出错 
>
>6， 当样本数量很多时比对测试依然正确， 可以确定方法a已经正确

使用对数器验证，优化后的方法是否正确。

```java
	public static int comparator(int[] arr) {
		if (arr == null || arr.length < 2) {
			return 0;
		}
		int res = 0;
		for (int i = 1; i < arr.length; i++) {
			for (int j = 0; j < i; j++) {
				res += arr[j] < arr[i] ? arr[j] : 0;
			}
		}
		return res;
	}

	public static int[] generateRandomArray(int maxSize, int maxValue) {
		int[] arr = new int[(int) ((maxSize + 1) * Math.random())];
		for (int i = 0; i < arr.length; i++) {
			arr[i] = (int) ((maxValue + 1) * Math.random()) - (int) (maxValue * Math.random());
		}
		return arr;
	}

	public static int[] copyArray(int[] arr) {
		if (arr == null) {
			return null;
		}
		int[] res = new int[arr.length];
		for (int i = 0; i < arr.length; i++) {
			res[i] = arr[i];
		}
		return res;
	}

	public static boolean isEqual(int[] arr1, int[] arr2) {
		if ((arr1 == null && arr2 != null) || (arr1 != null && arr2 == null)) {
			return false;
		}
		if (arr1 == null && arr2 == null) {
			return true;
		}
		if (arr1.length != arr2.length) {
			return false;
		}
		for (int i = 0; i < arr1.length; i++) {
			if (arr1[i] != arr2[i]) {
				return false;
			}
		}
		return true;
	}

	public static void printArray(int[] arr) {
		if (arr == null) {
			return;
		}
		for (int i = 0; i < arr.length; i++) {
			System.out.print(arr[i] + " ");
		}
		System.out.println();
	}

	public static void main(String[] args) {
		int testTime = 500000;
		int maxSize = 100;
		int maxValue = 100;
		boolean succeed = true;
		for (int i = 0; i < testTime; i++) {
			int[] arr1 = generateRandomArray(maxSize, maxValue);
			int[] arr2 = copyArray(arr1);
			if (smallSum(arr1) != comparator(arr2)) {
				succeed = false;
				printArray(arr1);
				printArray(arr2);
				break;
			}
		}
		System.out.println(succeed ? "Nice!" : "Fucking fucked!");
	}
```

#### 时间复杂度

>master公式的使用
>
>T(N) = a*T(N/b) + O(N^d)
>
>1) log(b,a) > d -> 复杂度为O(N^log(b,a))
>
>2) log(b,a) = d -> 复杂度为O(N^d * logN)
>
>3) log(b,a) < d -> 复杂度为O(N^d)
>
>[补充阅读]<www.gocalf.com/blog/algorithm-complexity-and-mastertheorem.html>

使用master公式估算终解时间复杂度，O(n*logn)

#### 最后

- `res += arr[p1] < arr[p2] ? (r - p2 + 1) * arr[p1] : 0;` 与 `help[i++] = arr[p1] < arr[p2] ? arr[p1++] : arr[p2++];` 的 `<` 能否换成 `<=` ?

不能，归并排序中 `<=` 作用为保持排序的稳定性；小和问题中 `<` 作用为计算有多少个大于a的数。

![_config.yml]({{ site.baseurl }}/images/xche2.jpg)

- [源码](https://github.com/A11Might/SomePracticeCode/blob/master/learningCode/SmallSum.java)
