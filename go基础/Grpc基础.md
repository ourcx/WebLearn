#  Grpc基础

- 单体架构

  只能对整体扩容，伸缩

  代码耦合，项目开发者往往需要知道整个系统的流程

- 微服务架构

  可以按照服务进行单独扩容，升级

  各个服务之间可独立开发，独立部署

  问题在于代码冗余和服务之间的沟通问题 



####  rpc

 rpc是远程过程调用

就是让调用远程方法和调用本地方法一样

grpc就是一个具体的实现库

- 安装相关的调用库

  ![75660344774](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756603447748.png)

环境变量添加到path路径下

exe哪个文件一般是在gopath文件里面

然后编写一个hello.proto文件

```go
syntax = "proto3"; // 指定proto版本
package hello_grpc;     // 指定默认包名

// 指定golang包名
option go_package = "/hello_grpc";

//定义rpc服务
service HelloService {
  // 定义函数
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// HelloRequest 请求内容
message HelloRequest {
  string name = 1;
  string message = 2;
}

// HelloResponse 响应内容
message HelloResponse{
  string name = 1;
  string message = 2;
}

```

注意看我的路径

在grpc路径下执行

通过protobuf生成go文件

```shell
// protoc -I . --go_out=plugins=grpc:. .\hello.protoCopyErrorOK!
```

或者是在项目根目录下

```shell
// protoc -I . --go_out=plugins=grpc:.\grpc .\grpc\hello.protoCopyErrorOK!
```

![img](http://python.fengfengzhidao.com/pic/20230415154411.png)

如果生成了这个文件，那么恭喜你，难倒50%人的环境问题已经搞定了





####  例子

```go
//服务端
package main

import (
  "fmt"
  "net"
  "net/http"
  "net/rpc"
)

type Server struct {
}
type Req struct {
  Num1 int
  Num2 int
}
type Res struct {
  Num int
}

func (s Server) Add(req Req, res *Res) error {
  res.Num = req.Num1 + req.Num2
  return nil
}

func main() {
  // 注册rpc服务
  rpc.Register(new(Server))
  rpc.HandleHTTP()
  listen, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println(err)
    return
  }
  http.Serve(listen, nil)
}



//客户端
package main

import (
  "fmt"
  "net/rpc"
)

type Req struct {
  Num1 int
  Num2 int
}
type Res struct {
  Num int
}

func main() {
  req := Req{1, 2}
  client, err := rpc.DialHTTP("tcp", ":8080")
  if err != nil {
    fmt.Println(err)
    return
  }
  var res Res
  client.Call("Server.Add", req, &res)
    //从服务端请求到了一个方法
  fmt.Println(res)
}

```

原生rpc的问题,编写相对复杂要自己关注实现过程,没有代码提示 





####  写一个grpc

 写一个grpc服务端:hello world

> 服务端的实现
>
> 1. 编写一个结构体,名字不重要
> 2. 最重要的得实现protobuf中的所有方法
> 3. 监听端口
> 4. 注册服务

```go
//定义接口
syntax = "proto3";
package hello_grpc;//指定默认的包名
option go_package = "/hello_grpc";//生成go的时候的包名,还有其他的一些配置 .;xxx生成在当前目录下

//proto文件只是定义接口方法
service HelloService {
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
    string name = 1;
    string message = 2;
}

message HelloResponse {
    string name = 1;
    string message = 2;
}
```

现在只是写了一个接口,定义了一些内容,但是具体的实现还是得搞完HelloService

```go
package main

import (
	"context"
	"log"
	"net"

	pb "path/to/hello_grpc" // 替换为实际的go_package路径
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Name:    "Server",
		Message: "Hello " + req.Name + ", " + req.Message,
	}, nil
}
//注册方法

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
//调用接口,注册服务


```

grpc最好的地方在于跨语言



> 客户端的实现:
>
> 1. 建立链接
> 2. 调用方法

```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_study/grpc_proto/hello_grpc"
	"log"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := hello_grpc.NewHelloServiceClient(conn)
	result, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "枫枫",
		Message: "ok",
	})
	fmt.Println(result, err)
}
```

如果要经常使用这个链接,就可以把client的创建复用起来

目录结构

grpc_study/
├── go.mod
├── go.sum
├── grpc_proto/
│   ├── hello_grpc/
│   │   ├── hello_grpc.proto    # Protocol Buffer 定义文件
│   │   └── hello_grpc.pb.go    # 自动生成的代码 (由 protoc 生成)
│   └── compile_proto.sh        # 可选的编译脚本
├── server/
│   └── main.go                 # gRPC 服务器实现
└── client/
    └── main.go                 # gRPC 客户端代码

1. **hello_grpc.proto** - 定义服务接口和消息格式
   - 使用 `protoc` 工具编译生成 `hello_grpc.pb.go`
   - 生成的文件包含客户端和服务端所需的接口和结构体
2. **server/main.go** - 服务端实现
   - 导入生成的 `hello_grpc` 包
   - 实现 `HelloServiceServer` 接口
   - 监听指定端口并提供服务
3. **client/main.go** - 客户端代码
   - 导入生成的 `hello_grpc` 包
   - 创建连接并调用远程方法









####  proto文件

语法:service对应的就是go里面的接口,可以作为服务端,客户端

rpc对应的就是结构体中的方法

message对应的也是结构体

![75660710800](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756607108002.png)

对应关系

![75660736155](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756607361551.png)

![75660744410](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756607444109.png)

![75660745388](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756607453880.png)

go的行为，没有赋值就不会响应了

**数组**

![75660804065](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756608040657.png)

![75660805222](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756608052222.png)

![75660808879](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756608088793.png)

这些等于1等于2的是序号,相当于枚举



编写风格:

1. 文件名建议下划线，例如：my_student.proto
2. 包名和目录名对应
3. 服务名,方法名,消息名均为大驼峰
4. 字段名改为下划线



####  多服务

![75661121122](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756611211228.png)

在这里注册多个服务





####  多个proto文件

![75661124770](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756611247705.png)

非常任意遇到坑

多个proto的头部最好是一样的

![75662432887](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756624328875.png)

common.proto来存放request,response这些结构体

![75662558428](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756625584287.png)

其他proto文件就不用放这些结构体,然后用import引入进来,报错不用怕,因为没有对应的插件

![75662564058](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756625640586.png)

下一步就是去进行相关的编译

![75662580958](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756625809586.png)

编译命令

还要继续去生成另外几个go文件,那在同一个包,下面,就都不会报错了

然后再去写具体的函数实现





####  grpc流式传输

1. 一问一答式-普通的rpc进行交流
2. 客户端请求一次,服务端一直传输sse
3. 客户端流式,客户端一直请求发送消息
4. 双向的响应



**一问一答式**

![75662651232](C:\Users\zxh\Desktop\前端\go基础\Grpc基础.assets\1756626512326.png)

服务端客户端同上



**服务端流式**

```protobuf
syntax = "proto3";
option go_package = "/proto";
service Simple{
    rpc Fun(Request)return(Response){}
}

message Request{
    string name = 1;
}

message Response{
    string  Text = 1;
}

service ServiceStream{
    rpc Fun(Request)return(stream Response){}
}
//加了一个关键字stream就变成了流式响应了

```

服务端go文件

```go
type ServiceStrean struct{}

func (ServiceStream)func(request *proto.Request,stream proto.ServiceStream.FunServer) error {
    for i := 0;i<10;i++{
        stream.Send(&proto.Response{
            Text:fmt.Sprintf("第%d轮数据",i)
        })
    }
    return nil
    //这边要一直发送十轮数据
}

func main(){
    listen,err := net.Listen("tcp",":8080")
    if err != nil{
        //xxxx
    }
    server:=grpc.NewServer()                           proto.RegisterServiceStreamServer(server,&ServiceStream{})
    server,Serve(listen)
}
```

客户端go文件

客户端不知道服务端什么时候结束

```go

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
    client := proto.NewServiceStreamClint(conn)
    
    stream,err := client.Fun(context.Background(),&proto.Request{
        Name:'张三'
        //给服务端发一个请求
    })
    for i := 0;i<10;i++{
        response,err := stream.Recv()
        //接收数据
        fmt.Println(response, err)
    }
    //可以一直循环去遍历他,反正不知道什么时候结束
}
```

> 案例: 下载文件的时候,需要一直开启流式传输

 尽量不传输string,而是传bytes,因为文件

配合go的读写使用.缓冲读,缓冲写





**客户端流式**

```protobuf
syntax = "proto3";
option go_package = "/proto";
service Simple{
    rpc Fun(Request)return(Response){}
}

message Request{
    string name = 1;
}

message Response{
    string  Text = 1;
}

service ServiceStream{
    rpc Fun(Request)return(Response){}
}
//加了一个关键字stream就变成了流式响应了

service ClientStream{
    rpc UploadFile(stream FileRequest) retrun(Response){}
}
```

服务端

```go
type ClientStream struct{}

func (ClientStream)UploadFile(stream proto.ClientStream_UploadFileServer)error {
    for i:=0;i<10;I++{
        res,err:=stream.Recv()
        
    }
    stream.SendAndClose(&proto.Response{Text:"j结束"})
    return nil
    //这边要一直接收十轮数据
}

func main(){
    listen,err := net.Listen("tcp",":8080")
    if err != nil{
        //xxxx
    }
    server:=grpc.NewServer()                           proto.RegisterServiceStreamServer(server,&ClientStream{})
    server,Serve(listen)
}

```

客户端

```go
func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}

	defer conn.Close()
	// 初始化客户端
	client := proto.NewClientStreamClient(conn)

	stream, err := client.UploadFile(context.Background())
	for i := 0; i < 10; i++ {
		stream.Send(&proto.FileRequest{FileName: fmt.Sprintf("format 第%d次", 1)})
	}
}
```

反正就是一种缓存的读写的时候使用





####  双向流

proto文件

```protobuf
service BothStream {
    rpc Chat(stream Request)return(stream Response){}
}
```

服务端

```go
type ClientStream struct{}

func (b *BotnStream) Chat(stream proto.BothStream_ChatServer) error {
    for i := 0; i < 10; i++ {
        request, err := stream.Recv() // 修正了方法名 Rdev → Recv
        if err != nil {
            return err
        }
        fmt.Println(request)
        
        err = stream.Send(&proto.Response{ // 修正了方法名 Semd → Send
            Text: "你好",
        })
        if err != nil {
            return err
        }
    }
    return nil
}


func main(){
    listen,err := net.Listen("tcp",":8080")
    if err != nil{
        //xxxx
    }
    server:=grpc.NewServer()                           proto.RegisterServiceStreamServer(server,&ClientStream{})
    server,Serve(listen)
}
```

服务端

```go
func main() {
    addr := ":8080" // 添加冒号表示端口号

    // 使用 grpc.Dial 创建一个新的本地化的 gRPC 连接
    // 此处使用了不安全的证书来实现连接（仅用于测试）
    conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(fmt.Sprintf("grpc connect addr [%s] 失败: %s", addr, err))
    }
    defer conn.Close() // 修正方法名 Eccess → Close

    // 初始化客户端
    client := proto.NewBothStreamClient(conn) // 修正客户端名称 Botostream → BothStream
    
    stream, err := client.Chat(context.Background())
    if err != nil {
        log.Fatal("创建流失败: ", err)
    }
    
    for i := 0; i < 10; i++ {
        // 发送请求
        err := stream.Send(&proto.Request{ // 修正方法名 Remd → Send
            Name: fmt.Sprintf("第%d次", i), // 修正格式字符串
        })
        if err != nil {
            log.Fatal("发送请求失败: ", err)
        }
        
        // 接收响应
        response, err := stream.Recv()
        if err != nil {
            log.Fatal("接收响应失败: ", err)
        }
        fmt.Println(response, err) // 修正方法名 PrintIn → Println
    }
    
    // 关闭流
    stream.CloseSend()
}
```

