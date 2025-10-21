#  vite

![75954920933](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759549209337.png)

vite也支持直接构建react，angular，svelte项目

vite优势的冰山一角。



####  目录

![75954958859](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759549588591.png)

浏览器只人数html，css，js文件

1. ts文件：遇到这个要使用tsc将ts代码转化为js的代码
2. react/vue：要使用相关的编译器进行转化，react-compiler/vue-complier转化为render函数
3. less/sass/postcss/component-style又要安装相关的loader编译工具
4. 语法降级：babel-》将es的新语法转化为旧版浏览器接受的语法
5. 体积优化：uglifyjs-》将代码进行压缩，减小体积
6. ....

这些都是一个企业级项目可能会使用的,构建工具就是干这个的

还干了

1. 模块化开放,支持使用node-modules＋多种模块化支持
2. 处理代码兼容性,降级,语法转化(构建工具将对应的处理工具集成进来了)
3. 提高项目性能,压缩文件,代码分割
4. 优化开发体验,热更新,自动调用相关的打包工具,在浏览器重新运行
5. 开发服务器的跨域问题



> 甚至配置文件都不是必须的,让我们不用关心代码的构建,不用关心代码怎么在浏览器运行,只关注自己的业务开发



####  相对于webpack

快

![ 75955485179](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759554851796.png)

多种模块化方法会被转化成

![75955529046](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759555290464.png)

支持多种模块化就导致它在开始就要统一模块化，把全部文件读一遍，导致项目越来越大的时候，导致花费了很多时间在启动开发服务器

而vite要求使用esmodules来写，这两个构建工具侧重点不一样，vite更更关注浏览器的开发体验，webpack侧重兼容性



####  初体验

![75955757401](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759557574016.png)

vue-cli和create-vite的区别

暂时先不用管create-vite这个命令

vue-cli给我们提供好的预设，给我们需要的模板

![75955808488](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759558084884.png)

create-vite和vue-cli分的不是很清楚，因为他们是同一个团队的

![75955919248](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759559192480.png)

浏览器做了这个东西会影响性能，引入了太多的东西

所以要使用vite





####  vite依赖构建

找寻依赖的过程是自当前目录依次向上寻找的过程，直到搜寻到根目录或者搜寻到对应依赖为止

yarn dev --->开发(每次依赖项构建所构建的相对路径绝对正确)

生产环境会全权交给一个rollup的库去完成生产环境的打包

可能会导致路径错误

![75955996622](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759559966228.png)

**依赖预构建**:首先vite会找到对应的依赖,然后调用esbuild(对js语法进行处理的一个库),将其他规范的代码转化为esmodule规范,然后放到当前目录下,node_modules/.vite/deps,把模块进行统一集成

1. 解决了不同的第三方包会有不同的导出格式
2. 对路径的处理可以直接使用.vite/deps,方便路径重写
3. 网络多包传输的性能问题,这个也是浏览器(esmodule)不敢支持node_modules的原因,不管有多少个额外的export和import,vite都会尽可能的将他们进行集成

![75956051488](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759560514883.png)

直接把外部模块变成一个函数,浏览器就不知道是多模块了



####  vite配置文件

1. 配置文件的语法提示

   - 使用webstorm有很好的语法补全

   - 类型标注,写给i编译器看的

     ![75956100289](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759561002891.png)

     ​

2. 关于环境的处理

   过去我们使用webpack的时候,我们要区分一个配置文件的一个环境

   -webpack.dev.config

   -webpack.prod.config

   -webpack.base.config

   -webpackmerqe

   vite可以

   ![75956141215](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759561412159.png)

   不同环境走不同的代码,在vite.config.js,传扩展运算符或者传空括号

   或者使用if...else进行判断


####  vite环境变量

方便不同环境的切换

![75959712857](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759597128576.png)

 dotenv会自动提取.env文件，并解析这个文件的对应的环境变量，并将其注入到process对象下（但是vite考虑到和其他配置的一些冲突问题，它不会直接注入到process对象下）

涉及到root和envDIr配置

envDir用来配置当前环境变量文件的地址

 我们可以调用vite提供的方法来确认配置文件

![75959798531](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759597985316.png)

 这样env就拿到配置文件了

![75959840662](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759598406625.png)

yarn dev --mode development会将mode设置成development传递进来

当我们调用nodeenv的时候，会做：

1. 找到。env文件，并解析其中的环境变量，放到对象里面

2. 会将传进来的mode的值进行拼接，并根据我们提供的目录去取相关配置文件，并放入一个对象

3. ![75959872576](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759598725766.png)

   新提取的会替换1里面提取的值

   重名替换

   ![75959878183](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759598781835.png)

   这三个文件是关键

上面就是服务器对环境的区分



客户端就是：

通过import.meta.env进行注入环境变量

但是vite做了一个拦截,防止将隐私变量送到import.meta.env里面,不是以VITE开头的,他就不会注入到客户端

前缀可以修改

![75959909514](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759599095144.png)



![75959727644](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759597276443.png)

配置文件可以使用esmodule的原因



####  vite怎么启动开发的

vite的开发服务器

实现

```javascript
const Koa = require("Koa")
const fs = require("fs")
const path = require("path")
//node内置模块直接request
//不同的宿主环境会给js不同的能力
//不能用esmodule必须使用commonjs
const app = new Koa()

//当请求来临以后会直接进入到use注册的回调函数里面
app.use(async(ctx)=>{
    //context上下文
    console.log("ctx",,ctx.request,ctx.response)
    if(ctx.request.url === "/"){
        //意味着需要根路径
        const indexCountent = await fs.promises.readFile(path.resolve(__dirname,"./index.html"))
        //一般在服务端不会这样使用,而是使用文件流防止占用太多内存
        ctx.response.body = indexCountent
        //json --> aplication/json text/html text/javascript
        ctx.response.set("Content-Type","text/html")
        //使用html去解析
    }
    if(ctx.request.url === "/api/getUserInfo"){
        
    }
    fi(ctx.request.url === "./App.vue"){
        const indexCountent = await fs.promises.readFile(path.resolve(__dirname,"./index.html"))
        //如果是vue文件,会做一个字符串替换,匹配到template这种会进行全部字符串替换,ast,虚拟节点那种处理
        //假设这里处理完成了,就可以直接返回给浏览器
        ctx.response.body = indexCountent
        ctx.response.set("Content-Type","text/javascript")
        //强制使用js解析
        //在浏览器和服务器眼里,你的文件都是字符串,只是能否解析的区别
    }
    //....
    //实际上不要那么难去取,实际里面有中间件完成这些处理
    
})

app.listen(5174,()=>{
   //开启这个服务器
    
})

```

koa是node端的框架

服务端可以处理后直接去返回.vue文件,因为已经被vue编译过了









####  vite对css的解析和处理

vite天生就支持对css文件的直接处理

1. vite读取到main.js引用的index.css
2. 直接去使用fs模块读取文件内容
3. 直接创建一个style标签,将idnex直接copy进去
4. 将这个标签插入到html的head中
5. 将css文件中的内容替换成js脚本,方便热更新和模块化
6. 设置为content-type为js，从而让浏览器以js脚本的形式来执行该css



场景

![75965112671](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759651126712.png)

所以需要cssmodule去解决这个问题

命名文件成xxxx.module.css来完成模块化

再这样进行使用

![75965126133](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759651261337.png)

> 大概原理是
>
> 会对css类名进行一些处理,module.css是一个约定
>
> 添加了哈希名
>
> 同时创建了一个映射对象
>
> 将替换后的内容塞进style标签然后放到head标签里面
>
> 将对应css文件替换成js脚本
>
> 将创建的映射对象在脚本中进行默认导出

less要使用的话就要去安装







####  vite对处理css的配置

![75965287563](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759652875634.png)

配置类名转化.驼峰和非驼峰

![75965360882](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759653608821.png)

一些在css的modules要放置的配置

![75965370298](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759653702986.png)

具体的使用方法





####  css预处理器配置流程

![75971455881](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759714558811.png)

![75971458927](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759714589279.png)

假设我们的代码被压缩过了，这时候程序出错，会显示是压缩后的代码的报错，如果设置了sourceMap，就会有一个索引文件，导航到源文件的某一行报错

![75971483618](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759714836186.png)







####  postcss

![75971528776](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759715287767.png)

postcss主要是为了讲解浏览器的兼容问题，浏览器是讲解不了的

同样的babel是让js执行起来万无一失

 

1. postcss让高级的css语法降级成更低的css语法
2. 前缀补全，让不同浏览器适配

![75971759473](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759717594736.png)

使用

目前来说,less和sass的postcss插件已经停止维护了,但postcss是后处理器,sass,less编译完后才进行处理,但它能做的事情很多

> yarn add postcss-cli postcss -D

书写描述文件

![75971940120](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759719401201.png)

vite配置postcss

postcss的配置是在css下面进行配置的

![75972012927](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759720129275.png)

自己去建postcss.config.js文件也是可以的

使用的一次额未来特性的css不是给less，sass处理，而是给postcss进行处理







####  为什么服务端一定要使用path

![75991645418](C:\Users\zxh\Desktop\前端\vite\vite.assets\1759916454180.png)

node会自动拼接。。

会注入一些变量__dirname是当前的目录

path本质是一个字符串处理模块



####  vite处理静态资源

静态资源：除了动态api都被视为静态资源

在前端就是图片和文件



> 如果不使用vite,其他构建工具会将导入作为一个字符串进行解析
>
> vite会对你的资源进行摇树优化，控制你的导入，把没有用到的东西进行去除

现在很多视频和图片，是流文件传输，一开始只是传输buffer二进制字符串给你

在vite.base.config.js配置别名路径

![76000905595](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760009055952.png)







####  别名,resolve.alias

![76001065448](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760010654489.png)

这里没有做相对路径的处理，但是实际上path是这样做的

#####  vite对svg的处理

vite对svg仍然是开箱即用的

svg是可伸缩矢量图形

1. svg不会失真
2. 尺寸小
3. 没有办法很好表示层次丰富的图片信息

svg适合做图标

![76001330370](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760013303702.png)

![76001332700](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760013327001.png)



 ####  vite对静态文件的处理

![76001489992](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760014899925.png)

使用hash算法进行对文件内容的处理

配置

![76001538472](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760015384720.png)

配置文件名的hash

和对某个文件小于某个值的都换成base64



使用outDir来修改打包生成的文件









####  vite插件

vite会在不同的生命周期调用不同的插件以达到不同的目的

1. 生命周期：vite从开始执行到执行结束
2. webpack输出html的插件，在结束的时候。清除输出目录：在开始的时候：clean-webpack-pliugin
3. 中间件，插件





####  vite-aliases

插件学习

这个插件可以帮助我们自动生成别名：检测你当前目录下包括src在内的所有文件夹，并帮助我们生成别名

![76024055416](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760240554162.png)

直接使用，插件就像一个函数，vite会在不同时间运行





手写vite-aliases插件

手写的插件要在vite-alises执行前去修改vite0alises的文件

插件要导出一个函数



```javascript
const fs = require("fs")
const path = require("path")

function getTotalDir(dirFilesArr=[],basePath=""){
    const result = {
        dirs:[],
        files:[]
    }
    dirFilesArr.forEach(name={
        const currentFileStat = fs.statSync(path.resolve(__dirname,basePath+"/"+name));
        const isDirectory = currentFileStat.isDirectory();
        
        if(isDirectory){
        result.dirs.push(name)
    }else{
                        result =.files.push(name)
                        }
    })
    
    return result;
}

function getTotalSrcDir(){
    const result = fs.readdirSync(path.resolve(__dirname,"../src"))
    const diffResult = diffDirAndFile(result,"../src")
    const resolveAliasesObj = {};
    diffResult.dirs.forEach(dirName=>{
        const key = `@${dirName}`;
        const absPath = path.resolve(__dirname,"../src"+"/"+dirName)
        resolveAliasesObj[key] = absPath
    })
    
    return resolveAliasesObj
}

module.exports = ({keyName="@"})=>{
    return {
        config(config,env){
            //config:目前的一个配置对象
            //production:development server build yarn dev yarn build
            //env:mode:string,command:string
            //config函数可以返回一个对象,这个对象是部分的viteconfig配置,就是你想要改的那部分
            //const result = fs.readdiSync(path.resolve(__dirname,"../src"))
            const resolveAliasesObj = getTotalSrcDir()
            //读取这里面的模块
            return {
                resolve:{
                    alias:resolveAliasesObj
                }
            }
        }
    }
}
```

这样就手写了一个vite的插件





####  vite-plugin-html

webpack --> webpack-html-plugin/clean-webpack-plugin

但是vite内置了很多插件，减少了我们的心智负担

这个东西明显的就是![76037225056](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760372250564.png)

控制html

![76037227762](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760372277623.png)

手写插件

![76037302567](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760373025672.png)



####  mock数据

![76037346986](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760373469860.png)

使用vite-plugin-mock和modkjs

![76041739919](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760417399193.png)

在mock文件夹下的index。js文件

![76041741807](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760417418079.png)

配置项哪里要写viteMockServe()

多去看文档

vite会拦截这个服务器的请求,而不让他发到后端

```javascript
const fs = 

export default (options)=>{
    //拦截http请求
    //处理请求地址
    //本地开发的时候去拦截
    return {
        configureServer(server){
            server.middleware.use((req,res,next)=>{
                //req是请求,res是响应,next是放行给下一个中间件
if(req.url==="/api/users"){
const mockStat = fs.statSync("mock")
const insDIrectory = mockStat.isDirectory()
if(insDIrectory){
let children = fs.readdirSync("mock")
const result = require(path.resolve(process.cwd(),"mock/index.js"))
}
}else{
next()
}
            })
        }
    }
}
```

......很多很麻烦

![76043095429](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760430954293.png)

vite拿到他的内部插件，配置文件解析完后会调用configResolved()

configurePreviewServer()是对一些预览的处理





####  vite和ts的结合

vite对ts是天生支持的,ts提供了一种强类型锁定,让它不容易出错

ts的配置(主要是让ts直接输出到控制台,阻塞你的开发)

**使用vite-plugin-checker**

刚开始安装这个东西的,要使用配置文件取消检测node_module去跳过检测

![76043235404](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760432354049.png)

vite和ts结合是一定要用这个包的,也就是架构!

打包的时候检查ts语法规范:`tsc --noEmit && vite build`

这就是一个综合的构建方法

![76043262447](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760432624478.png)

设置编译出来的js格式

**三斜线指令**

/// <reference type="vite/client" />

用来给d.ts来写声明文件

![76043279748](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760432797488.png)

给环境变量写相关的配置



####  vite优化概述

- 开发时态的构建速度优化

  ![76043425208](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760434252083.png)‘

- js逻辑

  - 注意副作用的清除，特别是计时器，在频繁的挂载和卸载里面，计时器没有清除，下次再生成计时器，又会多一个计时器的线程

    ```javascript
    const [timer,setTimer] = useState(null)
    useEffect(()=>{
        setTimer(setTimer(()=>{}))
        return ()=>clearTimeout(timer)
    })
    ```

  - 这个写法上特别要注意浏览器的帧率是16.6ms刷新一次

    ![76043465207](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760434652077.png)

  - 防抖节流,自己不熟还是要用lodash库,人家写的性能好

  - 控制作用域

    ![76043476420](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760434764207.png)

    让arr.length不要去被多次访问,优雅

- css

  关注继承属性,能继承的就不用重复写

  避免过深的嵌套

- 构建优化

  rollup的优化










####  分包策略

浏览器缓存策略：浏览器会缓存名字不变的静态资源，所以要使用hash对文件名字进行替代

分包就是把不会常规更新的文件，进行单独的打包处理

![76051866572](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760518665725.png)

默认名字就是vendor‘





####  gzip压缩

文件资源太大了，把所有静态资源文件进行压缩，已达到减少体积的目的

![76051927222](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760519272225.png)

服务端会去读取gzip文件

![76052049476](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760520494760.png)

体积不大就不用压缩了，其他就压缩一下





####  动态导入

![76052074193](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760520741931.png)

类似的就是路由：根据不同的地址，展现不同的组件

这样写是读了全部的文件，会加载很多

![76052106301](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760521063019.png)

所以这里可以用一些动态加载

原理

![76052145243](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760521452432.png)

![76052148841](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760521488418.png)









####  cdn加速

![76057791953](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760577919534.png)

cdn链接有一个好处，它会去找你最近的cdn服务器，这样减少了一些不必要的时间损耗

![76057892016](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760578920167.png)

把lodash指向这个cdn服务器





####  vite处理前端跨域

1. 同源策略，只在浏览器发生
2. ![76057976982](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760579769825.png)

如果360不提示，那这个请求就会被拦截

![76057992508](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760579925089.png)

![76058043393](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760580433937.png)

在vite.config.ts进行一个配置

![76058052150](C:\Users\zxh\Desktop\前端\vite\vite.assets\1760580521504.png)
































