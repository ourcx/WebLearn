#  go-redis

> 安装go get github.com/go-redis/redis/v8

```go
var rdb *redis.Client
func init(){
    rdb:=redis.NewClient(&redis.Options{
        Addr:"localhost:6379",
        Password:"",
        DB:0,
    })
}


func main(){
    ctx:=context.Background()
    err:=rdb.Set(ctx,"key","value",过期时间)
    //过期时间为10就是10秒,为0就是不过期
    if err!=nil{
        panic(err)
    }
    
    //拿到对应的数据
    val,err:=rdb.Get(ctx,"key").Result()
    if err !=nil{
        panic(err)
    }
    fmt.Println(val)
    
    rbd.Do(ctx,"get","key").Result()
    
    
}
```





####  string类型

```go
//GetSet拿到旧值和更新值
oldval,err:=rdb.GetSet(ctx,"key","newvalue").Result()
if err!=nil{
    panic(err)
}

fmt.Println(oldval)


//MGet批量查询key的值
vals,err:=rdb.MGet(ctx,"key1","key2","key3").Result()
if err!=nil{
    panic(err)
}
fmt.Peintln(vals)

//MSet批量设置key和value
err:=rdb.MSet(ctx,"key1","value1","key2","value2","key3","value3").Result()
if err!=nil{
    panic(err)
}


//针对incr,incrBy针对一个key的数值进行递增操作
val,err:=rdb.Incr(ctx,"key").Result()
if err!=nil{
    panic(err)
}


//IncrBy函数可以指定每次递增多少
valBy,err:=rdb.IncrBy(ctx,"key",2).Result()
if err!=nil{
    panic(err)
}
//incrByFloat是自增浮点数的函数
//使用Decr,DecrBy完成自减的操作

//Del可以删除key,支持批量删除

rdb.Del(ctx,"key")
err:=rdb.Del(ctx,"key1","key2","key3").Err()
if err != nil{
    panic(err)
}

//Expire单独设置过期时间
rdb.Rxpire(ctx,"key",3*time.Second)
```





####  Hash类型

hash特别适合存储对象 

- Hset

​        根据key和field字段设置，field字段的值

![74676532423](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746765324232.png)

- HGet

  根据key和field字段，查询field字段的值

  ![74676640233](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746766402332.png)

- HGetAll

  根据key查询所有字段和值

  ![74676657468](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746766574682.png)

- HIncrBy

  根据key和field字段，累加字段的数值

  ![74676665342](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746766653427.png)

- HKey

  根据key返回所有字段名

  ![74676683558](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746766835588.png)

- HLen

  根据key，查询hash的字段数量

  ![74676704414](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767044142.png)

- HMGet

  根据key和多个字段名，批量查询多个hash字段值

  ![74676709536](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767095364.png)

![74676720373](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767203730.png)

![74676747420](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767474204.png)

![74676758073](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767580737.png)

![74676758968](C:\Users\zxh\Desktop\前端\go基础\go-redis.assets\1746767589680.png)



####  List





