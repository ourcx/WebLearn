#  websocket

####  技术栈

![74467409354](C:\Users\zxh\Desktop\前端\websocate\websocket.assets\1744674093541.png)

####  使用

```javascript
//前端
((doc,Socket,stroage)=>{
    const oList = doc.querySelector('#list');
    const oMsg = doc.querySelector('#message');
    const oSendBtn = doc.querySelector('#send');
    const ws = new Socket('ws:localhost:8000');
    let username=''
    //实例化一个websocket
    
    const init=()=>{
           bindEvent();   
    }
    
    
    function bindEvent(){
  osendBtn.addEventListener('click',handleSendBtnClick,false);
    ws.addEventListener('open',handleOpen,false);
    ws.addEventListener('close',handleCloss,false);
    ws.addEventListener('error',handleError,false);
    ws.addEventListener('message',handleMessage,false);
    }
    
    
    function handleSendBtnClick(){
        const msg = oMsg.value;
        
        if(!msg.trim().length){
            return
        }
        
        ws.send({
            user:username,
            dataTime:new Date().getTime(),
            message:msg
        })
        
        omsg.value=''
    }
    
    function handleCloss(){
        
    }
    
    function handleOpen(e){
        username=storage.getItem("username")
        if(!username)return
        location.href=""
        return
    }
    
    function handleError(e){
        
    }
    
    function handleMessage(e){
        const msgData=JSON.parse(e.data)
        oList.appendChild(createMsg(msgData))
    }
    
    function createMsg(data){
         const{user,dateTime,messgae}=data;
        const oItem =doc.creatElemnet("li")
        oItem.innerHtml=`
<p>
<span>${user}</span>
<i>${new Data(dataTime)}</i>
</p>
<p>消息:${message}</p>
`
        return Oitem
    }
     
    
    init()
})(document,webSocket,localStorage)
```





```javascript
//后端//用nodejs//nodemon监控index.js文件
const Ws =require('ws')

;((ws)=>{
    const serve =new WS.server({port:8000});
    
    const init=()=>{
        
    }
    
    
    function bindEvent(){
        serve.on('open',handleOPen);
        serve.on('close',handleCloss);
        serve.on('error',handleError);
        serve.on('connection',handleConnection);
    }
    
    function handleOPen(){
        
    }
    
    function handleCloss(){
        
    }
    
    function handleError(){
        
    }
    
    function handleConnection(ws){
        
        ws.on('message',handleMessage)
    }
    
    function handleMessage(msg){
        server.clients.forEach(
            function(c){
                c.send(msg)
            })
        //记录了每一个客户端
    }
    
    init()
    
})(ws) 
```



vue版本的聊天室









  









####  go的websocket

> 使用包“github.com/gorilla/websocket”

websocket是http的升级形式

```go
//服务端
var p = websocket.Upgrader{
    ReadBufferSize:1024,
    WriteBufferSize:1024 
}

//先写一个http架构
func main(){
    http.HandleFunc("/",handler)
    http.ListenAndServe(":8888",nil)
}

func handler(w http.ResponseWriter,r *http.Requset){
    conn,err:=UP.Upgrade(w,r,nil)
    //链接
    if err!= nil{
        log.Println(err)
        return
    }
    conns = append(conns,conn)
    for{
        m,p,e:=conn.ReadMessage()
        //持续链接
        if e != nil {
            break
        }
        
        for i := range conns{
            conns[i].WriteMessage(websocket.TextMessage,[]byte('xxxxxx'))
        }
        
        
        conn.WriteMessage(websocket.TextMessage,[]byte('ssssss'))
        fmt.Println(m,string(p))
    }
    defrt conn.Close()
    //服务关闭
}
```

![74476071011](C:\Users\zxh\Desktop\前端\websocate\websocket.assets\1744760710119.png)



```go
//客户端
func main(){
    dl := websocket.Dialer{}
    conn,_,err :=dl.Dial("ws://127.0.0.1:8888",nil)
    if err!=mil{
        log.Println(err)
        return
    }
    conn.WriteMessage(websocket.TextMessage,[]byte("xxxx"))
    //拿到链接,开始传数据
    for{
        conn.ReadMessage(websocket.TextMessage,[]byte("xxxxxxxx"))
        //读取信息
    }
}
```









