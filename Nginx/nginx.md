#  nginx的安装和虚拟机的配置

1. 使用vm创建一个虚拟机，

2. 虚拟机固定网络ip

3. nginx常用版本(四大阵营)

   > nginx开源版本
   >
   > nginx plus商业版本
   >
   > openresty
   >
   > tengine(淘宝网)

4. nginx的功能：反向代理（把用户的请求转到不同的服务器上，用户是无感知nginx的），负载均衡

![75239972467](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752399724673.png)

并且占用内存非常小，甚至支持5w个并发链接



####  开源版本安装

是什么

轻量级的web服务器以及电子邮件代理服务器

1. 下载

2. 把nginx安装包传上去给linux,然后解压安装

   > 进到哪个目录里.使用编译安装(ls)去安装,它需要一些依赖(c语言\perl库\zlib库)

   然后执行

   make

   make install

3. 启动nginx

   进入安装目录里面:/usr/local/nginx/sbin

   ![74626682168](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746266821680.png)

   输入相关的命令

   ps aux|grep n ginx

   (没有相应的脚本,只能手动去启动)


![74650982035](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746509820350.png)

启动之后可以进行一些ssl的配置，这样服务器可以使用https来进行访问



4. 关于防火墙

   ![74626927032](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746269270321.png)

   内网基本不用防火墙，主要是外网用防火墙 





####  nginx目录结构

1. sbin

   存放了启动的主文件

2. conf

   配置文件和一些程序 

   nginx.conf  

3. html

   一些前端的站点文件

4. logs

   日志文件 ，可能会有存满的情况

####  nginx打包上线你的项目

![74627199333](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746271993339.png)

优化库的命令

可以生成一个项目分析的内容，将不需要的代码（比如mock.js这种虚拟数据文件）注销掉，防止出现在打包中

####  cdn加速

![74650604795](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746506047952.png)

打包的时候排除一些库，使用外链传入

使用cdn来访问这些文件，能有更好的优化

![74650693800](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746506938008.png)

然后再去上线





####  nginx工具

![74650760266](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746507602667.png)

在服务器的环境里进行运行，然后服务器有地址，可以被访问 

![74650787476](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746507874765.png)

打开下载好的文件夹，点击exe是window里面的

将打包的文件放到html那个文件夹里面

![74651762758](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1746517627587.png)

如果有多个项目，要在config里面做一个区分，保证独立性













####  反向代理

- 正向代理

  VPN这种

  ![75239987327](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752399873270.png)

  帮客户端代理的正向代理，帮服务器代理的就是反向代理

- 反向代理


####  负载均衡

轮询

![75247371719](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752473717194.png)

加权轮询

权重高的服务器收到的比较多

![75247372922](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752473729223.png)

iphash

![75247376199](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752473761994.png)

让用户的消息一定达到某个服务器，完成session共享，但现在一般使用redis来去完成共享

![75247383595](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752473835950.png)

静态和动态的数据进行分离



####  配置文件

全局配置

到最大链接数

到http80端口的配置

https443的配置

在这些里面还有一些

![75249166376](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752491663761.png)

特别的配置

 ![75249296176](C:\Users\zxh\Desktop\前端\Nginx\nginx.assets\1752492961767.png)

