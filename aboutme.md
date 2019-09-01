---
layout: page
title: About me
subtitle: Why you'd want to go on a date with me
---

我是胡启航，这是我的个人博客。

随便写写七七八八的东西。

欢迎给我留言。

## 联系

{% for website in site.data.social %}
* {{ website.sitename }}：[@{{ website.name }}]({{ website.url }})
{% endfor %}

## Skill Keywords

{% for category in site.data.skills %}
### {{ category.name }}
<div class="btn-inline">
{% for keyword in category.keywords %}
<button class="btn btn-outline" type="button">{{ keyword }}</button>
{% endfor %}
</div>
{% endfor %}
