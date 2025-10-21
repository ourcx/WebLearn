# 一、Vue2篇

### 1. 关于生命周期

##### 	1.1 生命周期有哪些？发送请求在created还是mounted？

```
请求接口测试：https://fcm.52kfw.cn/index.php?_mall_id=1&r=api/default/district
```

Vue2.x系统自带有8个

```
beforeCreate
created
beforeMount
mounted
beforeUpdate
updated
beforeDestroy
destroyed
```

发送请求在created还是mounted？

```
这个问题具体要看项目和业务的情况了，因为组件的加载顺序是，父组件引入了子组件，那么先执行父的前3个生命周期，再执行子的前4个生命周期，那么如果我们的业务是父组件引入子组件，并且优先加载子组件的数据，那么在父组件中当前的请求要房mounted中，如果当前组件没有依赖关系那么放在哪个生命周期中请求都是可以的。
```

##### 	1.2 为什么发送请求不在beforeCreate里？beforeCreate和created有什么区别？

为什么发送请求不在beforeCreate里？

```
因为：如果请求是在methods封装好了，在beforeCreate调用的时候，beforeCreate阶段是拿不到methods里面的方法的（会报错了）。
```

beforeCreate和created有什么区别？

```
beforeCreate没有$data
created中有$data

created是可以拿到methods的方法的
beforeCreate拿不到methods的方法
```

##### 	1.3 在created中如何获取dom

```
1. 只要写异步代码，获取dom是在异步中获取的，就可以了。
	例如：setTimeout、请求、Promise.xxx()等等...
2. 使用vue系统内置的this.$nextTick
```

##### 	1.4 一旦进入组件会执行哪些生命周期？

```
beforeCreate
created
创建前后
beforeMount
mounted
挂载前后
```

##### 	1.5 第二次或者第N次进去组件会执行哪些生命周期？

如果加入了keep-alive，只会执行一个生命周期

就是activated

如果没有加入keep-alive

![75083767346](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1750837673462.png)

会继续走这几个

##### 	1.6 父组件引入子组件，那么生命周期执行的顺序是？

像axios，基本都是最后执行的，这些微任务的执行是和宏任务不一样的  

```
父：beforeCreate、created、beforeMount
子：beforeCreate、created、beforeMount、mounted
...
父：mounted
先执行父组件的东西，在到子，子的完成后，就到父的挂载了
```

##### 	1.7 加入keep-alive会执行哪些生命周期？

```
keep-alive是用来缓存当前组件的东西
会新增两个生命周期
activated
deactivated
```

如果组件加入了keep-alive，第一次进入会执行5个生命周期

![75083758966](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1750837589661.png)

##### 1.8 你在什么情况下用过哪些生命周期？说一说生命周期使用场景

```
created 单组件请求的情况下,写到这里
同步获取dom，就在mounted，有子组件后父组件请求等情况下，就要在特定时候获取
activated判断id是否相同,相同不发送请求,而是使用缓存
destroyed记录播放到的时间,下次初始化就从这个地方开始播放
```

### 2. 关于组件

##### 	2.1 组件传值（通信）的方式

```
父传后代
1. 使用props,在组件哪里绑定,单向的
子组件使用props接收
props:{
    xxx:{
        type:string
        default:""
    }
}
做不到直接传递给孙子,传递比较麻烦
2. =直接拿到
通过this.$parent.xxx使用父组件的数据
这种方式,子组件可以直接修改父组件的数据
3.依赖注入
优势:父组件可以直接向某个后代组件传值(不止一层一层传递)
父组件provied
子组件inject
```

```
后代传父
1.子组件传值给父组件
子 this.$emit("change",value)
父 <List @change='Btn'></List>
   Btn(val){
       console.log(val)
   }
   或者使用
   this.$on("change",val=>{
       console.log(val)
   })
2.使用this.$children拿到子组件数据数组 
3.使用<List ref='child'></List>
this.$refs.child
```

```
平辈之间传递数据

```

##### 	2.2 父组件直接修改子组件的值

```
使用<List ref='child'></List>
this.$refs.child
```

##### 	2.3 子组件直接修改父组件的值

```
子组件可以使用this.$parent.xxx去修改
新建一个文件来进行周转
或者直接用vuex或者pinia
```

##### 	2.4 如何找到父组件

```
this.$parent
点击
```

##### 	2.5 如何找到根组件

```
this.$root
这些拿到节点的多在插件里面使用
```

##### 	2.6 keep-alive

缓存组件的，提升性能

##### 	2.7 slot

```
什么情况下会用到插槽
减少代码的复用,不用重新去写一个组件
1.匿名插槽
2.具名插槽
3.作用域插槽
```

##### 	2.8 如何封装组件

```
有没有制作公共组件,有没有加上插槽,组件通信等

```

![75098812946](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1750988129461.png)

##### 2.9provide/jnject

```
依赖注入
```

### 3. 关于Vuex

##### 	3.1 Vuex有哪些属性

##### 	3.2 Vuex使用state值

##### 	3.3 Vuex的getters值修改

##### 	3.4 Vuex的mutations和actions区别

##### 	3.5 Vuex持久化存储

```
vuex本身不是持久化储存的数据,vuex只是一个状态管理库=>是放全局属性的地方
想实现持久化储存只能自己写和使用插件
```

### 4. 关于路由

##### 	4.1 路由的模式和区别

##### 	4.2 子路由和动态路由

##### 	4.3 路由传值

##### 	4.4 导航故障

![75145834646](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751458346464.png)

vue2同页面跳转同页面的时候,会发生错误

在push方法哪里处理一下

![75145841704](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751458417043.png)

调到原型链上面进行相关的处理

##### 	4.5 $router和$route区别

```
router是一个大类,是整个路由的方法对象
route包含当前路由对象
```

##### 	4.6 导航守卫

![75145850404](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751458504045.png)

 全局守卫用的比较少,一般用于验证身份

![75145887956](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751458879562.png)

路由独享的路由守卫,需要next放行

组件内路由守卫需要调用(在组件里面才能使用)

### 5. 关于API

##### 	5.1 $set

##### 	5.2 $nextTick

```
功能是获取更新后的dom的
原理:一个异步的行为
```

![75145769243](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751457692435.png)

大概就是使用dom拿到了一个回调

##### 	5.3 $refs

##### 	5.4 $el

##### 	5.5 $data

##### 	5.6 $children

##### 	5.7 $parent

##### 	5.8 $root

##### 	5.9 data定义数据

##### 	5.10 computed计算属性

##### 	5.11 watch

##### 	5.12 methods和computed区别

```
computed是计算属性,vue3基本不使用methods

1.computed是有缓存的,后面会走缓存的结果,methods没有缓存机制
2.computed适合根据已有数据衍生新的数据过滤和转化,提高性能
3.methods每次调用都会执行一次
```

### 6. 关于指令

##### 	6.1 如何自定义指令

![75145941526](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751459415268.png)

然后可以在组件根据v-demo来进行使用

![75145951437](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751459514372.png)

局部的自定义事件



##### 	6.2 vue单项绑定

v-bind

##### 	6.3 v-if和v-for优先级

![75145717441](D:\BaiduNetdiskDownload\前端程序员来了7.9面试题笔记\面试题\1-28道面试题\01 生命周期有哪些？发送请求在created还是mounted？\vue面试题.assets\1751457174410.png)

vue3里面v-if的优先级比较大

### 7. 关于原理

##### 	7.1 $nextTick原理

##### 	7.2 双向绑定原理

### 8. axios二次封装

# 二、Vue3篇

### 1. Vue2和Vue3区别

