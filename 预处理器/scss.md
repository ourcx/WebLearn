#  less的编译

less的复代码太多了，包含颜色值，容器大小等，那我们是否可以使用js变量声明的方式解决这种问题，原本的css不支持，所以要使用less

less是一款css预处理语言，支持变量，混合，函数，嵌套，循环等特点，在css的基础上加了很多大小，比css丰满，为css增加了一些特性，将css作为目标生成文件，然后开发者就只需要使用这种语言进行编码工作

注意：

- 浏览器不认识less，写的less代码，需要编程成css
- 项目中使用预处理语言



**编译方法**

- 安装node
- 安装less     `npm i -g less`
- 编译less  

1. 方法1

less

```less
@width:980px;
@height:width+100px;
@color:skyblue;
//定义变量v  
#header{
    width:@width;
    height:@height;
    background-color:@color;
}
```

使用命令`lessc style.less style.css` 

css

```css
#header{
    width:980px;
    height:1080px;
    background-color:skyblue;
}
```

2. 方法2

   页面直接引入less.js

   通知vscode open with live server

   ![74908380253](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749083802535.png)

3. 方法3

   ![74908384187](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749083841872.png)

   开发的时候使用这些插件

   ![74908386336](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749083863369.png)

   写完less代码后,ctrl+s自动将less转化为css

4. 方法4

   在项目中,使用工程化的打包工具

   webpack



####  变量

less允许使用@符号定义变量，变量分配使用冒号：完成，less的变量声明格式为@变量名：变量值

![74908691812](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749086918123.png)

选择器的名字也可以使用变量，url地址也可以使用变量，变量可以先使用在声明

 

####  混合

![74908740476](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749087404769.png)

快速复用代码

如果不想编译.box1()

![74908743901](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749087439018.png)

混合也可以给他进行传参的操作,这些参数是混合时传递给选择器块的变量

![74908771258](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749087712586.png)



####  嵌套

![74937040217](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749370402174.png)

一个嵌套的写法

可以一直嵌套下去

![74937061356](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749370613566.png)

&的意思表示当前的父层级,比如这里的&就表示一个li







####  运算

基本运算符+ - * / 可以对任何数字 变量或者颜色进行运算,如果可以的话,算术运算会在加减或者比较之前会进行单位换算,计算的结果以最左侧操作数的单位类型为准,如果单位换算无效或者失去意义,则忽略单位

![74937203251](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749372032514.png)

实列





####  函数

less内置了很多用于转化颜色和处理字符串,算术运算的函数,这些函数在less函数手册有记录

![74937243532](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749372435320.png)





####  作用域

和js是一样的,先查找本层,再去父层查找

![74937254932](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749372549322.png)

####  注释和导入

注释和js是一样的,按ctrl加/就可以了

导入是

使用

> @import '相关的路径'





#  sass

也是和less很像的一种预编译语言

创建文件的文件后缀是scss

1. 编译sass

   使用`live sasscompiler `插件或者使用webpack进行编译

   ![74937332507](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749373325074.png)

   点哪个底部的按钮可以插件编译

   ![74937334926](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749373349267.png)

2. 变量

   sass使用$来标识变量

   ![74937376860](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749373768607.png)

3. 嵌套

   也可以像less一样使用嵌套去进行

4. 导入sass文件

   和less一样,变量冲突就选择最近的变量

5. 注释

   和less一样

6. 混合

   需要提前定义混合

   ![74937422085](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749374220852.png)

   使用@mixin混合

   使用@include导入混合,并且支持传参

7. 选择器继承

   要用@extend进行继承

   ![74937429428](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749374294285.png)

8. 运算

   也是相关的加减乘除都可以,和less很像

   指令

   ![74937435048](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749374350488.png)

   ![74937435798](C:\Users\zxh\Desktop\前端\预处理器\scss.assets\1749374357985.png)

   for和if可以去使用

   ​