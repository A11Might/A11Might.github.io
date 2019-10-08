---
layout: post
title: 图的广度优先遍历
tags: [leetcode]
---

使用广度优先遍历求无权图的最短路径

### 1. 完全平方数 leetcode[[279]](https://leetcode-cn.com/problems/perfect-squares/) [题解](https://github.com/A11Might/leetcode/blob/master/codes/lc279.java)

> 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

对问题建模，将整个问题转化为一个图论问题

从n到0，每个数字表示一个节点，如果两个数字x到y相差一个完全平方数则连接一条边，就可以得到一个无权图，原问题转化为，求这个无权图中从n到0的最短路径

![Crepe](/img/post/modeling.png){: .center-block :}

```java
public int numSquares(int n) {
    if (n == 0) {
        return 0;
    }
    boolean[] visited = new boolean[n + 1]; // 标记当前节点是否访问过
    Deque<Pair<Integer, Integer>> queue = new ArrayDeque<>();
    // 从n点出发寻找最短路径，当前步数为0
    queue.offer(new Pair<>(n, 0));
    visited[n] = true;
    while (!queue.isEmpty()) {
        Pair cur = queue.poll();
        int num = (int) cur.getKey();
        int step = (int) cur.getValue();

        // 向所有能走的节点前进
        for (int i = 1; num - i * i >= 0; i++) {
            // 若当前节点之前访问过，则跳过(若再走，步数一定大于之前访问的步数)
            if (!visited[num - i * i]) {
                // 走到0点，返回步数
                if (num - i * i == 0) return step + 1;
                queue.offer(new Pair<>(num - i * i, step + 1));
                visited[num - i * i] = true;
            }
        }
    }

    // 无解(其实肯定有解，1)
    throw new IllegalStateException("no solution");
}
```

### 2. 单词接龙 leeetcode[[127]](https://leetcode-cn.com/problems/word-ladder/) [题解](https://github.com/A11Might/leetcode/blob/master/codes/lc127.java)

> 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。
>
> 转换需遵循如下规则：   
> 每次转换只能改变一个字母。
> 转换过程中的中间单词必须是字典中的单词。
>
> 说明:   
> 如果不存在这样的转换序列，返回 0。
> 所有单词具有相同的长度。
> 所有单词只由小写字母组成。
> 字典中不存在重复的单词。
> 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

对问题建模，将整个问题转化为一个图论问题

每个单词表示一个节点，差距只有一个字母的两个单词之间连一条边，就可以得到一个无权图，原问题转化为，从这个无权图中找到从起点beginWord到终点endWord的最短路径

算法中最重要的步骤是找出相邻的节点，也就是只差一个字母的两个单词，具体如下

```java
private ArrayList<String> getNeighbors(HashSet<String> dict, String word) {
    ArrayList<String> ans = new ArrayList<String>();
    int n = word.length();
    char[] chrs = word.toCharArray();
    // 将当前单词的每一位替换为其它字母
    for (int i = 0; i < n; i++) {
        for (char chr = 'a'; chr <= 'z'; chr++) {
            if (chrs[i] == chr) continue;
            char oldChr = chrs[i];
            chrs[i] = chr;
            // 判断修改后的单词是否存在wordList中
            // 以此找到所有与当前单词只相差一个字母的所有单词
            if (dict.contains(String.valueOf(chrs))) {
                ans.add(String.valueOf(chrs));
            }
            chrs[i] = oldChr;
        }
    }

    return ans;
}
```

使用广度优先搜素找到beginWord到endWord之间的最短路径

```java
public int ladderLength(String beginWord, String endWord, List<String> wordList) {
    if (!wordList.contains(endWord)) {
        return 0;
    }
    // 将list转化为hashset，加快查找速度
    // 使用hashmap记录beginWord到遍历过的节点的路径
    HashSet<String> dict = new HashSet<String>(wordList); 
    HashMap<String, Integer> distance = new HashMap<String, Integer>(); 
    // 使用队列进行广度优先遍历
    // 从beginWord节点出发寻找最短路径，当前路径长度为1
    Deque<String> queue = new ArrayDeque<String>();
    queue.offer(beginWord);
    distance.put(beginWord, 1);
    while (!queue.isEmpty()) {
        int curLevelSize = queue.size();
        // 遍历当前层的每一个节点
        for (int i = 0; i < curLevelSize; i++) {
            String cur = queue.poll();
            int curDistance = distance.get(cur);
            // 获取当前节点所连接下一层节点
            ArrayList<String> neighbors = getNeighbors(dict, cur);
            // 遍历当前节点所连接下一层的所有节点
            for (String neighbor : neighbors) {
                if (!distance.containsKey(neighbor)) { // distance中为遍历过的节点(有visited功能)
                    if (neighbor.equals(endWord)) { // 找到最短路径
                        return curDistance + 1;
                    }
                    queue.offer(neighbor);
                    distance.put(neighbor, curDistance + 1);
                }
            }
        }
    }

    return 0;
}
```

### 3. 单词接龙 II leeetcode[[126]](https://leetcode-cn.com/problems/word-ladder-ii/) [题解](https://github.com/A11Might/leetcode/blob/master/codes/lc126.java)

> 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。
>
> 转换需遵循如下规则：    
> 每次转换只能改变一个字母。
> 转换过程中的中间单词必须是字典中的单词。
>
> 说明:   
> 如果不存在这样的转换序列，返回一个空列表。
> 所有单词具有相同的长度。
> 所有单词只由小写字母组成。
> 字典中不存在重复的单词。
> 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

如上题，对问题建模

使用bfs找到beginWord到endWord最短路径，并使用hashmap记录所有经过的节点到beginWord的距离(用于dfs正确的找出所有最短路径)

```java
private void bfs(String beginWord, String endWord, HashSet<String> dict, HashMap<String, ArrayList<String>> allNeighbors, HashMap<String, Integer> distance) {
    Deque<String> queue = new ArrayDeque<String>();
    queue.offer(beginWord);
    distance.put(beginWord, 0);
    while (!queue.isEmpty()) {
        int curLevelSize = queue.size(); // 当前层所含节点数
        boolean findEndWord = false; // 标记是否找到最短路径
        // 遍历当前层的每一个节点
        for (int i = 0; i < curLevelSize; i++) {
            String cur = queue.poll();
            int curDistance = distance.get(cur);
            // 获取当前节点所连接下一层节点
            // 并使用hashmap记录下来，供dfs使用(dfs无需再次使用getNeighbors获取)
            ArrayList<String> neighbors = getNeighbors(dict, cur);
            allNeighbors.put(cur, neighbors);
            // 遍历当前节点所连接下一层的所有节点
            for (String neighbor : neighbors) {
                if (!distance.containsKey(neighbor)) { // distance中为遍历过的节点(有visited功能)
                    if (neighbor.equals(endWord)) { // 找到最短路径
                        findEndWord = true;
                    }
                    queue.offer(neighbor);
                    distance.put(neighbor, curDistance + 1);
                }
            }
        }

        // 找到最短路径并遍历完最短路径所在的最后一层即可返回(再往下遍历不可能是最短路径)
        if (findEndWord) {
            break;
        }
    }
}
```

通过dfs找到所有长度等于最短路径的路径

```java
private void dfs(String cur, String endWord, HashMap<String, ArrayList<String>> allNeighbors, HashMap<String, Integer> distance, ArrayList<String> solution, ArrayList<List<String>> res) {
    solution.add(cur);
    if (cur.equals(endWord)) {
        res.add(new ArrayList<String>(solution));
    } else {
        // 由于bfs找到最短路径时，遍历完最短路径那一层就停了，所以那一层没有neighbors，使用getOrDefault防止空指针
        for (String neighbor : allNeighbors.getOrDefault(cur, new ArrayList<String>())) {
            // 防止neighbor是当前节点的上一层节点(确保neighbor是当前节点的下一层节点，否则会栈溢出)
            if (distance.get(cur) + 1 == distance.get(neighbor)) {
                dfs(neighbor, endWord, allNeighbors, distance, solution, res);
            }
        }
    }
    solution.remove(solution.size() - 1);
}
```

{: .box-note}
**Note:** bfs找到所有可能为最短路径的节点，再使用dfs将所有最短路径的节点串联起来

### 4. 参考

- 玩转算法面试 - 波波老师

- leetcode官方题解 - [单词接龙](https://leetcode-cn.com/problems/word-ladder/solution/dan-ci-jie-long-by-leetcode/)

- leetcode高赞解 - [Cheng_Zhang
](https://leetcode.com/problems/word-ladder-ii/discuss/40475/My-concise-JAVA-solution-based-on-BFS-and-DFS)