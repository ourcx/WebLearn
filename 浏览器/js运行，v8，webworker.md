####  js引擎和运行机制

![75674227367](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756742273674.png)

js是混合了编译和解释的一种实现，临时编译型

![75679072936](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756790729361.png)

v8引擎是使用最多的js解析引擎

浏览器：js引擎，渲染引擎blink



####  深入v8引擎

![75679094305](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756790943050.png)

字符串-》编译/解释AST-》字节码 -》机器码 -》执行

![75679106841](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791068415.png)

![75679116722](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791167229.png)

![75679140858](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791408585.png)

反优化

![75679145942](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791459421.png)

打破热点函数之后要做的

可以进行一些回退的操作

所以要使用typescript，利于v8的优化





####  代码储存

 ![75679169046](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791690469.png)

var和let和const

 ![75679183257](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756791832572.png)

![75679202383](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792023831.png)

var会放到变量环境，let会被放到词法环境

![75679205869](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792058699.png)





#### 处理机制

 js是单线程的，但是有异步任务执行机制

![75679239386](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792393865.png)

![75679256929](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792569295.png)

![75679258034](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792580345.png)

执行一个宏任务之后，立刻清理所有微任务，然后再进行ui渲染，开启下一个宏任务

![75679267722](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756792677224.png)



####  浏览器和node的事件循环存在差别

![75685755468](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857554685.png)

![75685757980](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857579807.png)

####  内存垃圾回收



![75685761116](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857611162.png)

标记清除算法，只清除不整理

老版本是标记清除整理的方法

![75685766884](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857668845.png)



####  webworker也是一个性能优化的点

![75685774914](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857749143.png)

就是创建后台线程，用post message的方式去交换信息

![75685781156](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857811560.png)

![75685781839](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857818398.png)

主线程监听worker线程

> 使用worker线程,小程序也能使用

![75685793257](C:\Users\zxh\Desktop\前端\浏览器\js运行，v8，webworker.assets\1756857932574.png)



消耗性能的都能丢进去

