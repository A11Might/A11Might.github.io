---
date: '2019-01-15'
draft: false
title: '爬取贴吧图片'
---

一直有在贴吧白嫖漫画的习惯，正好最近又在看Python爬虫，于是无聊就想把每个帖子里的漫画都下载到本地。思路是先分析爬取一页帖子内的漫画图片，然后再分析爬取精品栏内所有帖子的漫画图片。

>目标：爬取我的英雄学院百度贴吧精品贴漫画图片
>
>工具：Python Chrome

#### 第一步：爬取当前目标帖子内的所有图片

目标帖子URL：<https://tieba.baidu.com/p/5993286218?see_lz=1>

开启只看楼主，发现页数不超过一页，所以不用翻页操作，只需爬取当前一页图片即可。打开Chrome开发者工具，观察发现网页原始返回信息中就包含图片链接，所以直接使用requests库GET请求目标帖子URL，再使用PyQuery库解析返回的网页源码即可获取图片链接，然后再下载图片到本地。

使用Chrome开发者工具分析网页源码，可以发现：
> 帖子所有楼层信息都存在一个class="p_postlist"的div标签中
>
> 每层楼的信息存在class="l_post"的div标签中
>
> 每层楼中图片链接是class="d_post_content"的div标签中img标签的src属性

分析完源码后开始编写爬虫代码：

- 定义get_one_page(url)方法，使用requests库获取当前目标页面的网页源码（try-except防止网页请求错误，导致程序终止）

  ```python
  import requests
  from requests.exceptions import RequestException

  def get_one_page(url):
      try:
          url = url
          response = requests.get(url)
          return response.text
      except RequestException:
          print('第'+page_number+'页，网页请求失败')
  ```

- 定义parse_one_page(html)方法，使用pyquery库解析当前网页源码。根据上面的分析，使用css选择器**'#pb_content .p_postlist .l_post .BDE_Image'**获取全部img标签，并返回img标签的生成器（方便后面循环调用）

  ```python
  from pyquery import PyQuery as pq

  def parse_one_page(html):
      doc = pq(html)
      images = doc('#pb_content .p_postlist .l_post .BDE_Image').items()
      return images
  ```

- 定义get_title(html)方法，同样使用pyquery库获取当前帖子标题，方便后续把每个帖子的图片存入名称为帖子名的文件夹中

  ```python
  from pyquery import PyQuery as pq

  def get_title(html):
      doc = pq(html)
      title = doc('#j_core_title_wrap > h3').text()
      return title
  ```

- 定义download_images_to_folder(images,html)方法，使用获取的图片链接下载图片到名称为帖子名的文件夹中。首先获取帖子名称建立文件夹，再循环遍历parse_one_page(html)方法返回的生成器，获取每个img标签的src属性，并下载该图片到文件夹中，图片名按顺序为1-n

  ```python
  import requests
  from requests.exceptions import RequestException
  from pyquery import PyQuery as pq
  import os

  def download_images_to_folder(images,html):
      title = get_title(html)
      if not os.path.exists(title):
          os.mkdir(title)
      try:
          i = 1
          for image in images:
              response = requests.get(image.attr.src) 
              image_path = '{0}/{1}.jpg'.format(title,i)      
              with open(image_path,'wb') as f:
                  f.write(response.content)
              print('保存'+title+'第'+str(i)+'张图片成功')
              i = i + 1
      except RequestException:
          print('保存图片失败')
  ```

- 定义main()方法测试一下，成功获取目标网页图片

  ```python
  def main():
      html = get_one_page(url)
      download_images_to_folder(parse_one_page(html),html)

  if __name__ == '__main__':
      main()
  ```

运行结果：
![Crepe](/images/baidu_one_page.png)

#### 第二步：爬取贴吧精品栏中所有帖子的图片

目标帖子URL：<https://tieba.baidu.com/f?kw=%E6%88%91%E7%9A%84%E8%8B%B1%E9%9B%84%E5%AD%A6%E9%99%A2&ie=utf-8&tab=good>

在我的英雄学院吧的精品栏中发现，帖子前缀名为*【雄英支援科】* 的帖子是含有漫画图片的，每页精品栏的URL的差别为末尾的参数`&cid=&pn=offset` ，offset是50的倍数，所以只要获得每页精品栏中所有前缀名为*【雄英支援科】* 的帖子URL即可。

同样分析目标网页源码：
> id="thread_list"的ul标签包含该页所有帖子的简略信息
>
> class="j_thread_list"的li标签包含每个帖子的简略信息
>
> 每个帖子链接是class="threadlist_title"的div标签中a标签的href属性
分析完源码后开始编写爬虫代码：

- 同上获取并解析目标URL的源码。使用css选择器**'#thread_list .j_thread_list .threadlist_title a'**获取所有包含链接信息的a标签，调用它的href属性获取帖子链接的后半部分，再通过字符串的组合得到完整链接后，判断是否是含有漫画图片的帖子链接（直接通过名字中是否含有【雄英支援科】判断），返回每个含有漫画图片帖子URL的生成器

  ```python
  import requests
  from pyquery import PyQuery as pq

  def get_one_page_url(first_url):
      html = get_one_page(first_url)
      doc = pq(html)
      items_a = doc('#thread_list .j_thread_list .threadlist_title a').items()  
      for a in items_a:
          href = a.attr.href
          everyone_url = 'https://tieba.baidu.com'+ href +'?see_lz=1'
          if a.attr.title[:7] == '【雄英支援科】':
              yield (everyone_url)
  ```

- 定义main()方法，爬去所有爬取精品贴中所有帖子的漫画图片。通过观察精品贴不同页数的链接规律，使用offset参数构造每页精品贴的链接，再调用多线程快速爬取。

  ```python
  from multiprocessing.pool import Pool

  def main(offset):
      url_good = 'https://tieba.baidu.com/f?kw=%E6%88%91%E7%9A%84%E8%8B%B1%E9%9B%84%E5%AD%A6%E9%99%A2&ie=utf-8&tab=good&cid=&pn=' + str(offset)
      for url in get_one_page_url(url_good):
          html = get_one_page(url)
          download_images_to_folder(parse_one_page(html),html)

  if __name__ == '__main__':
      pool = Pool()
      pool.map(main,[i*50 for i in range(4)])
  ```

运行结果：
![Crepe](/images/baidu_all_page.png)

#### 总结

其实只是简单爬取了我的英雄学院贴吧精品栏中帖子前缀名是*【雄英支援科】* 的所有帖子中的漫画图片，并按帖子名称分文件夹保存。
[源码链接](https://github.com/A11Might/practicecode/blob/master/code/spider/TieBa.py)
