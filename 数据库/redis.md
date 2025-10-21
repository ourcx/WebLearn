# redis  

#### 介绍

> ------

#### NoSQL非关系型数据库：

1. 非结构化

   ![74385933666](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1743859336665.png)

2. 非关联的(数据库本身不会帮你维护数据,要自己处理)

 ![74385961345](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1743859613454.png)

3. 非SQL(没有固定的语法格式去查询)
4. 非事务(不满足acid的特性)

> 例子:键值类型的redis\文档类型的mongodb\列类型的hbase\图类型的noe4j
>
> ![74385997234](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1743859972341.png)
>
> ​                                                                sql                                           nosql

#### 认识redis

 ![74386197846](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1743861978464.png)

####  安装redis

redis基本都是基于linux来安装的,基本没有windows的方式,看文档来进行安装

记得设置开机自启动

#### 3.3.1.依赖库

Redis是基于C语言编写的，因此首先需要安装Redis所需要的gcc依赖：

```
yum install -y gcc tcl
```

#### 3.3.2.上传安装包并解压

然后将课前资料提供的Redis安装包上传到虚拟机的任意目录：

 

例如，我放到了/usr/local/src 目录： 

解压缩：

```
tar -xzf redis-6.2.6.tar.gz
```

解压后：

进入redis目录：

```
cd redis-6.2.6
```

运行编译命令：

```
make && make install
```

如果没有出错，应该就安装成功了。

默认的安装路径是在 `/usr/local/bin`目录下：

该目录已经默认配置到环境变量，因此可以在任意目录下运行这些命令。其中：

- redis-cli：是redis提供的命令行客户端
- redis-server：是redis的服务端启动脚本
- redis-sentinel：是redis的哨兵启动脚本

#### 3.3.3.启动

redis的启动方式有很多种，例如：

- 默认启动
- 指定配置启动
- 开机自启

#### 3.3.4.默认启动

安装完成后，在任意目录输入redis-server命令即可启动Redis：

```
redis-server
```





#### redis客户端

- 命令行客户端
- 图形化桌面客户端
- 编程客户端

\



####  数据结构介绍

![74426785688](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744267856887.png)

####  通用命令

1. KEYS：查看符合模板的所有key

![74426887907](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744268879078.png)

底层模糊查询，对服务器压力很大，不建议在生产环境下使用

2. DEL：删除一个指定的key
3. EXISTS：指定key是否存在
4. EXPIRE：给一个key设置有效期限，有效期到了会自动删除
5. 使用help[command]可以查看一个命令的具体用法
6. TTL查看一个key的剩余有效期



####  string类型

字符串类型，是redis里面最简单的存储类型

其value是字符串，不过根据·字符串的格式不同，可以分为：

- string：普通字符串，可以自增、自减
- int：整数类型，可以自增、自减
- float：浮点类型，可以自增、自减

三种方式只是编码方法不同，但都不能超过512m的最大空间

![74432882898](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744328828981.png)

- 示例：

![74432912232](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744329122327.png)

####  对key进行分类

![74433052381](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744330523819.png)![74433091090](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744330910903.png)

####  hash类型

hash类型，也是散列，其value是一个无序字典，类似于java的hashmap结构，

string结构是将对象序列化为JSON字符串后存储，当需要修改对象某个字段时很不方便

![74433174319](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744331743197.png)

- 常见命令，根据上面那个表来做

  ![74433191910](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744331919100.png)

  ​

####  List类型

这个list底层可以看为一个双向链表，支持正向检索和反向检索

- 有序
- 元素可以重复
- 插入和删除速度快
- 查询速度一般

######  查询方法

![74441574862](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744415748627.png)

 指定等待时间的BLOPOP和BRPOP





####  Set类型

一个value为null的hashMap。因为也是一个hash表，因此具备与HashSet类似的特征

- 无序
- 元素不可重复
- 查找快
- 支持交际并集差集功能

![74441787114](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744417871146.png)



####  SortedSet类型

可排序的set类型集合，里面的每个元素都带有score属性，可以基于score属性对元素进行排序，底层实现一个跳表加hash表

- 可排序
- 元素不重复
- 查询速度快

多用于实现排行榜这种东西

![74441907494](C:\Users\zxh\Desktop\前端\数据库\redis.assets\1744419074949.png)



























