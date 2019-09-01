---
layout: post
title: Markdown教程
tags: other
---

> Markdown 是一种轻量级标记语言，创始人为 John Gruber。它允许人们「使用易读易写的纯文本格式编写文档，然后转换成有效的 XHTML（或者 HTML）文档」。——维基百科

本文档的目的不在于面面俱到地介绍 Markdown，只是作为我对其理解的笔记整理，希望能同时帮助一些对 Markdown 感兴趣的人快速上手，或是作为一个工具，供对其已经有所了解的人在需要时参考。

## 优点

1. 专注于文字内容；

2. 纯文本，易读易写，可以方便地纳入版本控制；

3. 语法简单，没有什么学习成本，能轻松在码字的同时做出美观大方的排版。

## 语法

### 标题

**Markdown：**

```
# atx-style 一级标题

## 二级标题

###### 六级标题

Setext-style 一级标题
===

二级标题
---
```

**预览效果：**

> # atx-style 一级标题
> 
> ## 二级标题
> 
> ###### 六级标题
>
> Setext-style 一级标题
> ===
>
> 二级标题
> ---

**对应 HTML：**

```html
<h1>atx-style 一级标题</h1>

<h2>二级标题</h2>

<h6>六级标题</h6>

<h1>Setext-style 一级标题</h1>

<h2>二级标题</h2>
```

### 段落

中间没有空行的连续不断的几行文字被视为一个段落。

**Markdown：**

```
白日依山尽，

黄河入海流。
（句号后面没空格）

欲穷千里目，

更上一层楼。  
（句号后面有俩空格）
```

**预览效果：**

白日依山尽，

黄河入海流。
（句号后面没空格）

欲穷千里目，

更上一层楼。  
（句号后面有俩空格）

**对应 HTML：**

```html
<p>白日依山尽，</p>

<p>黄河入海流。
（句号后面没有空格）</p>

<p>欲穷千里目，</p>

<p>
  更上一层楼。
  <br>
  （句号后面有俩空格）
</p>
```

### 行内格式

对段落或者部分文本的强调效果。

**Markdown：**

```
后面俩字**加黑**

后面俩字*斜体*
```

**预览效果：**

后面俩字**加黑**

后面俩字*斜体*

**对应 HTML：**

```html
<p>
  后面俩字
  <strong>加黑</strong>
</p>
<p>
  后面俩字
  <em>斜体</em>
</p>
```

### 引用块

**Markdown：**

```
> 引用块段落一。
>
> 引用块段落二。
>> 内嵌引用块段落一。
>
> ### 引用块内的标题
```

**预览效果：**

> 引用块段落一。
>
> 引用块段落二。
>
> > 内嵌引用块段落一。
>
> ### 引用块内的标题

**对应 HTML：**

```html
<blockquote>
  <p>引用块段落一。</p>
  <p>引用块段落二。</p>
  <blockquote>
    <p>内嵌引用块段落一。</p>
  </blockquote>
  <h3 id="引用块内的标题">引用块内的标题</h3>
</blockquote>
```

### 超链接

Markdown 支持行内式链接和引用式链接。

**Markdown：**

```
行内式 [博客](https://mazhuang.org "我的个人博客") 链接，带 title。

行内式 [GitHub](https://github.com/mzlogin) 链接。

引用式 [博客][1] 链接。

引用式 [GitHub][2] 链接，带 title。

[1]: https://mazhuang.org
[2]: https://github.com/mzlogin "我的 GitHub 主页"
```

**预览效果：**

行内式 [博客](https://mazhuang.org "我的个人博客") 链接，带 title。

行内式 [GitHub](https://github.com/mzlogin) 链接。

引用式 [博客][1] 链接。

引用式 [GitHub][2] 链接，带 title。

[1]: https://mazhuang.org
[2]: https://github.com/mzlogin "我的 GitHub 主页"

**对应 HTML：**

```html
<p>行内式 <a href="https://mazhuang.org" title="我的个人博客">博客</a> 链接，带 title。</p>

<p>行内式 <a href="https://github.com/mzlogin">GitHub</a> 链接。</p>

<p>引用式 <a href="https://mazhuang.org">博客</a> 链接。</p>

<p>引用式 <a href="https://github.com/mzlogin" title="我的 GitHub 主页">GitHub</a> 链接，带 title。</p>
```

### 图片

在超链接的写法前加一个 `!`，就是引用图片的方法。

**Markdown：**

```
![Alt text](https://mazhuang.org/favicon.ico "favicon")
```

**预览效果：**

![Alt text](https://mazhuang.org/favicon.ico "favicon")

**对应 HTML：**

```html
<img src="https://mazhuang.org/favicon.ico" alt="Alt text" title="favicon">
```

### 列表

包括有序列表和无序列表。

**Markdown：**

```
- 苹果
- 葡萄
- 榴莲

1. 苹果
2. 葡萄
3. 榴莲
```

**预览效果：**

- 苹果
- 葡萄
- 榴莲

1. 苹果
2. 葡萄
3. 榴莲

**对应 HTML：**

```html
<ul>
  <li>苹果</li>
  <li>葡萄</li>
  <li>榴莲</li>
</ul>
<ol>
  <li>苹果</li>
  <li>葡萄</li>
  <li>榴莲</li>
</ol>
```

其中无序列表的标记可以使用 `+`、`-` 或 `*`，有序列表前的数字可以是乱序的。

### 代码块

支持行内代码和代码块。

**Markdown：**

    Android 里使用 `TextUtils` 类的 `isEmpty` 方法来判断字符串是否为空。

    ```java
    if (TextUtils.isEmpty(text)) {
        return null;
    }
    ```

**预览效果：**

Android 里使用 `TextUtils` 类的 `isEmpty` 方法来判断字符串是否为空。

```java
if (TextUtils.isEmpty(text)) {
    return null;
}
```

**对应 HTML：**

```html
<p>Android 里使用 <code>TextUtils</code> 类的 <code>isEmpty</code> 方法来判断字符串是否为空。</p>

<div class="highlight highlight-source-java"><pre><span class="pl-k">if</span> (<span class="pl-smi">TextUtils</span><span class="pl-k">.</span>isEmpty(text)) {
    <span class="pl-k">return</span> <span class="pl-c1">null</span>;
}</pre></div>
```

上例中的语言标记 `java` 可选填，可用于在编辑器和渲染后的效果里添加语法高亮。

块式代码也可以对整个代码段缩进四个空格，或一个 Tab 来实现。

### 水平分割线

使用一个单独行里的三个或以上 `*`、`-` 来生产一条水平分割线，它们之间可以有空格。

**Markdown：**

```
***

-----

- - -
```

**预览效果：**

***

-----

- - -

**对应 HTML：**

```
<hr />

<hr />

<hr />
```

### 嵌入 HTML

Markdown 标记语言的目的不是替代 HTML，也不是发明一种更便捷的插入 HTML 标签的方式。它对应的只是 HTML 标签的一个很小的子集。

对于那些没有办法用 Markdown 语法来对应的 HTML 标签，直接使用 HTML 来写就好了。

## 扩展语法

本节的内容是介绍一些受到广泛支持的 Markdown 扩展语法。

### 表格

**Markdown：**

    | 编号  | 姓名（左） | 年龄（右） | 性别（中） |
    | ----- | :--------  | ---------: | :------:   |
    | 0     | 张三       | 28         | 男         |
    | 1     | 李四       | 29         | 男         |

**预览效果：**

| 编号  | 姓名（左） | 年龄（右） | 性别（中） |
| ----- | :--------  | ---------: | :------:   |
| 0     | 张三       | 28         | 男         |
| 1     | 李四       | 29         | 男         |

**对应 HTML：**

```html
<table>
  <thead>
    <tr>
      <th>编号</th>
      <th align="left">姓名（左）</th>
      <th align="right">年龄（右）</th>
      <th align="center">性别（中）</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>0</td>
      <td align="left">张三</td>
      <td align="right">28</td>
      <td align="center">男</td>
    </tr>
    <tr>
      <td>1</td>
      <td align="left">李四</td>
      <td align="right">29</td>
      <td align="center">男</td>
    </tr>
  </tbody>
</table>
```

### 任务列表

在 GitHub / GitLab 里有较好的支持。

**Markdown：**

```
- [x] 洗碗
- [ ] 清洗油烟机
- [ ] 拖地
```

**预览效果：**

- [x] 洗碗
- [ ] 清洗油烟机
- [ ] 拖地

**对应 HTML：**

```html
<ul class="contains-task-list">
  <li class="task-list-item"><input type="checkbox" id="" disabled="" class="task-list-item-checkbox" checked=""> 洗碗</li>
  <li class="task-list-item"><input type="checkbox" id="" disabled="" class="task-list-item-checkbox"> 清洗油烟机</li>
  <li class="task-list-item"><input type="checkbox" id="" disabled="" class="task-list-item-checkbox"> 拖地</li>
</ul>
```

如果是在 GitHub / GitLab 的 Issue 里，会附赠任务完成比例提示效果：

![task list 1](https://raw.githubusercontent.com/mzlogin/markdown-intro/master/assets/task-list-1.png)

还可以直接在网页上拖动调整顺序，勾选和取消勾选。

![task list 2](https://raw.githubusercontent.com/mzlogin/markdown-intro/master/assets/task-list-2.png)

### 删除线

**Markdown：**

```
后面三个字打上~~删除线~~。
```

**预览效果：**

后面三个字打上~~删除线~~。

**对应 HTML：**

```html
<p>后面三个字打上<del>删除线</del>。</p>
```

### 自动链接

自动链接扩展，即：当识别到 URL，或用 `<`、`>` 包括的 URL 时，会自动为其生成 `a` 标签。

**Markdown：**

```
https://github.com

<example@gmail.com>
```

**预览效果：**

https://github.com

<example@gmail.com>

**对应 HTML：**

```html
<p><a href="https://github.com">https://github.com</a></p>

<p><a href="mailto:example@gmail.com">example@gmail.com</a></p>
```

### emoji

以 GitHub Pages 为例。

**Markdown：**

```
:camel: :blush: :smile:
```

**预览效果：**

:camel: :blush: :smile:

**对应 HTML：**

```html
<p>
  <img class="emoji" title=":camel:" alt=":camel:" src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f42b.png" height="20" width="20">
  <img class="emoji" title=":blush:" alt=":blush:" src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f60a.png" height="20" width="20">
  <img class="emoji" title=":smile:" alt=":smile:" src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f604.png" height="20" width="20">
</p>
```

## 奇技淫巧

脑洞清奇的工程师们还发掘了很多使用 Markdown 的方法，大部分都是引入第三方 JavaScript 插件来实现。对这部分我只做简述，对其中的部分功能比如作图等，还是推荐用专门的可视化工具去做。

### 画流程图和时序图

有部分网站和编辑器实现了对 Markdown 里流程图和时序图的支持，比如我们使用的项目管理工具 TAPD 的在线编辑器，还有 VSCode + 插件 Markdown Preview Enhanced 等。

以我们使用的项目管理工具 TAPD 的在线编辑器为例：

![流程图](https://raw.githubusercontent.com/mzlogin/markdown-intro/master/assets/tapd-markdown-flowchart.png)

![时序图](https://raw.githubusercontent.com/mzlogin/markdown-intro/master/assets/tapd-markdown-seq.png)

### 插入数学公式

仍然以 TAPD 为例：

![数学公式](https://raw.githubusercontent.com/mzlogin/markdown-intro/master/assets/tapd-markdown-math.png)

应该是利用 JavaScript 支持了 LaTeX 公式语法。

### 用 Markdown 做 PPT

有专门的工具 [Marp](https://github.com/yhatt/marp)，另外使用 VSCode + 插件 Markdown Preview Enhanced 也可以实现。

### 用 Markdown 写微信公众号

可以将公众号素材用 Markdown 编辑好后，贴到在线排版工具以后，复制到公众号编辑器里即可。有多种页面主题和代码主题可选择。

### 在表格单元格里换行

借助于 HTML 里的 `<br />` 实现。

示例代码：

```
| Header1 | Header2                          |
|---------|----------------------------------|
| item 1  | 1. one<br />2. two<br />3. three |
```

示例效果：

| Header1 | Header2                          |
|---------|----------------------------------|
| item 1  | 1. one<br />2. two<br />3. three |

### 图文混排

使用 `<img>` 标签来贴图，然后指定 `align` 属性。

示例代码：

```
<img align="right" src="https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png"/>

这是一个示例图片。

图片显示在 N 段文字的右边。

N 与图片高度有关。

刷屏行。

刷屏行。

到这里应该不会受影响了，本行应该延伸到了图片的正下方，所以我要足够长才能确保不同的屏幕下都看到效果。
```
示例效果：

<img align="right" src="https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png"/>

这是一个示例图片。

图片显示在 N 段文字的右边。

N 与图片高度有关。

刷屏行。

刷屏行。

到这里应该不会受影响了，本行应该延伸到了图片的正下方，所以我要足够长才能确保不同的屏幕下都看到效果。

### 控制图片大小和位置

标准的 Markdown 图片标记 `![]()` 无法指定图片的大小和位置，只能依赖默认的图片大小，默认居左。

而有时候源图太大想要缩小一点，或者想将图片居中，就仍需要借助 HTML 的标签来实现了。图片居中可以使用 `<div>` 标签加 `align` 属性来控制，图片宽高则用 `width` 和 `height` 来控制。

示例代码：

```
**图片默认显示效果：**

![](https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png)

**加以控制后的效果：**

<div align="center"><img width="65" height="75" src="https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png"/></div>
```

示例效果：

**图片默认显示效果：**

![](https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png)

**加以控制后的效果：**

<div align="center"><img width="65" height="75" src="https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/demo.png"/></div>

### 格式化表格

表格在渲染之后很整洁好看，但是在文件源码里却可能是这样的：

```
|Header1|Header2|
|---|---|
|a|a|
|ab|ab|
|abc|abc|
```

不知道你能不能忍，反正我是不能忍。

好在广大网友们的智慧是无穷的，在各种编辑器里为 Markdown 提供了表格格式化功能，比如我使用 Vim 编辑器，就有 [vim-table-mode](https://github.com/dhruvasagar/vim-table-mode) 插件，它能帮我自动将表格格式化成这样：

```
| Header1 | Header2 |
|---------|---------|
| a       | a       |
| ab      | ab      |
| abc     | abc     |
```

是不是看着舒服多了？

如果你不使用 Vim，也没有关系，比如 Atom 编辑器的 [markdown-table-formatter](https://atom.io/packages/markdown-table-formatter) 插件，Sublime Text 3 的 [MarkdownTableFormatter](https://github.com/bitwiser73/MarkdownTableFormatter) 等等，都提供了类似的解决方案。

### 使用 Emoji

这个是 GitHub 对标准 Markdown 标记之外的扩展了，用得好能让文字生动一些。

示例代码：

```
我和我的小伙伴们都笑了。:smile:
```

示例效果：

我和我的小伙伴们都笑了。:smile:

更多可用 Emoji 代码参见 <https://www.webpagefx.com/tools/emoji-cheat-sheet/>。

### 行首缩进

直接在 Markdown 里用空格和 Tab 键缩进在渲染后会被忽略掉，需要借助 HTML 转义字符在行首添加空格来实现，`&ensp;` 代表半角空格，`&emsp;` 代表全角空格。

示例代码：

```
&emsp;&emsp;春天来了，又到了万物复苏的季节。
```

示例效果：

&emsp;&emsp;春天来了，又到了万物复苏的季节。

### 展示数学公式

如果是在 GitHub Pages，可以参考 <http://wanguolin.github.io/mathmatics_rending/> 使用 MathJax 来优雅地展示数学公式（非图片）。

如果是在 GitHub 项目的 README 等地方，目前我能找到的方案只能是贴图了，以下是一种比较方便的贴图方案：

1. 在 <https://www.codecogs.com/latex/eqneditor.php> 网页上部的输入框里输入 LaTeX 公式，比如 `$$x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}$$`；

2. 在网页下部拷贝 URL Encoded 的内容，比如以上公式生成的是 `https://latex.codecogs.com/png.latex?%24%24x%3D%5Cfrac%7B-b%5Cpm%5Csqrt%7Bb%5E2-4ac%7D%7D%7B2a%7D%24%24`；

   ![](https://raw.githubusercontent.com/mzlogin/mzlogin.github.io/master/images/posts/markdown/latex-img.png)

3. 在文档需要的地方使用以上 URL 贴图，比如

   ```
   ![](https://latex.codecogs.com/png.latex?%24%24x%3D%5Cfrac%7B-b%5Cpm%5Csqrt%7Bb%5E2-4ac%7D%7D%7B2a%7D%24%24)
   ```

   示例效果：

   ![](https://latex.codecogs.com/png.latex?%24%24x%3D%5Cfrac%7B-b%5Cpm%5Csqrt%7Bb%5E2-4ac%7D%7D%7B2a%7D%24%24)

### 任务列表

在 GitHub 和 GitLab 等网站，除了可以使用有序列表和无序列表外，还可以使用任务列表，很适合要列出一些清单的场景。

示例代码：

```
**购物清单**

- [ ] 一次性水杯
- [x] 西瓜
- [ ] 豆浆
- [x] 可口可乐
- [ ] 小茗同学
```

示例效果：

**购物清单**

- [ ] 一次性水杯
- [x] 西瓜
- [ ] 豆浆
- [x] 可口可乐
- [ ] 小茗同学

## 转载

- [码志](https://mazhuang.org/)
