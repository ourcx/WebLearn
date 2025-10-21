#  go文件操作

####  打开/关闭一个文件

```go
func main(){
    file,err := os.open(name,falg);
    //name是一个地址
    if err == nil{
        //文件打开成功
    }else{
        //文件打开失败
        return
    }
    defer file.Close()
    //挂起一个文件关闭的操作,防止忘记关闭文件的事情发生
    //os.open()返回文件指针和一个错误
    
}
```







####  带缓冲的文件读写

1. **以只读的方式打开一个文件,创建其带缓冲的读取器,逐行读取到末尾**

```go
func main(){
    //open就是对openfile的一个封装,简便调用
    file,err := os.OpenFile(name,flag,perm);
    //name是文件明,flag是读取方式,perm是权限,返回文件的指针和一个可能的错误
    if(err == nil){
        //创建成功
    }else{
        //失败
    }
    
    defer func(){
        file.Close()
    }
    
    //创建缓冲区
    reader:=bufio.NewReader(file)
    //读取器
    for{
        str err := reader.ReadString(delim)
    //一个字符串字符串的去读,读到第一个delim字节时返回读到的内容,delim是一个整型 返回string或者err
        if err == mil{
            fmt.println(str)
        }else{
            if err == io.EOF{
                //读到文件末尾不用在读了
            }
            break
        }
    }
    //文件读取完毕
      
}
```

> 关于文件读写的数字
>
> 这些权限位可以组合，形成不同的权限组合。例如，`0644` 表示用户有读写权限，组和其他用户有读权限。
>
> 那么`0400`这四位数字又分别代表什么，左起第一个`0`代表的是文件类型，`0`是普通类型,左起第二位`4`代表的是当前用户权限，第三位`0`代表的是用户组权限，第四位`0`是其他用户权限
>
> 同时:4是读权限,2是写权限,1是执行权限
>
> 
>
> 读取方式:
>
> ![74522798596](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1745227985968.png)
>
> go定义了一系列的读写代码,让你可以设置文件的读取方式
>
> 只读
>
> 只写
>
> 可读可写
>
> 追加
>
> 创建文件
>
> 文件不存在时才创建
>
> 不允许并发操作
>
> 覆盖



####  便捷读写

使用ioutil来实现便捷读写

```go
func main(){
    //读入指定文件的全部数据,返回[]byte类型的原始数据
    bytes,err:=ioutil.ReadFile("xxxx")
    if err == nil{
        contentStr := string(bytes)
        //成功
    }else{
        //错误
    }
    //这个包对文件的关闭做好了封装,不用你干
}
```



####  缓冲式写出

打开一个文件，缓冲式写出几行数据，缓冲区空了之后就退出

追加是在文件末尾加东西，覆盖是直接改整个文件

```go
func main(){
    //创建只写追加
    file,err := os.openFile("路径"，os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
    //错误处理
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    writer.WriteString("内容")
    writer。Flush()
    //把缓冲区全部写入文件
}
```





####  便捷写出

```go
func main(){
    //读入指定文件的全部数据,返回[]byte类型的原始数据
    //``反引号是保留原始格式,""是转化字符串,在``里面换行不用\n
    bytes,err:=ioutil.WriteFile("xxxx路径","写入的东西","权限")
    if err == nil{
        contentStr := string(bytes)
        //成功
    }else{
        //错误
    }
    //这个包对文件的关闭做好了封装,不用你干
}
```

####  判断文件是否存在

![74530007043](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1745300070437.png)





####  文件拷贝

方法

1. 简易的把文件读出来写到另外一个文件,比较简陋

2. 使用io.copy进行文件拷贝

   ```go
   func main(){
       scrFile,_:=os.OpenFile("xxx","xxxxx",0666)
       dstFile,_:=os.OpenFile("xxx","xxxxx",0666)
       //要去打开文件,没有就创建
       written,err := io.Copy(dstFIle,srcFile)
       //把srcFile拷贝给dstFile
       if err == nil{
           fmt.Println("xxxx")
       }else{
           //错误
       }
       //返回拷贝的字节数和错误,写入两个文件的东西
       
   }
   ```

3. 缓冲式文件拷贝:用于拷贝超大的文件，不能一次性全部拷进去

   ```go
   func main(){
       //打开源文件
       srcFile,_:=so.OpenFile("xxxx",os.O_RDONLY,0666)
       //打开目标文件
       dstFile,_:=os.OpenFile("xxx","xxxxx",0666)
       
       defer func(){
           srcFile.Close()
           dstFile.Close()
           fmt.PrintIn("文件全部关闭")
       }()
       
       
       //创建源文件缓冲读取器
       reader:=bufio.NewReader(srcFile)
       //创建目标文件的写出器
       writer:=bufio.NewWriter(dstFile)
       
       //创建小水桶缓冲池
       
       buffer:= make([]byte,1024)
       //1k大小的水桶
       
       //一部分一部分读入数据，直到io.EOF
       for{
           _,err := reader.Read(buffer)
           if err!=nil{
               if err==io.EOF{
                   //源文件读取完毕
                   break
               }else{
                   //读取发生错误
               }
               
           }else{
           _,err:=writer.Write(buffer)
           if err!=nil{
               //写入发生错误
               return
           }
       }
       }
       
       
   }
   ```

   ​


####  总结

使用的包有读写包,缓冲包,便携读写包

![74641313695](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1746413136958.png)

- 读写包

  ![74641331119](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1746413311190.png)

  perm是指定其读写权限等

  ![74641342760](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1746413427603.png)

  文件的状态和是否存在的判断

- 便携包

  ![74641354587](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1746413545875.png)

- 缓冲包

  ![74641358099](C:\Users\zxh\Desktop\前端\go基础\echo.assets\1746413580992.png)

  缓冲写入和缓冲读取

  ​





#  JSON数据读写

javascript对象标记语言
































