# Java 内存区域详解

<!-- TOC -->

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

## JVM 运行时的数据区域

首先获取一个直观的认识：

![JVM运行时的数据区域.png](./pic/JVM运行时的数据区域.png)

总共也就这么 5 个区（直接内存不属于 JVM 运行时数据区的一部分），除了程序计数器其他的地方都有可能出现 OOM (OutOfMemoryError)，其中像是程序计数器和两个栈（Java 虚拟机栈 & 本地方法栈）都是每个线程要有一个的，所以肯定是线程隔离的。而其他 2 个区就是线程共享的了，也就是说，如果有多个线程要同时访问这两个区的数据，是会出现线程安全问题的。接下来，我们将对这些区域进行详细的介绍。

### 程序计数器

- 当前线程所执行的字节码的行号指示器，字节码解释器工作时就是通过改变这个计数器的值来确定下一条要执行的字节码指令的位置
- 执行 Java 方法和 native 方法时的区别：
	- 执行 Java 方法时：记录虚拟机正在执行的字节码指令地址；
	- 执行 native 方法时：无定义；
- 是 5 个区域中唯一不会出现 OOM 的区域。

### Java 虚拟机栈

- Java 方法执行的内存模型，每个方法执行的过程，就是它所对应的栈帧在虚拟机栈中入栈到出栈的过程；
- 服务于 Java 方法；
- 可能抛出的异常：
	- OutOfMemoryError（在虚拟机栈可以动态扩展的情况下，扩展时无法申请到足够的内存）；
	- StackOverflowError（线程请求的栈深度 > 虚拟机所允许的深度）；
- 虚拟机参数设置：`-Xss`.

### 本地方法栈

- 服务于 native 方法；
- 可能抛出的异常：与 Java 虚拟机栈一样。

### Java 堆

- 唯一的目的：存放对象实例；
- 垃圾收集器管理的主要区域；
- 可以处于物理上不连续的内存空间中；
- 可能抛出的异常：
	- OutOfMemoryError（堆中没有内存可以分配给新创建的实例，并且堆也无法再继续扩展了）。
- 虚拟机参数设置：
	- 最大值：`-Xmx`
	- 最小值：`-Xms`
	- 两个参数设置成相同的值可避免堆自动扩展。

### 方法区

- 存储已被虚拟机加载的类信息、常量、静态变量、即时编译器编译后的代码等数据；	
	- 类信息：即 Class 类，如类名、访问修饰符、常量池、字段描述、方法描述等。
- 垃圾收集行为在此区域很少发生；
	- 不过也不能不清理，对于经常动态生成大量 Class 的应用，如 Spring 等，需要特别注意类的回收状况。
- 运行时常量池也是方法区的一部分；
	- Class 文件中除了有类的版本、字段、方法、接口等描述信息外，还有一项是常量池，用于存放编译器生成的各种字面量（就是代码中定义的 static final 常量）和符号引用，这部分信息就存储在运行时常量池中。
- 可能抛出的异常：
	- OutOfMemoryError（方法区无法满足内存分配需求时）。

### 直接内存

- JDK 1.4 的 NIO 类可以使用 native 函数库直接分配堆外内存，这是一种基于通道与缓冲区的 I/O 方式，它在 Java 堆中存储一个 DirectByteBuffer 对象作为堆外内存的引用，这样就可以对堆外内存进行操作了。因为可以避免 Java 堆和 Native 堆之间来回复制数据，在一些场景可以带来显著的性能提高。
- 虚拟机参数设置：`-XX:MaxDirectMemorySize`
	- 默认等于 Java 堆最大值，即 `-Xmx` 指定的值。
- 将直接内存放在这里讲解的原因是它也可能会出现 OutOfMemoryError；
	- 服务器管理员在配置 JVM 参数时，会根据机器的实际内存设置 `-Xmx` 等信息，但经常会忽略直接内存（默认等于 `-Xmx` 设置值），这可能会使得各个内存区域的总和大于物理内存限制，从而导致动态扩展时出现 OOM。



## HotSpot 虚拟机堆中的对象

这一小节将对 JVM 对 Java 堆中的对象的创建、布局和访问的全过程进行讲解。

### 对象的创建（遇到一条 new 指令时）

1. 检查这个指令的参数能否在常量池中定位到一个类的符号引用，并检查这个符号引用代表的类是否已被加载、解析和初始化过。如果没有，先把这个类加载进内存；
2. 类加载检查通过后，虚拟机将为新对象分配内存，此时已经可以确定存储这个对象所需的内存大小；
3. 在堆中为新对象分配可用内存；
4. 将分配到的内存初始化；
5. 设置对象头中的数据；
6. 此时，从虚拟机的角度看，对象已经创建好了，但从 Java 程序的角度看，对象创建才刚刚开始，构造函数还没有执行。

第 3 步，在堆中为新对象分配可用内存时，会涉及到以下两个问题：

**如何在堆中为新对象划分可用的内存？**

- 指针碰撞（内存分配规整）
	- 用过的内存放一边，没用过的内存放一边，中间用一个指针分隔；
	- 分配内存的过程就是将指针向没用过的内存那边移动所需的长度；
- 空闲列表（内存分配不规整）
	- 维护一个列表，记录哪些内存块是可用的；
	- 分配内存时，从列表上选取一块足够大的空间分给对象，并更新列表上的记录；

**如何处理多线程创建对象时，划分内存的指针的同步问题？**

- 对分配内存空间的动作进行同步处理（CAS）；
- 把内存分配动作按照线程划分在不同的空间之中进行；
	- 每个线程在 Java 堆中预先分配一小块内存，称为本地线程分配缓冲（Thread Local Allocation Buffer，TLAB）；
	- 哪个线程要分配内存就在哪个线程的 TLAB 上分配，TLAB 用完需要分配新的 TLAB 时，才需要同步锁定；
	- 通过 `-XX:+/-UseTLAB` 参数设定是否使用 TLAB。

### 对象的内存布局

- 对象头：
	- 第一部分：存储对象自身运行时的数据，HashCode、GC分代年龄等（Mark Word）；
	- 第二部分：类型指针，指向它的类元数据的指针，虚拟机通过这个指针来判断这个对象是哪个类的实例（HotSpot 采用的是直接指针的方式访问对象的）；
	- 如果是个数组对象，对象头中还有一块用于记录数组长度的数据。
- 实例数据：
	- 默认分配顺序：longs/doubles、ints、shorts/chars、bytes/booleans、oops (Ordinary Object Pointers)，相同宽度的字段会被分配在一起，除了 oops，其他的长度由长到短；
	- 默认分配顺序下，父类字段会被分配在子类字段前面。

*注：HotSpot VM要求对象的起始地址必须是8字节的整数倍，所以不够要补齐。*

### 对象的访问

Java 程序需要通过虚拟机栈上的 reference 数据来操作堆上的具体对象，reference 数据是一个指向对象的引用，不过如何通过这个引用定位到具体的对象，目前主要有以下两种访问方式：句柄访问和直接指针访问。

#### 句柄访问

句柄访问会在 Java 堆中划分一块内存作为句柄池，每一个句柄存放着到对象实例数据和对象类型数据的指针。

优势：对象移动的时候（这在垃圾回收时十分常见）只需改变句柄池中对象实例数据的指针，不需要修改reference本身。

![对象的访问定位_句柄访问.png](./pic/对象的访问定位_句柄访问.png)

#### 直接指针访问

直接指针访问方式在 Java 堆对象的实例数据中存放了一个指向对象类型数据的指针，在 HotSpot 中，这个指针会被存放在对象头中。

优势：减少了一次指针定位对象实例数据的开销，速度更快。

![对象的访问定位_直接指针访问.png](./pic/对象的访问定位_直接指针访问.png)

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


#### 1.1 最简单的服务器

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

- 注解配置

```java
@Mapper
public interface UserDAO {
    String TABLE_NAME = " user ";
    String INSERT_FIELDS = " name, password, salt, head_url ";
    String SELECT_FIELDS = " id, " + INSERT_FIELDS;

    @Insert({"insert into ", TABLE_NAME, "(",  INSERT_FIELDS,
            ") values (#{name},#{password},#{salt},#{headUrl})"})
    int addUser(User user);

    @Select({"select ", SELECT_FIELDS, " from ", TABLE_NAME, " where id=#{id}"})
    User selectUserById(int id);

    @Select({"select ", SELECT_FIELDS, " from ", TABLE_NAME, " where name=#{name}"})
    User selectUserByName(String name);

    @Update({"update ", TABLE_NAME, " set password=#{password} where id=#{id}"})
    void updatePassword(User user);

    @Delete({"delete from ", TABLE_NAME, " where id=#{id}"})
    void deleteById(int id);
}
```

- XML配置

    - 在相同的包目录下定义的同名XML

    - XML相较于注解可以做一些简单的逻辑操作

```java
List<Question> selectLatestQuestions(@Param("userId") int userId,
                                     @Param("offset") int offset,
                                     @Param("limit") int limit);
```

```xml
<mapper namespace="com.nowcoder.dao.QuestionDAO">
    <sql id="table">question</sql>
    <sql id="selectFields">id, title, content, comment_count,created_date,user_id
    </sql>
    <select id="selectLatestQuestions" resultType="com.nowcoder.model.Question">
        SELECT
        <include refid="selectFields"/>
        FROM
        <include refid="table"/>

        <if test="userId != 0">
            WHERE user_id = #{userId}
        </if>
        ORDER BY id DESC
        LIMIT #{offset},#{limit}
    </select>
</mapper>
```

#### 2.5 ViewObject

使用vo整合实体(如问题)以及该实体相关信息(如发问题用户的信息和当前问题关注数)一并传递给模板

```java
public class ViewObject {
    private Map<String, Object> objs = new HashMap<>();

    public void set(String key, Object value) {
        objs.put(key, value);
    }

    public Object get(String key) {
        return objs.get(key);
    }
}
```

#### 3.1 拦截器(interceptor)

拦截器在链路上设置回调接口，所有的请求都会回调注册过的拦截器(相较于切面编程的好处是相关请求request和response等都已经包装好了)

![img](https://github.com/A11Might/A11Might.github.io/blob/master/img/nowcoder2016/3_1.jpg)

```java
public class PassportInterceptor implements HandlerInterceptor{
    // preHandle在处理controller前走拦截器，用于判断当前用户是否有权限访问当前页面
    boolean preHandle(HttpServletRequest var1, HttpServletResponse var2, Object var3) throws Exception;
    // 可以在拦截器中，直接返回false，结束该请求(controller之前就结束了)，用于请求提前返回，进行异常处理

    // postHandle 在渲染前走拦截器，用于将数据推入模板进行渲染
    void postHandle(HttpServletRequest var1, HttpServletResponse var2, Object var3, ModelAndView var4) throws Exception;

    // afterCompletion 所有请求结束后走拦截器，用于将之前生成的数据删除
    void afterCompletion(HttpServletRequest var1, HttpServletResponse var2, Object var3, Exception var4) throws Exception;
}
```

注册拦截器到链路上

```java
@Component
public class WendaWebConfiguration extends WebMvcConfigurerAdapter {
    @Autowired
    PassportInterceptor passportInterceptor;

    @Autowired
    LoginRequiredInterceptor loginRequiredInterceptor;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(passportInterceptor);
        registry.addInterceptor(loginRequiredInterceptor).addPathPatterns("user/*");
        super.addInterceptors(registry);
    }
}
```

#### 3.2 Ajax请求

ajax请求(异步的JavaScript请求)，不是提交到页面，而是提交到后台，后台返回一个json格式字符串后，再进行动态的刷新处理(翻页时不刷新页面，直接从json格式字符串中提取信息替换掉)

#### 3.3 敏感词过滤

```java
public String filter(String text) {
    if (StringUtils.isBlank(text)) {
        return text;
    }
    String replacement = DEFAULT_REPLACEMENT;
    StringBuilder result = new StringBuilder();

    TrieNode curNode = rootNode;
    int begin = 0;
    int position = 0;

    while (position < text.length()) {
        char c = text.charAt(position);
        if (isSymbol(c)) {
            // 如果符号在敏感词中间就跳过，否则加入
            if (begin == position) {
                result.append(c);
                begin++;
            }
            position++;
            continue;
        }
        curNode = curNode.getSubNode(c);
        // 以begin开始的字符串不是敏感词
        if (curNode == null) {
            result.append(text.charAt(begin));
            begin++;
            position = begin;
            curNode = rootNode;
        // 从begin到position之间的字符串是敏感词
        } else if (curNode.isKeywordEnd()) {
            result.append(replacement);
            position++;
            begin = position;
        //  从begin到position之间的字符串不确定是否是敏感词
        } else {
            position++;
        }
    }

    // 加入最后一段非完整敏感词文本
    result.append(text.substring(begin));

    return result.toString();
}
```

#### 3.4 多线程

- Thread

    1. extends Thread，重载run()方法
    2. implements Runnable()，实现run()方法

```java
new Thread(new Runnable() {
    @Override
    public void run() {
        Random random = new Random();
        for (int i = 0; i < 10; ++i) {
            sleep(random.nextInt(1000));
            System.out.println(String.format("T%d : %d", tid, i));
        }
    }
}, String.valueOf(i)).start();  
```

- Synchronized－内置锁

    1. 放在方法上会锁住所有synchronized方法
    2. synchronized(obj) 锁住相关的代码段

```java
public static void testSynchronized1() {
    synchronized (obj) {
        Random random = new Random();
        for (int i = 0; i < 10; ++i) {
            sleep(random.nextInt(1000));
        }
    }
}
```

- BlockingQueue 同步队列

多线程访问的队列，用于实现异步处理

- ThreadLocal

    1. 线程局部变量。即使是一个static成员，每个线程访问的变量是不同的
    2. 常见于web中存储当前用户到一个静态工具类中，在线程的任何地方都可以访问到当前线程的用户
    3. 参考HostHolder.java里的users

- Executor

Executor是一个任务框架，将各种任务提交进来，它会根据已有的线程(可以是一个线程或线程池)，一个个去执行

```java
ExecutorService service = Executors.newSingleThreadExecutor(); // 单线程任务框架
ExecutorService service = Executors.newFixedThreadPool(2); // 多线程任务框架
// 提交任务
service.submit(new Runnable() {
    @Override
    public void run() {
        for (int i = 0; i < 10; ++i) {
            sleep(1000);
            System.out.println("Execute %d" + i);
        }
    }
});

// 在任务结束后关闭线程
service.shutdown();
```

- Atomic(原子性操作)

在进行多线程操作时，线程安全

```java
// 线程安全变量
private static AtomicInteger atomicInteger = new AtomicInteger(0);

public static void testWithAtomic() {
    for (int i = 0; i < 10; ++i) {
        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    Thread.sleep(1000);
                    for (int j = 0; j < 10; ++j) {
                        System.out.println(atomicInteger.incrementAndGet());
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }
}
```

- Future

实现线程与线程之间的通信(线程与线程数据传递)

```java
public static void testFuture() {
    ExecutorService service = Executors.newSingleThreadExecutor();
    // 返回异步结果
    Future<Integer> future = service.submit(new Callable<Integer>() {
        @Override
        public Integer call() throws Exception {
            //Thread.sleep(1000);
            // 任务中抛出异常，阻塞等待任务返回结果(future.get())，可以获取线程中的Exception
            throw new IllegalArgumentException("异常");
            //return 1;
        }
    });

    service.shutdown();
    try {
        // 阻塞等待返回结果()
        System.out.println(future.get());
        // 阻塞等待返回结果()，有过期时间
        //System.out.println(future.get(100, TimeUnit.MILLISECONDS));
    } catch (Exception e) {
        e.printStackTrace();
    }
}
```

#### 3.5 redis

- redis.conf

```conf
# after 900 sec (15 min) if at least 1 key changed
save 900 1

# 存储的文件
dbfilename dump.rdb
```

RDB持久性按指定的时间间隔执行数据集的时间点快照

AOF持久性会记录服务器接收的每个写入操作，这些操作将在服务器启动时再次播放，以重建原始数据集

- Jedis

Java对redis接口的包装，Redis官方推荐的Java库，官方网站：https://github.com/xetorthio/jedis

为了使用连接池，需要根据业务将jedis的操作包装成service

```java
public long sadd(String key, String value) {
    Jedis jedis = null;
    try {
        jedis = pool.getResource();
        return jedis.sadd(key, value);
    } catch (Exception e) {
        logger.error("发生异常" + e.getMessage());
    } finally {
        if (jedis != null) {
            jedis.close();
        }
    }
    return 0;
}
```

- [Redis数据结构](https://github.com/A11Might/wenda/blob/master/src/main/java/com/nowcoder/util/RedisTest.java)

- redis的使用

PV、点赞、关注、排行榜、验证码、缓存、异步队列、判题队列
