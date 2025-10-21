# 微服务

指的是一种开发应用的架构形式

![img](https://image.fengfengzhidao.com/rj_0731/20231026111717.png)

将大型应用分解成多个独立的组件，每一个组件有各自的责任领域，在处理一个用户请求时，基于微服务的应用可能调用许多内部微服务来共同生成其响应

- 单体项目和微服务

> 多个独立的组件,指的就是微服务
>
> 独立服务一般有明确定义的api,用于通讯
>
> 独立服务通常由一个小团队负责
>
> 独立服务的开发,部署,运营维护,扩展都不应该影响其他服务
>
> 每个服务应该针对一组功能进行设计,专注于解决特定问题
>
> 不强调每个服务使用同一种语言
>
> 非常独立的东西

- 会更复杂啊,不易维护,所以还要使用很多组件:

  ![74368704348](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743687043487.png)![74368715509](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743687155090.png)



####  go的微服务实践:

1. 标准库/自研派系:

   **不要让框架束缚开发,框架限制了开发者的思维**

   rpc:http\tcp\gypc来进行通讯,rpc是一种网络远程调用

   ![74368740000](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743687400003.png)

2. Web框架派系:

   **gin比较简洁,更稳定,适配比较好**

   ![74368772206](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743687722069.png)

3. 大一统框架: 

   **go-zero**

   ![74368791655](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743687916554.png)

####  第一个helloworld:

-  安装go

- goctl:

  > go install github.com/zeromicro/go-zero/tools/goct@1.4.4


![74377578274](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743775782742.png)

![74377591562](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743775915623.png)

-  进入项目目录,输入go mod tidy进行初始化 所有依赖

- 使用相关的api,可以去生成一些代码

  > ​
  > ├── etc//相关配置
  > │   └── greet.yaml
  > ├── go.mod
  > ├── greet
  > │   └── greet.go//是主包的一部分，用于初始化配置、依赖和服务，并启动服务
  > ├── greet.go//主程序入口或 HTTP 处理器的实现。可能包含 `main` 函数，用于启动 gRPC 服务器或 HTTP 服务。
  > ├── greet.proto 
  > ├── internal
  > │   ├── config
  > │   │   └── config.go//加载和解析配置文件（如 `etc/greet.yaml`）
  > │   ├── logic
  > │   │   └── pinglogic.go//实现核心业务逻辑（如处理 `Ping` 请求）
  > │   ├── server
  > │   │   └── greetserver.go// 实现 gRPC 服务端接口（定义在 `greet_grpc.pb.go` 中）
  > │   └── svc
  > │       └── servicecontext.go//依赖注入容器，聚合配置、数据库连接、日志等资源
  >
  > |    |——model//数据库模型等东西
  >
  > └── pb//由 `protoc` 工具从 `greet.proto` 生成的代码，包含 Protobuf 消息结构体的序列化/反序列化逻辑。
  > ├── greet.pb.go
  > └── greet_grpc.pb.go

  ​




#### api和rpc的区别

**RPC是远程过程调用（Remote Procedure Call Protocol，简称RPC），像调用本地服务(方法)一样调用服务器的服务(方法)**

**API 是应用程序编程接口，它定义了不同软件组件之间相互交互的规范和协议**

![74494021294](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744940212941.png)







#### Etcd

是给高可用的分布式键值存储系统,主要用于共享配置信息和服务发现,采用Raft一致性算法来包扎数据的强一致性,并且支持对数据进行监视和更新

- 为什么使用

![74391797180](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1743917971804.png)

![img](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\20231026110344.png)

配置etcd之后的交互图







####  一个简单的dome引用

1. 先编写一个rpc的proto文件

2. 生成代码

   > goctl rpc protoc user/rpc/user.proto --go_out=user/rpc/types --go-grpc_out=user/rpc/types --zrpc_out=user/rpc/

3. 填写必要的逻辑

> 在对外的api里调用rpc
>
> 再去完善yaml配置
>
> 再运行这些东西

![74402766309](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744027663095.png)

####  api语法讲解

api文件是对这个服务所有的api的描述

服务名，函数名，路径，请求方法，请求参数，响应参数  

```go
//api语法
syntax=“v1”
//goctl的版本

info{
    author:"xxx"
    date:"2022"
    docs:"api"
}

type LoginRequest {
  UserName string `json:"userName"`
  Password string `json:"password"`
}

type Response {
  Code int    `json:"code"`
  Data string `json:"data"`
  Msg  string `json:"msg"`
}

type UserInfo {
  UserName string `json:"userName"`
  Addr     string `json:"addr"`
  Id       uint   `json:"id"`
}

type UserInfoResponse {
  Code int      `json:"code"`
  Data UserInfo `json:"data"`
  Msg  string   `json:"msg"`
}

@server{
    group:user
    //路由分组
    prefix:userapi/v1
    //	添加路由分组
    jwt: Auth
    //需要鉴权
}

service users {
  @handler login
    //这个是请求这个url会执行的参数
  post /api/users/login (LoginRequest) returns (Response)
    //请求体和响应体
    //出参和入参，用来处理一些相关的gin
  
  @handler userInfo
  get /api/users/info returns (UserInfoResponse)
}
// goctl api go -api v1.api -dir .
```

####  语法

```go
/**
 * api语法示例及语法说明
 */
// api语法版本
syntax = "v1"
// import literal
import "foo.api"
//导入api文件
// import group
import (
    "bar.api"
    "foo/bar.api"
)
info(
    author: "songmeizi"
    date:   "2020-01-08"
    desc:   "api语法示例及语法说明"
)
//info是对这个api服务的藐视文件
// type literal
//type就是结构体
type Foo{
    Foo int `json:"foo"`
}
// type group
type(
    Bar{
        Bar int `json:"bar"`
    }
)
// service block
//定义api服务
@server(
    jwt:   Auth
    group: foo
    //jwt	声明当前service下所有路由需要jwt鉴权，且会自动生成包含jwt逻辑的代码	jwt: Auth
//group	声明当前service或者路由文件分组	group: login
//middleware	声明当前service需要开启中间件	middleware: AuthMiddleware
//prefix	添加路由分组	prefix: /api
)
service foo-api{
    @doc "foo"
    @handler foo
    //对handler层的描述
    post /foo (Foo) returns (Bar)
}
```





生成好的代码直接go mod tidy生成依赖

 在user.go文件里面

```go
package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"xiaohaiyun/app/user/api/desc/user/internal/config"
	"xiaohaiyun/app/user/api/desc/user/internal/handler"
	"xiaohaiyun/app/user/api/desc/user/internal/svc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")
//指定了配置文件,yaml就是配置文件
func main() {
	flag.Parse()
    //解析

	var c config.Config
	conf.MustLoad(*configFile, &c)
    //将配置文件映射到config

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
    //配置一些依赖

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
    //注册路由,对应handler的文件

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
    //启动服务
}
```

在handler文件夹里面

```go
package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"xiaohaiyun/app/user/api/desc/user/internal/logic"
	"xiaohaiyun/app/user/api/desc/user/internal/svc"
	"xiaohaiyun/app/user/api/desc/user/internal/types"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    //传入svcviceContext文件
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
        //参数解析到请求体上面
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
        //业务逻辑
        //写好的业务逻辑会被传到handler
		resp, err := l.User(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
```

//handler层的文件是会去使用逻辑层的代码的

所以逻辑层的代码应该是更广阔一些

goctl model 为 goctl 提供的数据库模型代码生成指令，目前支持 MySQL、PostgreSQL、Mongo 的代码生成，MySQL 支持从 sql 文件和数据库连接两种方式生成，PostgreSQL 仅支持从数据库连接生成

####  生成model

1. 通过sql语句生成model

2. 通过数据库的表生成model

   ![74428175493](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744281754938.png)

   使用sh生成model

   > sh genModel.sh zero-demo user

   ```go
   package model

   import (
   	"github.com/zeromicro/go-zero/core/stores/cache"
   	"github.com/zeromicro/go-zero/core/stores/sqlx"
   )

   var _ UserModel = (*customUserModel)(nil)

   type (
   	// UserModel is an interface to be customized, add more methods here,
   	// and implement the added methods in customUserModel.
   	UserModel interface {
   		userModel
   	}

   	customUserModel struct {
   		*defaultUserModel
   	}
   )

   func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
   	return &customUserModel{
   		defaultUserModel: newUserModel(conn, c, opts...),
   	}
   }
   ```


   ```

   #####  在logic里面调用model

   ![74428366362](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744283663629.png)

   ![74428367729](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744283677290.png)

   把yaml的DB解析到config,最后在svc的配置里面

   ![74428373097](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744283730976.png)

   然后就可以使用UserMOdel这个变量

   在逻辑里面使用

   ![74428388211](C:\Users\zxh\Desktop\前端\go基础\go-zero.assets\1744283882113.png)

   查询数据库的操作,在判断错误的时候,它内部设计把数据库错误和你的错误区分开来了,要针对的去写相关的操作

   #####  model代码详解

   






























   ```