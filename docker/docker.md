#  docker

![75119937171](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751199371715.png)

![75119954555](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751199545554.png)



# 介绍

![75119990841](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751199908417.png)

docker的前身是LXC基于Linux内核的东西，进行隔离

![75119997351](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751199973518.png)

![75120001701](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751200017018.png)









#  docker入门

# Docker 基础

cnb 云原生开发环境中已经预装了 docker，无需学员手动安装，直接体验即可，可使用如下命令来查看 docker 信息

```shell
docker version  #查看版本信息
docker info     #查看运行时信息
```

## 运行你的第一个容器

```shell
docker run hello-world
# 不加地址默认去docker hub去找对象
#加地址就能从其他地方拉过来
```

跟所有的技术学习起点一样，上述命令使用 docker 运行一个最简单的 hello-world 容器，它的功能仅仅是在屏幕上打印一些文字。

`docker run` 首先会去本地寻找 hello-world 镜像，如果本地没有，则会从默认的 Docker 镜像仓库中拉去，也就是 [Docker Hub](https://hub.docker.com/)。

![75128503410](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751285034102.png)

从镜像拉出代码，打印应该hello-world

![75128511674](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751285116745.png)

虚拟机解决了硬件分配的问题，形成机群

docker做了环境隔离和分享机制，扩大了自己的影响力

![75128586825](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751285868252.png)

这些是linux内核拥有的东西，进行了封装

namespace

容器本身进行相关的隔离,本身就是一个被隔离的东西

先创建一个特殊进程

![75128608298](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751286082986.png)

这个是对docker的控制,进行资源限制



- 这样就变成了一个容器

![75128614799](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751286147998.png)

docker也想成携带所有环境的进程

把文件分成,对真正变化的东西进行下载,形成联合文件类型





镜像的格式一般为

```text
<repository>/<image>:<tag>
```

如果 repository 为空，默认为 Docker Hub, tag 为空，则默认为 latest，
如下是一个从 cnb 制品库中的镜像示例, 其中 `repository` 为 docker.cnb.cool，
`image` 为 looc/git-cnb，`tag` 为 latest。

```text
docker.cnb.cool/looc/git-cnb:latest
```

镜像分为 public 和 private 两种，对于 public 的镜像无需登录即可拉取，对于 private 的镜像则需要登录后才能拉取，登录命令如下

```shell
docker login <repository>
```

## 实践案例: 运行 Alpine Linux 容器

Alpine 镜像在企业生产环境中被广泛应用，它是一个极简的 Linux 发行版，
只包含最基本的命令和工具，因此镜像非常小，只有 5MB 左右，并且内置包管理系统 `apk`, 使其成为许多其他镜像的常用起点。

拉取镜像

``` shell
docker pull alpine  #拉取镜像
docker image ls     #查看镜像
```

运行容器

``` shell
docker run alpine ls -a  #运行容器
docker ps -a             #查看容器
```

交互式运行容器

docker run 命令默认使用镜像中的 Cmd 作为容器的启动命令，Cmd 可通用如下命令来查看。

``` shell
docker inspect alpine --format='{{.Config.Cmd}}'
# 查看相关的镜像信息
```

可以看到默认的 Cmd 为 `["/bin/sh"]`，因此直接使用 `docker run alpine` 会启动 /bin/sh 这个 shell,
我们 期望可以在这个 shell 中执行一些命令，但实际上它只是启动了 shell，退出了 shell，然后就停止了容器。

``` shell
docker run -it alpine  #等效于 docker run -it alpine /bin/sh, 使用 -it 参数启动容器进入交互式终端
```

后台运行容器

`run -it` 命令会启动一个交互式终端，退出终端后容器也会停止，如果希望容器在后台运行，可以使用 `-d` 参数，如下命令会启动一个后台运行的容器。

``` shell
docker run -it -d alpine #后台运行容器
```

加上 -d 参数后，容器不会执行完命令后立即退出，而是会进入后台运行，此时会返回一个唯一的 ID，使用 `docker ps` 命令可以查看容器运行状态。





docker attach 用于连接到一个正在运行的容器，主要作用是 访问容器的主进程（PID=1）的标准输入输出流

``` shell
docker attach <container_id>
```

由于 attach 是接管了 PID=1 的进程，因此如果这个进程是守护进程，
那么 attach 退出后，容器也会退出。所以一般不推荐使用 attach 命令。
而是使用 `docker exec` 命令来连接容器。

``` shell
docker exec -it <container_id> /bin/sh
```

此时进入到容器中使用`ps -a`命令可以看到容器中存在两个进程，其中 PID=1 的进程为 /bin/sh，
而另一个 /bin/sh 进程则是我们通过 exec 命令启动的，这个进程退出不会影响 PID=1 的进程，也就不会导致容器的退出。





- 有了docker可以标准化使用软件,不用你去配置,打包,编译了,使用docker可以把镜像打入进去
- 容器起来以后,还是需要环境的,镜像就是环境文件
- cnb的使用






####  docker和镜像基础

![75173265142](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751732651427.png)

首先交互式运行应该alpina容器

> docker run -it alpina

然后再去执行一些命令，比如安装应该软件，然后退出容器

> apk update
>
> 更新索引
>
> apk add figlet
>
> 安装相关的工具
>
> figlet "hello"
>
> 文本展示工具
>
> exit
>
> 退出当前的工具

这样就在alpina容器里面安装了figlet工具,当燃,可以安装一些更由于的软件,比如git,niginx,然后通过分享给别人,使用commit将容器保存为有个镜像

> docker ps -a//查看容器
>
> docker commit <container_id>

这样,我们就创建了有个装有figlet的镜像,我们可以通过docker images命令查看

> docker image is

上一个命令拿到了新创建的镜像的id,将其重新tag为有个alpina-figlet

> docker tag <image_id> alpina-figlet

然后我们就可以使用这个新的镜像了

> docker run alpine-figlet figlet "hlelo"

最后也可以使用docker push 命令将镜像推送到镜像仓库,其他人也可以使用docker pull来使用这个镜像了

![75173396251](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751733962514.png)



####  dockerfile核心机制

![75173414041](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751734140413.png)

dockerfile构建核心,引入了git,复用机制好

这个是一种更灵活的镜像创建方式

使用docker build构建镜像

> docker build -t alpine-figlet-from-dockerfile

同时可以使用这个镜像

> docker run alpine-figlet-from-dockerfile figlet "hello"

![75173487063](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751734870630.png)

重新构建镜像,并运行

![75173490869](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751734908699.png)

####  指令实战和讲解

- 指令

  - FROM指定基础镜像

  - RUN执行构建命令

  - COPY/ADD文件操作

  - WORKDIR设置工作目录

- 运行时的指令

  - EXPOSE端口声明
  - CMD容器启动命令
  - 环境变量动态配置env





例子

#### 使用 Dockerfile 构建一个 jupyter notebook 镜像

接下来让我们使用 Docker 来构建一个真实可用的镜像，比如 jupyter notebook 镜像。[Dockerfile](./jupyter_sample/Dockerfile)

```shell
docker build -t jupyter-sample jupyter_sample/
```

该镜像使用 RUN 指令来安装 jupyter notebook，使用 WORKDIR 指令设置工作目录，
使用 COPY 指令将代码复制到镜像中，使用 EXPOSE 指令来暴露端口，
最后使用 CMD 指令来启动 jupyter notebook 服务。

使用上述镜像来启动 jupyter notebook 服务。

```shell
docker run -d -p 8888:8888  jupyter-sample
```

我们使用了 -p 参数来将容器内的 8888 端口映射到宿主机的 8888 端口，在 cnb 上我们可以通过添加一个端口映射来实现外网访问。

点击这个浏览器图标，就可以访问 jupyter notebook 服务了。

####  使用多阶段构建来打包一个 golang 应用

在实际开发中，我们经常需要构建 golang 应用。
如果使用传统的单阶段构建，最终的镜像会包含整个 Go 开发环境，导致镜像体积非常大。
通过多阶段构建，我们可以创建一个非常小的生产镜像。

创建一个 [main.go](./golang_sample/main.go) 文件，
一个普通构建的 [Dockerfile](./golang_sample/Dockerfile.single)
以及一个多阶段构建的 [Dockerfile](./golang_sample/Dockerfile.multi)

构建镜像：

```shell
docker build -t golang-demo-single -f golang_sample/Dockerfile.single golang_sample/
docker build -t golang-demo-multe -f golang_sample/Dockerfile.multi golang_sample/
```

运行容器：

```shell
docker run -d -p 8080:8080 golang-demo-single
docker run -d -p 8081:8081 golang-demo-multe
```

容器运行成功后可以通过如下命令行来访问，可以看到两个容器都是在运行我们写的 golang 服务。

```shell
curl http://localhost:8080
curl http://localhost:8081
```

让我们来对比一下单阶段构建和多阶段构建的区别：

![75176164462](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751761644628.png)

WORKDIR /root/

COPY --from=builder /app/server

EXPOSE ${PORT}

CMD ["./server"]

多阶段和单阶段构建

 单阶段会打包整个环境,多阶段只是基础的东西

多阶段有体积优化

![75176216064](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751762160645.png)

![75176221527](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751762215278.png)

优化方向

![75176232282](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751762322825.png)

你会发现最终的镜像只有几十 MB，而如果使用单阶段构建（直接使用 golang 镜像），镜像大小会超过 1GB。这就是多阶段构建的优势：

- 最终镜像只包含运行时必需的文件
- 不包含源代码和构建工具，提高了安全性
- 大大减小了镜像体积，节省存储空间和网络带宽

这种构建方式特别适合 Go 应用，因为 Go 可以编译成单一的静态二进制文件。在实际开发中，我们可以使用这种方式来构建和部署高效的容器化 Go 应用。







####  docker的储存和网络

**储存**

1. 回顾：

   docekr生命周期管理

   基本指令

   apk包管理器

   dockerfile的优势

2. 储存方式

   ![75195588449](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751955884495.png)

   临时储存和volumes、挂载储存

   1. **默认存储**：容器内的数据随容器删除而丢失
   2. **Volumes（卷）**：由 Docker 管理的持久化存储空间，完全独立于容器的生命周期
   3. **Bind Mounts（绑定挂载）**：将主机上的目录或文件直接挂载到容器中

3. ![75195596457](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751955964573.png)

   场景测试

   ```shell
   # 运行一个 nginx 容器
   docker run -d --name web-default -p 8000:80 nginx

   # 在容器中创建一个测试页面
   docker exec -it web-default sh -c 'echo "<h1>Hello from Default Storage</h1>" > /usr/share/nginx/html/index.html'

   # 访问页面验证内容
   curl http://localhost:8000

   # 删除容器
   docker rm -f web-default

   # 用同样的配置重新运行容器
   docker run -d --name web-default -p 8000:80 nginx

   docker ps
   # 查看后台的容器列表

   # 再次访问页面，会看到默认的 Nginx 欢迎页面，之前的内容已经丢失
   curl http://localhost:8000

   ```
   临时的数据储存需求适合这个，提高了容器间数据隔离特性，确保每个容器拥有独立的数据环境，避免交叉污染

4. Volunmes卷储存

   ![75195651092](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751956510926.png)

   ```shell
   # 创建一个 Docker volume
   docker volume create nginx_data
   # 可以是自动创建的
   # 运行 Nginx 容器并挂载卷
   docker run -d --name web-volume -p 8081:80 -v nginx_data:/usr/share/nginx/html nginx
   # nginx_data是卷名后面跟着目录

   # 在容器中创建一个测试页面
   docker exec -it web-volume sh -c 'echo "<h1>Hello from Volume Storage</h1>" > /usr/share/nginx/html/index.html'

   # 访问页面验证内容
   curl http://localhost:8081

   # 删除容器
   docker rm -f web-volume

   #通过docker volume ls可以看到所有卷

   # 用同样的配置重新运行容器
   docker run -d --name web-volume-2 -p 8081:80 \
      -v nginx_data:/usr/share/nginx/html nginx
      
   # 再次重新挂载

   # 再次访问页面，内容仍然存在,是储存在数据卷里面的内容
   curl http://localhost:8081

   # 查看卷的详细信息
   docker volume inspect nginx_data
   ```
   ![75195661514](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751956615147.png)

   卷可以完成数据卷的共享

   多个容器共享一个卷,可能会发生同时读写的问题

   数据储存在docker管理区域，安全性比较好

5. 绑定挂载

   ![75195726733](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751957267338.png)

   使用本机的数据和储存功能

   ```shell
   # 创建本地目录
   mkdir nginx-content
   echo "<h1>Hello from Bind Mount Storage</h1>" > nginx-content/index.html
   ```


   # 运行 Nginx 容器并挂载本地目录
   docker run -d --name web-bind \
      -p 8082:80 \
      -v $(pwd)/nginx-content:/usr/share/nginx/html nginx
   # 相对路径会报错

   # 访问页面验证内容
   curl http://localhost:8082

   # 在主机上修改文件
   echo "<h1>Updated content from host</h1>" > nginx-content/index.html

   # 无需重启容器，直接访问更新后的内容
   curl http://localhost:8082

   # 删除容器
   docker rm -f web-bind

   # 用同样的配置重新运行容器
   docker run -d --name web-bind-2 -p 8082:80 \
      -v $(pwd)/nginx-content:/usr/share/nginx/html nginx

   # 再次访问页面，内容仍然存在
   curl http://localhost:8082
   ```
   ![75195752455](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751957524551.png)

   开发调试会快一点

   ![75195778956](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751957789560.png)

   ![75195781226](C:\Users\zxh\Desktop\前端\docker\docker.assets\1751957812266.png)

6. 清理操作

   完成实验后，可以进行清理：

   ```shell
   # 清理容器
   docker rm -f web-default web-volume web-volume-2 web-bind web-bind-2

   # 清理卷
   docker volume rm nginx_data

   # 清理本地目录
   rm -rf nginx-content
   ```

   清理容器，清理卷，清理本地目录

7. 使用Volume部署MySQL数据库

   **创建并管理Volume**

   ```shell
   docker volume create mysql_data

   docker volume inspect mysql_data

   docker volume ls
   ```

   **使用Volume运行MySQL**

   ```shell
   # 运行 MySQL 容器并挂载卷
   docker run -d \
     --name mysql_db \
     -e MYSQL_ROOT_PASSWORD=mysecret \
     -v mysql_data:/var/lib/mysql \
     mysql:8.0
     
   docker exec -it mysql_db mysql -uroot -pmysecret -h127.0.0.1

   mysql> CREATE DATABASE test_db;
   mysql> USE test_db;
   mysql> CREATE TABLE users (id INT, name VARCHAR(50));
   mysql> INSERT INTO users VALUES (1, 'John Doe');
   mysql> exit
   # 这些操作要进入MySQL完成
   ```

   **删除容器再创建来验证数据完整性**

   ​



、



####  docker网络基础和实践

![75259337946](C:\Users\zxh\Desktop\前端\docker\docker.assets\1752593379464.png)

![75259339496](C:\Users\zxh\Desktop\前端\docker\docker.assets\1752593394964.png)、

![75259354396](C:\Users\zxh\Desktop\前端\docker\docker.assets\1752593543966.png)

都是一些比较简单的应用

Bridge网络

docker默认的bridge网络通过网桥实现容器间通信，其核心是宿主集上的docker虚拟交换机，所有未指定网络的容器会自动接入该网络

每次发请求都是从虚拟交换机发出

![75259374540](C:\Users\zxh\Desktop\前端\docker\docker.assets\1752593745404.png)基本命令介绍

- 桥接网络

  ![75259419911](C:\Users\zxh\Desktop\前端\docker\docker.assets\1752594199112.png)

  不同容器也支持名称访问了

  需要你去自定义

- 自定义桥接网络现在让我们看看自定义 bridge 网络的优势：

  ```shell
  # 创建自定义 bridge 网络
  docker network create \
      --driver bridge \
      --subnet=172.20.0.0/16 \
      --gateway=172.20.0.1 \
      my-bridge-network

  # 启动两个容器，连接到自定义网络
  docker run -d \
      --name custom-container1 \
      --network my-bridge-network \
      my-nginx

  docker run -d \
      --name custom-container2 \
      --network my-bridge-network \
      my-nginx

  # 现在可以通过容器名称访问
  docker exec -it custom-container1 curl http://custom-container2
  ```

- web和redis通信

  这个案例展示了在实际应用中如何使用自定义 bridge 网络：

  首先，创建一个简单的 Web 应用：

  ```
  # 构建web-app镜像
  docker build -t web-app web-app
  ```

  然后创建网络并运行容器：

  ```
  # 创建自定义 bridge 网络
  docker network create my-bridge-network

  # 启动 Redis 容器
  docker run -d \
      --name redis-server \
      --network my-bridge-network \
      redis:alpine

  # 启动 Web 应用容器
  docker run -d \
      --name web-app \
      --network my-bridge-network \
      -p 5000:5000 \
      web-app
  ```

  现在可以测试应用了：

  ```
  # 访问应用
  curl http://localhost:5000

  # 多次访问，观察计数器增加
  curl http://localhost:5000
  curl http://localhost:5000

  # 查看 Redis 中的数据
  docker exec -it redis-server redis-cli get hits
  ```

  清理环境：

  ```
  # 删除用到 my-bridge-network 网络的容器
  docker rm -f web-app redis-server
  docker rm -f custom-container1
  docker rm -f custom-container2
  # 删除自定义网络
  docker network rm my-bridge-network
  # 删除镜像
  docker rmi web-app redis:alpine
  ```





- host网络

  host移除了容器和docker主机之间的网络隔离。直接使用主机的网络线，从而获取最佳网络性能

  **特点：**

  - 最佳网络性能
  - 直接使用主机的网络栈
  - 没有网络隔离
  - 端口直接绑定到主机上

  ```shell
  # 使用 host 网络运行 Nginx
  docker run -d \
      --name nginx-host \
      --network host \
      my-nginx

  # 直接通过主机的 80 端口访问
  curl http://localhost:80

  # 因为使用了 host 网络，容器直接使用主机的 80 端口, 所以当我们再次启动一个 Nginx 容器时，会报端口冲突的错误
  docker run -d \
      --name nginx-host-2 \
      --network host \
      my-nginx
      
  docker logs nginx-host-2
  ```

​        host的报错比较隐晦



- 没有网络的环境

  完全禁用网络端口，使得容器在该网络模式下没有任何外部端口，处于完全隔离的网络环境

  就是 **none网络**

  实践案例：**使用 None 网络运行独立计算任务**

  ```shell
  # 运行一个计算密集型任务，不需要网络
  docker run --network none alpine sh -c 'for i in $(seq 1 10); do echo $((i*i)); done'

  ```

- 跨主机通讯overlay网络

  Overlay 网络是 Docker 用于实现跨主机容器通信的网络驱动，主要用于 Docker Swarm 集群环境。 它通过在不同主机的物理网络之上创建虚拟网络，使用 VXLAN 技术在主机间建立隧道，从而实现容器间的透明通信。 在 Overlay 网络中，每个容器都会获得一个虚拟 IP，容器之间可以直接通过这个 IP 进行通信， 而不需要关心容器具体运行在哪个主机上。这种网络类型特别适合于微服务架构、分布式应用以及需要跨主机通信的容器化应用， 例如分布式数据库集群、消息队列集群等。Overlay 网络支持网络加密，能确保跨主机通信的安全性， 同时还提供了负载均衡和服务发现等特性，是构建大规模容器集群的重要基础设施。

  **适合微服务的网络**

  ​

  ​