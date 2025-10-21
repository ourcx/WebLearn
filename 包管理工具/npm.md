#  包管理工具

 ![75404433318](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754044333180.png)

react默认使用yarn来管理，cnpm是国内常用的

pnpm有一些操作系统的概念





####  代码共享方案

  ![75405901733](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754059017335.png)

 ![75405930814](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754059308144.png)

使用包仓库来完成代码的存放，再进行分享

 

####  npm

 ![75406090506](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754060905063.png)

看包的官网和github有没有

![75406641611](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754066416111.png)





####  npm的配置文件package.json

记录一些项目的信息

 ![75409652339](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754096523392.png)

在真实开发里，我们要初始化自己的package.json

`npm init`初始化

使用脚手架可以省略一下

![75411390939](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754113909394.png)





####  项目依赖

常见属性

1. scripts脚本

   ![75411402932](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754114029326.png)

2. ![75413021039](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754130210394.png)

   有些脚本不用加run，就是上面加的一些东西

3. ![75413083690](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754130836908.png)

   依赖和开发依赖（生产环境会被去除的包）








####  版本管理

  ![75413230627](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754132306270.png)

不向前兼容的时候就要大版本更新的

小功能就更新次版本号就可以了

修订一般是修订BUG的

![75413261148](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754132611487.png)

![75413278902](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754132789022.png)

是对浏览器的兼容和指定一些版本控制

![75413281307](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754132813074.png)

一些的项目的package.json





####  npm安装包的细节补充

 -g是全局安装，没-g不是全局安装

![75414543352](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754145433527.png)

![75414750093](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754147500938.png)

全局安装会报一个警告

![75414753824](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754147538248.png)

就是-g可能会过期

局部安装是在目录下生产一个node_modules文件夹

![75414793711](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754147937110.png)





​      



####  package-lock.json的作用

![75414830086](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754148300864.png)

找缓存和去远程库拉取

package-lock。json这个文件，它会确实我们要的版本，保证版本一样

 ![75414948453](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754149484533.png)

 



####  install的原理和解析

![75464526165](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754645261651.png)



####  npm 其他指令的作用

![75464567095](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754645670957.png)





####  yarn

![75464591869](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754645918697.png)

但在npm5后npm的缺陷被升级修复了

使用yarn需要使用`npm install yarn直接安装这些东西`

 ![75464620853](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754646208538.png)





####  关于cnpm和淘宝镜像

![75464650802](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754646508026.png)

镜像服务器每十分钟就会从npm的仓库拉取新的镜像

 ![75464782103](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754647821032.png)

不想使用镜像,就是要cnpm吧

 



####  npx工具

![75464906848](C:\Users\zxh\Desktop\前端\包管理工具\npm.assets\1754649068489.png)

