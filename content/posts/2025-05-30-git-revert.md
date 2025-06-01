---
date: '2025-05-30T22:20:02+08:00'
draft: false
title: 'Git 实战：回滚合并提交'
tags: ['Git']
---

之前需求开发分支（dev）已经合到（merge）上线分支（release）后，产品又说不上了要回滚代码，shit。我能想到的是用 `git revert` 来回滚**公共分支**（release）的提交，但对 `git revert` 的使用还停留在 `git revert A` 上，抓瞎。下面是我实践和验证后的操作记录。

### 场景还原

提交历史如下所示，dev merge 到 release 上产生了 mrege 提交 M，后面可能还有其他开发者的提交或合并记录（这里的 4 和 5）。

```
                                       HEAD
                                        │
                                        ▼
                                      release
                                        │
                                        ▼
1 ────► 2 ────► 3 ────► M ────► 4 ────► 5
                        ▲ 
                        │
        A ────► B ──────┘
                ▲
                │
               dev
```

### 回滚操作

使用 -m 参数回滚 merge 提交，命令：

```bash
$ git revert -m 1 M
```

生成 merge 提交 M 的 revert W（W 像不像 M 倒过来，：P），此时 release 分支就完成了回滚，不包含不上需求的提交（dev 的 A 和 B），可以正常上线了。

```
                                               HEAD
                                                │
                                                ▼
                                             release
                                                │
                                                ▼
1 ────► 2 ────► 3 ────► M ────► 4 ────► 5 ────► W
                        ▲ 
                        │
        A ────► B ──────┘
                ▲
                │
               dev
```

其中的参数含义：

- -m 1：指定主线的父节点编号（从 1 开始计数），让 revert 操作**基于该父节点对应的分支状态进行变更反转**。简单来说：
  - -m 1：revert 到目标分支（这里的 release）的状态；

  - -m 2：revert 到来源分支（这里的 dev）的状态。

- M：合并提交的哈希值（这里的 `M`），也就是要撤销的合并操作。

### 二次上线操作

如果之后需求又要上线，回滚之前的回滚操作即可，命令：`git revert W`，其中 Y 是 W 的 revert。

```
                                                               HEAD
                                                                │
                                                                ▼
                                                             release
                                                                │
                                                                ▼
1 ────► 2 ────► 3 ────► M ────► 4 ────► 5 ────► W ────► 6 ────► Y
                        ▲ 
                        │
        A ────► B ──────┘
                ▲
                │
               dev
```

另外如果 dev 还有新修改（需求变更或 Bug 修复），直接在 dev 上进行修改后，然后再合到 release 上。

```
                                                                       HEAD
                                                                        │
                                                                        ▼
                                                                     release
                                                                        │
                                                                        ▼
1 ────► 2 ────► 3 ────► M ────► 4 ────► 5 ────► W ────► 6 ────► Y ────► *
                        ▲                                               ▲ 
                        │                                               │
        A ────► B ──────┘ ────────────────────────────► C ────► D ──────┘ 
                                                                ▲
                                                                │
                                                               dev
```

### References

- [git-revert - Revert some existing commits](https://git-scm.com/docs/git-revert)
- [revert-a-faulty-merge How-To](https://github.com/git/git/blob/7014b55638da979331baf8dc31c4e1d697cf2d67/Documentation/howto/revert-a-faulty-merge.adoc)

