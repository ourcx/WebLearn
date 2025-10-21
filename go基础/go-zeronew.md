####  go-zero框架

go-zero集成了各种实践和web和rpc的框架，通过弹性设计保障了大并发服务的稳定性，经受了充分的实战检验

![75689653602](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756896536021.png)

![75689676365](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756896763653.png)

![75689706725](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756897067255.png)

go-zero帮助我们减少写重复的代码

让开发者更关注业务



**何时使用go-zero**

![75689710774](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756897107743.png)

微服务适合大型的项目，个人项目不适合，除了学习



学习go-zero的工具，etcd，mysql，protoc，goctl，**doker**





####  环境配置

![75689724830](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756897248308.png)

这些都是使用go-zero的基础

注意go的版本问题，看官方文档

 安装完成后，输入` goctl api new api` 

![75689789828](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756897898288.png)

修改完这个文件,cd到api,使用go run api.go来运行

这个是创建api服务

![75689811742](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756898117425.png)

创建rpc服务

`goctl rpc new rpc`

跑的时候要做一些修改







####  ETcd

![75689853224](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756898532241.png)

etcd的数据可靠性比redis更强

![75689867441](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756898674414.png)

![75689877150](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756898771506.png)

类似redis的数据库,没有很繁琐的操作







####  最简单的微服务

![75689970782](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756899707820.png)

**创建一个user的rpc服务**

![75690060081](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1756900600813.png)

在生成的go文件里面写具体的逻辑,然后进行调用

**创建video的api服务**

我们做的

1. 编写用户proto
2. 生成代码
3. 添加自己的逻辑
4. 编写视频微服务的api服务的pai文件
5. 生成代码
6. 完善依赖配置
7. 添加自己的逻辑










####  api

api文件就是对这个服务所有的api的描述

服务名，函数名，路径，请求方法，请求参数，响应参数



编写api文件

```go
type LoginRequest {
    UserName string 'json:"userName"
    Password string 'json:"password"
}
//一定要大写

type Response {
    Code int 'json:"code"
    Data string 'json:"data"
    Msg string 'json:"msg"
}

type UserInfo {
    UserName string 'json:"userName"' [
    Addr string 'json:"addr"
    Id uint 'json:"id"
    }

type UserInfoResponse {
    Code int 'json:"code"
    Data UserInfo 'json:"data"
    Msg string 'json:"msg"
    }

        service users{
            @handler login
            post /api/users/login (LoginRequest) returns (Response)
            @handler userInfo
            get /api/users/login  returns (UserInfoResponse)
            //没有入参就不写那个括号
            
        }
        
        
        //转化 gotcl api go -api user.api(目录) -air(到当前的目录)
        //生成对应的文件后去logic去编写对应的业务
        
```

这里的type写的会很频繁,所以要注意使用

去进行封装->

 简写

```go
type LoginRequest {
    UserName string 'json:"userName"
    Password string 'json:"password"
}
//一定要大写

type UserInfo {
    UserName string 'json:"userName"' [
    Addr string 'json:"addr"
    Id uint 'json:"id"
    }


        service users{
            @handler login
            post /api/users/login (LoginRequest) returns (string)
            @handler userInfo
            get /api/users/login  returns (UserInfo)
            //没有入参就不写那个括号
            
        }
        //就这样直接去进行装换
```

**响应封装**

不把code,data,msg写在api里面,我们通过封装统一响应,在统一响应里面去加上code data msg

在common/response/enter.go

```go
package response

import(
    "github.com/zeromicro/go-zero/rest/httpx",
    "net/http"
)

type Body struct{
    Code uint32 `json:"code"`
    Msg  string `json:"msg"`
    Data interface{} `json:"data"`
}


func Response(r *http.Request, w http.ResponseWriter,resp interface{},err error){
    if err == nil{
        //根据不同的错误码,返回不同的错误信息
        r:=&Body{
            Code:0,
            Msg:"成功",
            Data:resp,
        }
        httpx.WriteJson(w,http.StatusOK,r)
        return
    }
    errCode := uint32(10086)
    errMsg := "服务器错误"
    
    httpx.WriteJson(w,http.StatusBadRequest,&Body{
        Code:errCode,
        Msg: errMsg,
        DataL:nil
    })
}
```





**模板定制化**

官方给了修改模板的方式,避免每次生成都要去改

先全局搜一下handler.tpl这个文件

没有就用`goctl template init`这个命令进行生成

看文档,用起来很麻烦







####  api前缀

对于用户服务来说,api的前缀是/api/users

```go
@server(
    prefix:/api/users
)

service users{
    @handler login
    post /login (LoginRequest) returns (string)
    @handler userInfo
    get /info returns (UserInfoResponse)
}
```

**jwt验证操作**

```go
@server(
    prefix:/api/users
)

service users{
    @handler login
    post /login (LoginRequest) returns (string)
}

@server(
    prefix:/api/users
    jwt:Auth
)

service users{
    @handler userInfo
    get /info returns (UserInfoResponse)
}

//一个有检验,一个没有检验
```

后面生成以后,在users.yaml配置相关的属性

**签发jwt需要我们自己进行,但是验证jwt是已经配置好了的**

签发

```go
func (l *LoginLogic) Login(req *types.LoginRequest)(resp string,err error){
    auth := l.svcCtx.Config.Auth
    
    token,err:=jwts.GenToken(jwts.JwtPayLoad{
        UserID:1,
        Username:"发疯",
        Role:1,
    },auth.AccessSecret,auth.AccessExpire)
    if err != nil{
        return "" , err
    }
    return token,err
}
```

再去请求不带请求头的jwt就报401错误

解析

```go
func (l *UserInfoLogic)UserInfo()(resp *types.UserInfoResponse,err error){
    userId := l.ctx.Value("user_id").(json.Number)
    username := l.ctx.Value("username").(string)
    uid,_ := userId.Int64()
    
    return &types.UserInfoResponse{
        UserId: uint(uid),
        Username: username,
    }, nil
}
```

接收的时候要注意jwt解析出来的类型

**jwt失败的响应**

不能一直是401,对前端很不友好

![75713997686](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757139976867.png)





####  生成api文档

后端对外的api，需要文档

1. 安装goctl-swagger

   `go install github.com/zeromicro/goctl-swagger@latest`

2. 生成app.json

   如果没有doc目录,则需要创建

   `goctl api plugin -plugin goctll-swagger="swagger -filename app.json -host localhost:8888 -basepath /" -api v1.pai -dir ./doc`

3. 使用docker,查看这个swagger页面

   `doker run -d --name swg -p 8087:8080 -e SMAGGER_JSON=/opt/app.json -v D:\IT\go_project3\go_test\v1\api\doc\:/opt swaggerapi/swagger-ui`

也可以完善相关的api信息给前端

> swagger不好用

公司项目有自己的api平台

团队项目可以用apifox

个人的化就用swagger

要给接口加一点信息,不然看不懂

```go
@server(
    prefix:/api/users
)

service users{
    @doc(
        summary:"用户登录"
    )
    @handler login
    post /login (LoginRequest) returns (string)
}

@server(
    prefix:/api/users
    jwt:Auth
)

service users{
     @doc(
        summary:"获取用户信息"
    )
    @handler userInfo
    get /info returns (UserInfoResponse)
}
```









####  go-zero操作mysql

![75717134950](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757171349505.png)

使用这些去生成表操作文件、

![75717187657](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757171876570.png)

![75717211130](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757172111300.png)

 ![75717244176](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757172441766.png)

在逻辑哪里调用

![75717257676](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757172576769.png)

直接拿到这些增删改查的操作

调用查找的功能需要你自己实现

这些都是一些原生的操作，go-zero





####  go-zero综合gorm

![75717412101](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757174121014.png)

![75717413853](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757174138532.png)

![75717414434](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757174144348.png)





####  rpc服务

```protobuf
syntax = 'proto';
package user;
option go_package = "./user";

message UserInfoRequest {
    uint32 user_id = 1;
}
message UserInfoResponse {
    uint32 user_id = 1;
    string username = 2;
}

message UserCreateRequest {
    string username = 1;
    string password = 2;
}

message UserCreateResponse {
    string err = 1;
}

service user{
    rpc UserInfo(UserInfoRequest)returns(UserInfoResponse);
    rpc UserCreate(UserCreateRequest)returns(UserCreateResponse);
}
```

  go-zero封装好了那些grpc,客户端都生成了

svc文件是依赖注入的地方,用于调用其他grpc

etc写一些地址的配置

server是一些代码的服务端,go-zero把服务端代码写了,他直接调logic的函数了





####  rpc服务分组

写多个server在proto里面然后在生成的问卷里面，的server可以找到多个你写的server





####  rpc使用gorm

![75723379801](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757233798013.png)

这就是一个操作相关的一个方法，用于用户登录

**配置文件，添加mysql的相关配置**

![75723402609](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757234026093.png)

 添加对应的配置映射，类型![75723408532](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757234085324.png)

在服务依赖的地方进行注入

![75723419484](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757234194842.png)

![75723419963](C:\Users\zxh\Desktop\前端\go基础\go-zeronew.assets\1757234199633.png)

创建逻辑

就是一些常见的数据库操作

会有一些models定义





####  rpc结合api

### 核心架构：API 网关 + RPC 服务

在 go-zero 的设计哲学中，通常采用如下架构：

1. **API 网关层 (API Gateway)**：对外暴露 HTTP/RESTful 接口。它负责接收、校验和路由外部客户端的请求，然后**作为客户端去调用内部的一个或多个 RPC 服务**，并最终聚合结果返回给外部客户端。这一层用 `goctl api` 命令创建。
2. **内部服务层 (RPC Service)**：对内提供具体的业务逻辑实现。它们之间通过 gRPC 进行通信，不直接对外暴露。这一层用 `goctl rpc` 命令创建。

API 网关通过 **RPC 客户端** 来调用下游的 RPC 服务。

### 实现步骤详解

让我们通过一个经典的“用户服务”示例来完整走一遍流程。假设我们需要提供一个 `GET /user/:id` 的 API 来获取用户信息，而用户数据由后端的 `user` RPC 服务管理。

#### 第 1 步：定义 RPC 服务（内部微服务）

首先，我们需要创建并定义内部的 RPC 服务。

1. **创建 RPC 服务**：

   bash

   ```
   goctl rpc new user
   ```

   这会生成一个名为 `user` 的 RPC 服务项目模板，结构如下：

   text

   ```
   user/
   ├── etc/                 # 配置文件目录
   │   └── user.yaml
   ├── internal/           # 内部逻辑（业务核心）
   │   ├── config/         # 配置结构定义
   │   ├── logic/          # 业务逻辑层（重点）
   │   ├── server/         # RPC 服务器实现
   │   └── svc/            # 服务上下文（依赖注入）
   ├── pb/                 # ProtoBuf 文件及生成的代码
   │   └── user.proto      # （重点）服务接口定义文件
   └── user.go            # 服务主入口
   ```

2. **编写 Proto 文件 (pb/user.prto)**：
   这是定义服务接口和消息体的核心。

   protobuf

   ```
   syntax = "proto3";

   package user;
   option go_package = "./user";

   // 定义请求和响应结构
   message GetUserRequest {
     string id = 1;
   }
   message GetUserResponse {
     string id = 1;
     string name = 2;
   }

   // 定义 RPC 服务
   service User {
     rpc GetUser(GetUserRequest) returns (GetUserResponse);
   }
   ```

3. **生成代码**：
   修改完 `proto` 文件后，需要重新生成 Go 代码和 gRPC 代码。

   bash

   ```
   goctl rpc protoc pb/user.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
   ```

   这个命令会生成：

   - `pb/user.pb.go` 和 `pb/user_grpc.pb.go`：标准的 gRPC 代码。
   - 为 go-zero 生成的 RPC 服务器代码（在 `internal` 目录下）。

4. **实现业务逻辑 (internal/logic/getuserlogic.go)**：
   在自动生成的 `GetUserLogic` 结构体的 `GetUser` 方法中编写具体逻辑。

   go

   ```
   func (l *GetUserLogic) GetUser(in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
       // 这里编写你的业务逻辑，例如：
       // 1. 校验参数
       if in.Id == "" {
           return nil, errors.New("id is required")
       }
       // 2. 查询数据库或其他操作
       // userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
       // ...

       // 3. 返回结果
       return &pb.GetUserResponse{
           Id:   in.Id,
           Name: "Username for " + in.Id, // 模拟数据
       }, nil
   }
   ```

5. **启动 RPC 服务**：
   运行 `go run user.go` 启动这个 RPC 服务，它会在指定的端口（如 `8080`）监听 gRPC 请求。

#### 第 2 步：定义 API 网关服务

现在，我们需要创建一个 API 网关来对外提供 HTTP 接口，并让它调用刚刚创建的 `user` RPC 服务。

1. **创建 API 服务**：

   bash

   ```
   goctl api new gateway
   ```

   这会生成一个 API 网关项目模板，结构如下：

   text

   ```text
   gateway/
   ├── etc/                # 配置文件
   │   └── gateway.yaml
   ├── internal/           # 内部逻辑
   │   ├── config/         # 配置结构定义
   │   ├── handler/        # HTTP 处理器（重点）
   │   ├── logic/          # 业务逻辑层（调用RPC的地方）
   │   ├── middleware/     # 中间件
   │   ├── types/          # 请求/响应结构体（重点）
   │   └── svc/            # 服务上下文（依赖注入，含RPC客户端）
   └── gateway.go          # 服务主入口
   ```

2. **编写 API 定义文件 (gateway.api)**：
   这是定义 HTTP 路由的地方。

   api

   ```go
   type GetUserRequest {
       Id string `path:"id"` // 从 URL 路径中获取参数
   }

   type GetUserResponse {
       Id   string `json:"id"`
       Name string `json:"name"`
   }

   service gateway {
       @handler GetUserHandler
       get /user/:id (GetUserRequest) returns (GetUserResponse)
   }
   ```

3. **生成代码**：

   bash

   ```
   goctl api go -api gateway.api -dir .
   ```

   这个命令会根据 `.api` 文件生成或更新 `handler`, `types`, `config` 等目录的代码。

#### 第 3 步：连接 API 网关和 RPC 服务（关键）

这是最核心的一步，让 API 网关能够找到并调用 RPC 服务。

1. **配置 RPC 客户端 (etc/gateway.yaml)**：
   在配置文件中添加 RPC 服务的连接信息（如 Etcd 地址或直接连接地址）。

   yaml

   ```
   Name: gateway
   Host: 0.0.0.0
   Port: 8888

   # 关键：配置 User RPC 服务
   UserRpc:
     # 方式一：通过 Etcd 服务发现（推荐）
     Etcd:
       Hosts:
         - localhost:2379
       Key: user.rpc
       #这里这个东西要和rpc哪里的配置是一样的
     # 方式二：直连（开发环境方便）
     # Target: 127.0.0.1:8080 # RPC 服务的监听地址
   ```

2. **注入配置和 RPC 客户端 (internal/config/config.go)**：

   go

   ```go
   type Config struct {
       rest.RestConf
       UserRpc zrpc.RpcClientConf // RPC 客户端配置，对应yaml中的`UserRpc`
   }
   ```

3. **创建 RPC 客户端并放入服务上下文 (internal/svc/servicecontext.go)**：

   go

   ```go
   type ServiceContext struct {
       Config   config.Config
       UserRpc  user.User // 这是由 pb/user_grpc.pb.go 生成的客户端接口
   }

   func NewServiceContext(c config.Config) *ServiceContext {
       return &ServiceContext{
           Config:  c,
           // 初始化 UserRpc 客户端
           UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
       }
   }
   ```

   `user.NewUser()` 函数是 go-zero 通过 proto 文件自动生成的，用于创建 RPC 客户端的 helper 函数。

4. **在 Logic 中调用 RPC 服务 (internal/logic/getuserlogic.go)**：
   API 网关的 Logic 层通过服务上下文 `svcCtx` 获取 RPC 客户端，然后像调用本地函数一样调用远程服务。

   go

   ```go
   func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (*types.GetUserResponse, error) {
       // 调用 UserRpc 的 GetUser 方法
       userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &pb.GetUserRequest{
           Id: req.Id,
       })
       if err != nil {
           return nil, err
       }

       // 将 RPC 的响应转换为 API 网关的响应类型
       return &types.GetUserResponse{
           Id:   userResp.Id,
           Name: userResp.Name,
       }, nil
   }
   ```

#### 第 4 步：启动并测试

1. **启动服务发现（如果使用 Etcd）**：

   bash

   ```
   etcd
   ```

   并在你的 User RPC 服务中配置并注册到 Etcd。

2. **启动 RPC 服务**：

   bash

   ```
   cd user && go run user.go
   ```

3. **启动 API 网关**：

   bash

   ```
   cd gateway && go run gateway.go
   ```

4. **测试**：
   用 curl 或浏览器访问 API 网关：

   bash

   ```
   curl -i http://localhost:8888/user/123
   ```

   输出应为：

   json

   ```
   {"id":"123","name":"Username for 123"}
   ```

   这个请求的完整流程是：
   `curl -> API Gateway(8888) -> (gRPC Call) -> User RPC Service(8080) -> DB -> Return`







