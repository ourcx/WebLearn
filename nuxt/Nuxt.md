####  同构渲染ssr

![75550679277](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755506792775.png)

服务端渲染，客户端渲染，同构渲染（一部分服务端，一部分客户端）

![75550811012](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755508110120.png)

![75550836328](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755508363289.png)

这时候的请求是一步步来的，按需使用

 ![75550871013](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755508710134.png)

首屏渲染速度需要优化，搜索引擎第一次请求网页，分析网页内容，给别人搜索，客户端渲染时，搜索引擎第一次拿到的是空的页面，不利于SEO，网站不好被搜索引擎记录

**相应的框架就是vue react等**

![75550904108](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755509041087.png)

第一次请求给完整的页面，后面的请求都是用ajax进行渲染

什么时候使用同步渲染：

项目对外，需要被搜索，且首屏渲染等

的时候考虑使用同步渲染



####  开始

直接下载或者克隆

![75552518356](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755525183566.png)

npm run dev -- -o

是开发服务器的启动命令



####  基本路由

![75552563939](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755525639395.png)

根组件哪里挖了一个路由出口：nuxt-page组件

 ![75552578543](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755525785437.png)

简单的路由设置

在路由目录下创建index.vue可以当成一个默认路由,在访问文件夹的时候使用

**父子组件**

![75552635979](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755526359799.png)

![75552637248](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755526372483.png)

  



####  路由跳转

1. 使用标签跳转

![75557708036](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755577080365.png)

```vue
<template>
  <div>
    <h1>我是根组件</h1>
    <nuxt-link to="/users">用户-列表</nuxt-link>
    <nuxt-link to="/users/create-or-edit">用户-添加或更新</nuxt-link>
    <nuxt-link to="/about">关于</nuxt-link>
    <nuxt-link to="/roles/admin">角色-管理员</nuxt-link>
    <nuxt-link to="/roles/normal">角色-普通用户</nuxt-link>
    
    <!-- 
  		nuxt-link 会被编译成 a 标签, 但是不推荐直接使用 a 标签
    	直接用 a 标签，点击会刷新跳转，相当于一次全新的向服务端发起的请求
    	使用 nuxt-link 是浏览器端本地切换页面，即SPA
  	-->
    <nuxt-page />
  </div>
</template>

<script setup>
</script>

<style>
  /* nuxt-link 会被编译成 a 标签 */
	a { margin: 20px; }
</style>
```

推荐组件有一个根标签

2. 自定义路由参数

   ![75557965227](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755579652276.png)

   query也是这样取出来

3. 自定义插件

   ![75558081894](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755580818941.png)

路由匹配失败等错误, 可以在项目根目录下创建 error.vue

中间件

![75558349011](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755583490110.png)

![75558350228](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755583502284.png)



####  使用组件

![75558132420](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755581324201.png)

三种使用方法:

在你隔壁的,需要你引入

在components下的,不用引入,但要区分文件夹和非文件夹,所以要用大小驼峰,有文件夹的使用驼峰





####  布局处理

![75558357197](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755583571979.png)

![75558355683](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755583556834.png)

页面默认布局



####  SEO的配置

在nuxt.config.ts进行相关的配置

![75577216594](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755772165946.png)

也可以在vue文件进行设置

![75577310168](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755773101685.png)



还有其他设置

![75577329967](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755773299676.png)

  或者用动态标题

![75577371923](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755773719235.png)

动态改变

> 上面的一些配置可能有冲突,注意删除\



####  静态资源的访存

> public目录:存放的内容相当于服务器的根目录去访问
>
> assets:构建工具会去处理的内容

1. 访问public

   ![75577408923](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755774089236.png)

2. 在assets

   ![75577417514](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755774175143.png)

   要用构建工具就用这个

3. 使用差别

   ![75577486830](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755774868307.png)







####  配置环境变量

![75583615487](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755836154878.png)

定义环境变量

![75583616837](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755836168377.png)

这样拿到相关的值

 .env文件的值会覆盖你的nuxt.config.ts文件

  ![75583714664](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755837146644.png)

在pbilic里面执行的值浏览器可以看到,通过isServer在浏览器和服务器不同的值来判断当前的环境

也可以在app.config.ts里面设置环境变量

![75583725868](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755837258684.png)

然后取值

![75583726699](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755837266997.png)





####  获取数据

![75583821682](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755838216824.png)

![75583843519](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755838435195.png)

nuxt框架提供了其他方式去发请求

这个是nuxt3调用http调用的首选方式

![75584017729](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755840177296.png)

只是下面那个多了一个请求的说明

 ![75584137262](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755841372621.png)

使用解构语法可以解构这几个值

**带lazy的请求**

不会阻塞在一起，会先返回

1. 使用useFetch请求数据

![75584186443](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755841864436.png)

![75584245499](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755842454994.png)



token可以在headers里面设置，也能在拦截器option里面的authorization里面设置



####  nuxt也能编写后端的代码

服务端

> 在server/api/aaaa.ts

请求对应的地址就可以了,一般在NUxt开发里面,是请求相关的后端服务器

![75595041566](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755950415662.png)

写出~/server/routers/文件下面,就不用写api了

也可以写成server/api/xxx/ss.ts

或者是变化的值,路由传参

![75595085104](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755950851040.png)

![75595099952](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755950999520.png)

实际拿到的age

前端

![75595034024](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755950340242.png)



请求方式

通过更改文件名进行请求方式的改变

![75595548513](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755955485139.png)

 拿到请求数据的body

![75595552681](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755955526814.png)



写一个api/[...].ts可以让前端匹配匹配不到的路由

**中间件**

![75595759320](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755957593201.png)

在nuxt服务端也能调用redis等数据库来传输资源





####  状态管理

![75595805196](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1755958051969.png)

写在composables/states.ts下面的,自带抓状态管理

在nuxt也能使用pinia做状态管理



####  自动导入

很多ref和computed和一些工具函数在nuxt可以自动导入

在配置文件里面可以关掉







####  构建nuxt

  直接build

然后node允许index。mjs页面

**node  .output/server/index.mjs**

端口和地址可以通过 process.env 配置, 但没必要, 不要纠结



或者使用PM2，一直运行

1. 安装 pm2

npm  i  pm2  -g      // 可能需要管理员权限

pm2  -v

![75601212132](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756012121329.png)



####  路由渲染策略

```javascript
export default defineNuxtConfig({
  routeRules: {
    '/blog/**': { swr: true },  // Static page generated on-demand, revalidates in background
    '/articles/**': { static: true },   // Static page generated on-demand once
    '/_nuxt/**': { headers: { 'cache-control': 's-maxage=0' } },   // 设置响应头信息
    '/admin/**': { ssr: false },  // Render these routes with SPA
    // Add cors headers
    '/api/v1/**': { cors: true },
    // Add redirect headers
    '/old-page': { redirect: '/new-page' },
    '/old-page2': { redirect: { to: '/new-page', statusCode: 302 } }
  }
})

```

简短的配置

```javascript
export default defineNuxtConfig({
  routes: {
    '/': { prerender: true },         // 每一次构建时，都重新预渲染页面 (透过 Builder),常用的页面数据不变化的可以变成静态渲染
    '/blog/*': { static: true },      // 接收到一个请求时，页面依照需求重新渲染页面 (透过 Lambda)
    '/stats/*': { swr: '10 min' },    // 接收到一个请求时，10 分钟的快取缓存过期后，将会再次的重新取得数据进行渲染 (透过 Lambda)
    '/admin/*': { ssr: false },       // 仅在客户端渲染
    '/react/*': { redirect: '/vue' }, // 路由重定向,没找到相关路由的时候跳转
    '/_nuxt/**': { headers: { 'cache-control': 's-maxage=0' } },   // 设置响应头信息
  }
})
```

在nuxt.config.js里面写这些,这里写了客户端渲染admin文件夹的内容,服务端不会渲染里面的内容,而是让客户端去请求

不同路径有不同的策略

**取消服务端渲染**

![75601304395](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756013043952.png)





####  SSG-预渲染

在服务器构建的时候使用的东西，加快页面拼接的速度

![75601362103](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756013621034.png)

执行generate生成静态页面，先生成好静态页面，直接在请求的时候返回

并且默认不生产动态路由

![75601374936](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756013749366.png)

要去nuxt.config.ts定义才能渲染动态路由



---

CSR、SSR和SSG各有优缺点，适用于不同的场景和需求。在选择使用哪种技术时，需要根据项目的具体需求来权衡利弊。

例如，对于需要丰富交互效果和实时数据的场景，可以选择CSR；对于需要优化首屏加载速度和SEO效果的场景，可以选择SSR；而对于内容更新不频繁、对性能要求高的场景，可以选择SSG。同时，也可以结合使用多种技术来实现更好的用户体验和性能优化。

----





导入组件库

![75601414960](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756014149607.png)

**![75601417670](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756014176701.png)**

两种导入方式

![75601421346](C:\Users\zxh\Desktop\前端\nuxt\Nuxt.assets\1756014213464.png)

组件懒加载方式

