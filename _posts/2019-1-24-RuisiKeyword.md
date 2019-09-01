---
layout: post
title: 通过关键词搜索睿思电影URL
tags: [spider]
---

最近在看running man，非常喜欢里面的智孝，想着是演员就去睿思看看有什么出演的电影，发现只能搜索标题中的关键字。于是利用爬虫通过关键字搜索电影名及其链接URL。

### 第一步：爬取一个导航页中每个电影帖子的URL

目标URL：<http://rs.xidian.edu.cn/bt.php?mod=browse&c=10>

打开Chrome开发者工具发现：

> 所有电影帖子信息放在一个class='t1'的table标签中
>
> 每个电影帖子信息对应table标签的一个tbody标签
>
> 电影帖子的URL为tbody标签中class='common'的td标签中的a标签的href属性

分析完网页源码开始编写爬虫：

- 定义get_one_page(url)方法，使用requests库获取网页源码

```python
def get_one_page(url):
    try:
        reponse = requests.get(url,headers=headers)
        return reponse.text
    except RequestException:
        print('网页请求失败')
```

- 定义parse_one_home_page(html)方法，使用pyquery库解析一页导航页源码：直接右键 `class='t1'的table标签` 复制它的css样式，通过CSS选择器获取整个大的table标签，在使用items()方法得到它的生成器，依次遍历tbody标签，得到其中电影帖子的URL

```python
def parse_one_home_page(html_home):
    base_url = 'http://rs.xidian.edu.cn'
    doc = pq(html_home)
    tbodys = doc('#wp > div.layout > div.col-main > div > table a').items()
    for tbody in tbodys:
        if tbody.text()[:4] == '[电影]':
            url_simple = base_url + tbody.attr.href[1:]
```

### 第二步：得到每个电影的的简介内容

目标URL：<http://rs.xidian.edu.cn/forum.php?mod=viewthread&tid=986193>

通过解析第一步我们获得的每个电影帖子URL的源码，可以获得每个电影的简介，打开Chrome开发者工具分析：

> 贴内电影简介的所有内容在class='t_f'的td标签中

分析发现每个帖子的简介内容排版不同，对应td标签中的源码也不同，所以直接获取整个td标签中的所有文本作以简介内容

- 定义get_intro(html)方法，完成如上操作

```python
def get_intro(html_simple):
    doc = pq(html_simple)
    td = doc('.t_f')
    return td.text()#每个帖子简介的html格式不统一，但都位于class=t_f的td标签中
```

- 定义get_name(html)方法，方便知道获取的URL对应的电影名字

```python
def get_name(html_simple):
    doc = pq(html_simple)
    span = doc('#thread_subject')
    return span.text()
```

### 第三步：判断关键字是否在简介中

通过re库在简介内容中查找关键字，若在则打印电影名和对应的帖子URL

- 定义keyword_in_intro_or_not(intro)方法，用re库查找关键词

```python
def keyword_in_intro_or_not(intro):
    result = re.search(keyword,intro,re.S)
    if result:
        return True
```

- 重写parse_one_home_page(html)方法，得到获取的每个URL的简介内容，并打印搜索到关键词的简介内容所对应的电影名及其帖子URL

```python
def parse_one_home_page(html_home):
    base_url = 'http://rs.xidian.edu.cn'
    doc = pq(html_home)
    tbodys = doc('#wp > div.layout > div.col-main > div > table a').items()
    for tbody in tbodys:
        if tbody.text()[:4] == '[电影]':
            url_simple = base_url + tbody.attr.href[1:]
            html_simple = get_one_page(url_simple)
            if keyword_in_intro_or_not(get_intro(html_simple)):
                print(get_name(html_simple) + '\n' + url_simple)
```

### 第四步：搜索范围扩大到每页导航页

- 定义main(offset)方法，搜索所有睿思中的电影简介内容，并返回包含关键字的电影名及其URL

```python
def main(offset):
    url_home = url + '&page=' + str(offset)
    html_home = get_one_page(url_home)
    parse_one_home_page(html_home)
```

- 使用多线程

```python
if __name__ == '__main__':
    pool = Pool()
    pool.map(main,[i for i in range(100)])
```

运行结果：没有智孝还在保种的电影

![Crepe](/img/post/ruisi-namesearch.png){: .center-block :}

### 总结

借助睿思人规范发帖，通过 *[电影]* 前缀获取所需帖子URL，并一页简介一页简介的搜索有没有关键字存在。简单暴力低效，即使多线程也要花上些许时间。[源码链接](https://github.com/A11Might/SomePracticeCode/blob/master/spider/RuiSi.py)
