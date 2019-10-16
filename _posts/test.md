<!-- TOC -->

- [重定向](#1.3-重定向)

- [最简单](#最简单的服务器)

- [Java 内存区域详解](#java-内存区域详解)
    - [JVM 运行时的数据区域](#jvm-运行时的数据区域)
        - [程序计数器](#程序计数器)
        - [Java 虚拟机栈](#java-虚拟机栈)
        - [本地方法栈](#本地方法栈)
        - [Java 堆](#java-堆)
        - [方法区](#方法区)
        - [直接内存](#直接内存)
    - [HotSpot 虚拟机堆中的对象](#hotspot-虚拟机堆中的对象)
        - [对象的创建（遇到一条 new 指令时）](#对象的创建遇到一条-new-指令时)
        - [对象的内存布局](#对象的内存布局)
        - [对象的访问](#对象的访问)
            - [句柄访问](#句柄访问)
            - [直接指针访问](#直接指针访问)

<!-- /TOC -->
# Document


- Part I

    - [1.1 最简单的服务器](#1.1-最简单的服务器)
    - [1.2 request/response](#1.2-request/response)
    - [1.3 重定向](#1.3-重定向)
    - [1.4 Error](#1.4-Error)
    - [1.5 IoC(控制反转)](#1.5-IoC(控制反转))
    - [1.6 AOP(面向切面编程)](#1.6-AOP(面向切面编程))

----

- Part II

    - [2.1 数据库字段的设计](#2.1-数据库字段的设计)
    - [2.2 数据库创建](#2.2-数据库创建)
    - [2.3 数据库操作(CRUD)](#2.3-数据库操作(CRUD))
    - [2.4 MyBatis](#2.4-MyBatis)
    - [2.5 ViewObject](#2.5-ViewObject)

----

- Part III

    - [3.1 interceptor(拦截器)](#3.1-拦截器(interceptor))
    - [3.2 Ajax请求](#3.2-Ajax请求)
    - [3.3 敏感词过滤](#3.3-敏感词过滤)
    - [3.4 多线程](#3.4-多线程)
    - [3.5 redis](#3.5-redis)


#### 最简单的服务器

![img](https://github.com/A11Might/A11Might.github.io/blob/master/img/nowcoder2016/1_1.jpg)

通过Controller定义网页入口，在网页入口指定访问地址，访问方法，通过访问地址和请求可以解析参数，有两种返回方式

- 直接返回字符串

```java
@RequestMapping(path = {"/profile/{groupId}/{userId}"})
@ResponseBody
public String profile(@PathVariable("userId") int userId, // 解析地址中的参数(/userId)
                      @PathVariable("groupId") String groupId,
                      @RequestParam(value = "type", defaultValue = "1") int type, // 解析请求的参数(?type=2)
                      @RequestParam(value = "key", required = false) String key) {
    return String.format("Profile Page of %s / %d, t:%d k: %s", groupId, userId, type, key);
}
```

- 渲染模板返回

```java
@RequestMapping(path = {"/vm"}, method = {RequestMethod.GET})
public String template(Model model) { // model将数据传递到模板中
    model.addAttribute("value1", "vvvvv1");
    List<String> colors = Arrays.asList(new String[]{"RED", "GREEN", "BLUE"});
    model.addAttribute("colors", colors);

    Map<String, String> map = new HashMap<>();
    for (int i = 0; i < 4; ++i) {
        map.put(String.valueOf(i), String.valueOf(i * i));
    }
    model.addAttribute("map", map);
    model.addAttribute("user", new User("LEE"));
    return "home";
}
```


#### request/response

HttpServletRequest、HttpServletResponse是网页请求request和服务器返回信息response的包装

- request

    参数解析、cookie读取、http请求字段、文件上传

- HttpServletRequest

    request.getHeaderNames();request.getMethod();request.getPathInf();request.getQueryString()

- response

    页面内容返回、cookie下发、http字段设置(headers)

- HttpServletResponse

    response.addCookie(new Cookie(key, value));response.addHeader(key, value);

#### 1.3 重定向

- http状态码301：永久重定向，意为旧的URL已经不在使用，已永久转移至新的地址

    永久跳转会从缓存中读取后跳转

- http状态吗302：临时重定向，意为某个时间段因为某些原因临时进行的跳转行为，旧的URL地址依然使用并存在

    重定向，通过redirect前缀跳转过去(默认是302临时跳转)

```java
@RequestMapping(path = {"/redirect/{code}"}, method = {RequestMethod.GET})
    public String redirect(@PathVariable("code") int code) {
        return "redirect:/";
    }
```

#### 1.4 Error

定义一个ExceptionHandler，抛出异常时，进行统一的异常处理

```java
// Spring MVC外的Exception或Spring MVC没有处理的Exception
@ExceptionHandler()
@ResponseBody
public String error(Exception e) {
    return "ERROR:" + e.getMessage();
}
```

#### 1.5 IoC(控制反转)

无需关注变量的初始化，只要使用注解表示其是来自bean池对象的初始化，就会自动设置上

```java
@Autowired
PassportInterceptor passportInterceptor;
```

#### 1.6 AOP(面向切面编程)

面向切面，所有业务都要处理的业务

![img](https://github.com/A11Might/A11Might.github.io/blob/master/img/nowcoder2016/1_6.jpg)

```java
@Aspect
@Component
public class LogAspect {
    private static final Logger logger = LoggerFactory.getLogger(LogAspect.class);

    // 通配符'*'，可以代表任何东西
    // '*'是返回值 com.nowcoder.controller.*Controller是类 *(..)方法以及一些参数
    // @Before是在执行所有Controller中方法之前执行这个方法
    @Before("execution(* com.nowcoder.controller.*Controller.*(..))")
    public void beforeMethod(JoinPoint joinPoint) {
        StringBuilder sb = new StringBuilder();
        for (Object arg : joinPoint.getArgs()) {
            if (arg != null) {
                sb.append("arg:" + arg.toString() + "|");
            }
        }
        logger.info("before method:" + sb.toString());
    }

    // @After是在执行所有IndexController中方法之后执行这个方法
    @After("execution(* com.nowcoder.controller.IndexController.*(..))")
    public void afterMethod() {
        logger.info("after method" + new Date());
    }
}
```

#### 2.1 数据库字段的设计

| id | name | password | salt | head_url |
|-----|-----|-----|-----|-----|
| 1 | Jim | ajsdf83d | kd8 | http://xx |
| 2 | Tom | adsf9nfd | xdi | http://yy |
| 3 | Lucy | gf2jnfpof | daf | http://zz |

关系型数据库中存储的是一张一张的表，每张表对应一种实体(用户，问题或者评论...)，如上为用户表，用于存储每一个用户的信息

- 每一行是一个实体(某一个用户)，对应数据库中的一条记录

- 每一列是实体的具体信息，如上用户表，包含用户姓名，密码，salt和头像

#### 2.2 数据库创建

使用workbrench图形化界面创建

int 整数类型；varchar(n) 可变字符；datetime 日期类型；float(m, d) 浮点类型；text 长字符串

#### 2.3 数据库操作（CRUD）

```sql
# 增
INSERT INTO table_name (列1, 列2,...) VALUES (值1, 值2,....)
# 查
SELECT */列名称/多个列名称 FROM 表名称
# 改
UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值
# 删
DELETE FROM 表名称 WHERE 列名称 = 值
```

```sql
# 选择语句
WHERE {CONDITION}
# 排序 升序/降序
ORDER BY {COL} ASC/DESC
# 分页
LIMIT {OFFSET},{COUNT}
# groups rows that have the same values into summary rows
GROUP BY {COL}
```

```sql
# 嵌套
select * from (select * from message order by created_date desc) tt group by conversatoin_id
# 返回指定列的值的数目
SELECT COUNT(column_name) FROM table_name
```

#### 2.4 MyBatis

- 注解
