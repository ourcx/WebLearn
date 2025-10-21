#  ts基础知识 

ts是动态类型的语言

- 运行时才做检查的语言
- js,py,ruby

静态类型的语言

- 编译阶段就做类型检查
- c/c++,c#,java



ts优势

- 程序更容易理解
- 效率高
- 错误少
- 包容性好、



####  细说

tsc是ts的编译器，将ts文件编译为js文件

####  原始值

![74580078178](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745800781785.png)

原始类型：boolean、null、undefined、number、bigint、string、symbol、object

尽量避免使用any



####  数组和元组

```typescript
let arrOfNumbers:number[]=[1,2,3]
arrOfNumbers.push(123)
//push进去的东西也必须符合number


//元组
let user : [string,number]=["ssss",20]
user.push(123)
uaer.push("sdsd")
```





####  接口

1. 对对象的形状进行描述(duck typing鸭子类型)


```typescript
interface Person{
    readonly id ：numeber
    //设置id为只读属性 
    name:string;
    age?:number
    //?表示age是可以不要的或者要的
}


let viking:Person = {
    name:"xxxx",
    age:20,
}
```



####  函数

```typescript
function add(x:number,y:number):number{
    return x+y
}

const add =(x:number,y:number):number=>{
    return x+y
}
```

类型推论,当你没有声明类型的时候,tsc会自动推出一个类型来使用

联合类型:

- > let numberOrstring:number|string
  >
  > 可以给这个变量赋值数字或者字符串

类型断言

- > 给联合类型使用
  >
  > 把其确定下来
  >
  > const n = n as number
  >
  > 或者转换类型
  >
  > typeof n == "string"



####  枚举

普通枚举

 ![74582943870](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745829438702.png)

字符串枚举:常量枚举,可以提升性能或者计算值枚举

 ![74582955569](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745829555696.png)

枚举类似一个对象





####  泛型

泛型类似一个占位符,为了解决一些东西约束不是很好的问题

```typescript
function echo <T>(arg:T):T{
   return arg
}
const result = echo(123)
//自动推论echo的返回值


//多个泛型的执行
function swap<T,U>(tuple:[T,U]):[U,T]{
    return [tuple[1],tuple[0]]
}
const result = swap(["string",123])


interface IWithLength{
    length:number
}
//属性约束

function echoWithArr<T exthends IWithLength>(arg:T):[]{
    console.log(arg.length)
    return arg
}
                     
//强制约束arg必须有长度这个属性，没有extence的（arr.T[]）只能传数组进去,很不方便,所以要进一步解决,使用extence

const arrs = echoWithArr([1,2,3])
```

 

####  泛型在类和接口中的使用

1. 类:

   ![74598872061](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745988720614.png)

   初始化类的时候传入类型,然后就可以去使用了

2. 接口:

    ![74598893857](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745988938574.png)

   ​



####  补充

1. 类型别名

   ![74598953000](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745989530005.png)

   重新给type类型声明了一个名字什么的

   ​


2. 字面量:const str:"xxxx"="xxxx"

   str只能等于xxxx,常量值不改变

3. 交叉类型:

   ![74598971827](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745989718275.png)

   实现交叉的时候用type,当实现extence,实现类的时候用interface

   ​



####  声明文件

建立一个xxx.d.ts文件

在这个文件写声明文件,写很多类型声明,没有任何实现代码

只有类型:比如interface,function,class

当类型文件放置在特定的文件夹里面,就不用引入和打开.d.ts文件来持久化

![74599126792](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745991267929.png)

使用的时候就是

calculator.plus([1,2,3])

并且导出成一个模块,要引入再来使用,第三方的,把它放node_modules的@types目录下可以按照导入第三方的方式进行导入



####  内置类型

就是js的一些内置类型



####  内置文件

tsconfig.json

![74599523337](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1745995233373.png)



####  文件结构和插件

1. tsconfig.json是浏览器环境文件，tsconfig.node.js是node环境文件

   vite.config.ts是一个vite的插件管理文件

2. index.html是入口文件

3. vite-env.d.ts是对文件类型的定义文件

插件

1. volar语法高亮和智能提示什么的 
2. 浏览器要用vue-tool



####  ESLint

代码规范问题,提高工作效率，使用espree将js代码解析成抽象语法树ast，通过ast来分析代码

rules：

一个rule有三个等级，0代表关闭，1代表warning输出警告而不是错误，2代表错误，会抛出错误

安装：

npm i eslint --save-dev

npm init @eslint/config

可以控制的：

加不加分号，加不加空格

可以使用别人的规则

extends这个字段控制了外部规则

比如默认的一些规则，可以简化配置

**配合vue去使用：**

![74628067061](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746280670619.png)

![74628121720](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746281217202.png)





####  vue3特性依赖注入

- Provide/inject

  在一个经典的vue应用中，数据是通过props属性由上向下（由父及子）的结构进行传递的，这种特性称之为prop Deilling

  Provide/inject的推出就是为了解决这个问题，它提供了一种组件之间共享此类值的方式，而不必通过组件树的每个层级显式地传递props

  目的是为了共享那些被认为对于一个组件树是全局的数据

  ```javascript
  //父组件

  provide(langKey,"值")
    //子组件

    inject("建")

    //使用东西去接受

    //当东西多的时候,容易重名,导致覆盖的问题出现

    //使用Symbol数据类型防止这种情况出现

    //每个Symbol()返回的symbol值都是唯一的,一个symbol值能作为对象属性的标识符

    在其他文件调用

    export const langKey = Symbol()

    回来导入和使用

    使用很多这个Symbol

    provide(langKey,"值")

    inject(langKey)
    
  ```

```

```

![74652908203](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746529082032.png)

  配置全局属性和全局组件

  ![74652921736](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746529217366.png)

  配置全局数据可传递

  在哪里都能用

- 同时.vue3可以有两个应用实例,和vue2不同



#  开始书写feiUI组件库



使用create-vue进行相应的封装vite

技术栈:vite+Vue3+TS+Eslint

####  button

![74653000240](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746530002405.png)

![74653020374](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746530203741.png)

精简的项目结构

![74653088699](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746530886993.png)

要完成的属性列表

![74653095006](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746530950061.png)

![74653138985](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746531389855.png)

封装一个类似button的原生组件的时候,也要考虑它是否适合原生属性的使用,这个是一定要注意的点

#####  css预处理器

变量、nesting（嵌套）、循环、类似函数（mixin）

使用sass或者less或者stylus或者postcss



####  色彩

- 系统色彩

基础色板

![74659976086](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746599760866.png)

中性色板

![74659983682](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746599836821.png)

- 产品色板

  品牌色（主要的）

  使用#a61b29（苋菜红为色阶）

  ![74659992923](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746599929235.png)

  功能色（明确的信息和状态）

  ![74659994128](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746599941282.png)

  ​

![74660088285](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746600882850.png)

![74660089435](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1746600894350.png)

我的主色

#b51e1e
#c03e3e
#ca5e5e
#d57e7e
#df9f9f
#eabfbf
#f4dfdf

辅助色

#96C24E
#A5CB67
#B4D381
#C3DC9A
#D2E5B3
#E1EECC
#F0F6E6

芽绿

#FED71A
#FEDD3B
#FEE25B
#FEE87C
#FFEE9D
#FFF4BE
#FFF9DE

佛手黄

#EFAFAD
#F1BAB9
#F4C6C4
#F6D1D0
#F8DDDC
#FAE8E8
#FDF4F3

无花果红

#BBB5AC
#C5C0B8
#CECAC4
#D8D5D0
#E2DFDB
#ECEAE7
#F5F4F3

铅灰



####  添加系统默认css样式

normalize.css

- 保护有用的浏览器默认样式
- 一般化的样式
- 修复浏览器自身的bug
- 优化css可用性

element-plus提供了更简单的，更贴合组件库的默认样式

postCSS插件-postCSS nested



####  使用postcss生成对应的颜色值

![74903529665](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749035296650.png)









####  总结button组件

流程：

- 确定props

- 确定组件的展示方式

- 确定事件

  ​

初始化项目以及项目结构

create-vue

项目结构



**组件编码**

* props的定义方式
  - 普遍的对象方式
  - ts类型方法
* 原生属性
* defineExpose定义实例导出
* 是否支持原生的事件
* 遗传

![74903680077](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749036800774.png)

传递的属性会直接被传到dom身上

**样式解决方案**

- 选择Postcss作为预处理器
- 了解色彩系统
- 使用css变量添加颜色系统
- 添加button样式
  - 善用变量覆盖
  - 使用postcss插件
- 使用postcss动态生成主题颜色

**你的yapostcss没有配置好，md**





###  collapse组件

就是做折叠面板或者是手风琴面板

![74903757781](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749037577813.png)

方案：

- 传入数组，进行遍历

- 语义化展示，甚至title可以用插槽实现

  ![74903777087](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749037770879.png)

使用语义化的方式进行数据的展示

那就要确定我能传什么属性进去了

确定事件：1. 改变事件

2. 更新事件





使用一个数组去维护collapse打开的状态

![74937626748](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749376267489.png)

![74937627680](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1749376276805.png)

使用v-model去动态双向添加元素，来操控collapse

还又手风琴等





- **动画效果**

  使用transation组件完成相关的动画效果

  ![75031413505](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750314135055.png)

  使用js方式的动画效果

  事件钩子，操作更精细

  ![75032898618](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750328986180.png)

  ![75032899557](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750328995579.png)





border-bottom和transation发生冲突

![75032902928](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750329029287.png)

添加一个父元素就行了

但是又有另外一个问题，里面的东西会立刻显示出来





总结：动画效果的实现方式，使用vue内置的Trasation和js的动画钩子进行配置



**Collapse总结**

- 类似列表结构组件，两种常用的实现方法

- ![75038579917](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750385799171.png)

   ![75038595239](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750385952397.png)

  语义化展示中，父子组件沟通的方法：

  - 数据状态以及处理放在父组件
  - 使用Provide/inject传递给子组件
  - 父组件为数据中心



- v-model的实现：

  - 是属性modelValue
  - 事件update：modelValue的语法糖
  - 特别注意属性赋值给内部响应式对象的更新问题（使用watch）

- 使用内置的Transition实现动画效果

  不提供动画，但是有classes表示动画和js的钩子自定义高级动画

  ​








####  图标库的发展

- 雪碧图CSS script

  一个图片拼接很多图片

  ![75038699326](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750386993261.png)

  ​

- Font Icon

  ![75038717443](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750387174437.png)

  ​

- inline SVG

  是完全可控的，font icon只能控制字符相关的属性

  font icon需要下载的字体文件大，除非自己切割

  font icon会有问题

- 使用现成的图标库

  ![75038727316](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750387273165.png)

  Fontawesome是这次要使用的开源图标库

  ![75038751650](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750387516503.png)

  第一个是核心库,第二个是对应的图标库(免费和付费),第三个是基于vue3的包装

npm i --save @fortawesome/fontawesome-svg-core

npm i --save @fortawesome/free-solid-svg-icons

npm i  --save @fortawesome/vue-fontawesome@latest-3







####  icon二次编码

完成相关类型的引入后,就可以根据自己的需求引入其他内容了

![75038964821](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750389648215.png)





**总结**

- 第一步：支持组件的原始属性
  - inheritAttrs：false不继承属性,让透传失效,让对应属性添加到子属性
  - 使用$props传递属性
  - 要注意不继承以后一些默认属性失效的问题
- 第二步
  - 我们添加的type/color属性
  - 过滤传递的属性lodash omit





**作业**

开发alert栏

![75039495981](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750394959815.png)

![75039499120](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750394991209.png)

 对应的属性有

![75039503221](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750395032218.png)

![75039507328](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750395073287.png)





####  为什么要有测试

- 出现的问题

  - 新建一堆不同属性的组件，肉眼观测
  - 更新代码需要手动再进行测试

- 测试再国内严重忽视

  - 没有时间
  - 需求一直修改
  - 不知道怎么测试

- 测试的优点

  - 自动化完成流程，保证代码的运行结果
  - 更早发现BUG
  - 重构和升级更加容易和可靠

- 测试金字塔

  ![75049227417](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750492274179.png)

  Unit单元测试

  Service部分测试

  UI模拟真实场景的测试

  ![75049243339](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750492433397.png)

  手动测试到下面的测试


####  在项目中添加测试

两种测试工具

- 通用测试框架
- 针对某种库的测试库



#####  通用测试框架

- Mocha
- jest
- Vitest

#####  Vitest的优点

- 基于vite的测试框架
- 兼容Jest
- 对ESM，TS，JSX特攻
- 智能

![75049314757](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750493147572.png)



 简单使用一个测试功能

![75049428666](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750494286663.png)





####  基于vue的测试工具

  ![75049511083](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750495110830.png)

![75049536076](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750495360761.png)

![75049616194](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750496161945.png)

组件测试

选择相对vue的测试框架，学习mount写法，将对应组件挂载上，配置文件的写法







####  vue-test-utils基本用法

在xxx-text-ts里面使用

 



####  使用Vnode以及Render Function

![75056219940](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750562199402.png)

![75056362846](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750563628461.png)

优点：

- 提供效率更高，计算需要的最小化操作，并完成更新
- 更方便操作dom

![75056439099](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750564390994.png)、

输出虚拟dom

![75056459178](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750564591786.png)

 



####  JSX/TSX

![75056500133](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750565001330.png)

![75056513636](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750565136369.png)

在tsx文件里操作







####  事件测试方法

目的:测试collapse的change事件

1. 传入事件的回调函数

   ![75065962798](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750659627980.png)

   创建回调

   ![75065963826](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750659638264.png)

   观测触发

2. 普通方式mount组件




####  测试用例

![75066090940](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750660909408.png)

在describe外部声明变量，在beforeAll赋值，然后进行操作

一次性注册实例







####  总结

![75066158535](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750661585358.png)

![75066172511](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750661725113.png)

![75066175991](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750661759916.png)

![75066180518](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750661805189.png)



####  vitest钩子函数

![75066183644](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750661836441.png)

描述生命周期

在用例开始前，所有用例开始前，还有两个事件修饰符







####  Tooltip组件需求分析

功能分析：

- 触发区
- 展示区

触发方式

- hover
- 点击
- 手动

重点是触发区，发生特定事件的时候，展示区的展示与隐藏

![75066248688](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750662486886.png)

属性

![75066250568](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750662505681.png)

![75066252212](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750662522121.png)

![75066252752](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750662527528.png)

实例为手动方式进行服务

这是一个通用组件

![75066309437](C:\Users\zxh\Desktop\前端\组件库\UI.assets\1750663094370.png)

浮层





####  安装popper.js

浮层显示库

`npm i @popperjs/core --save`

完成相关的安装





####  Tooltip开发计划

- 最基本的实现0 
- 支持click/hover两种触发
- 支持clickoutside的时候隐藏
- 支持手动触发
- 支持poper参数
- 支持延迟显示
- 样式
- 支持多个浮动窗口
- 支持渲染HTML
- 支持设置样式

![75074397771](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750743977713.png)

根据click和hover添加不同的事件

 



- 使用debounce完成延迟操作

  安装

  `npm install lodash-es --save`





####  Tooltip总结

分析

- 基础组件

  - 触发层
  - 展示层

- 开发复杂组件的时候

  - 需求分析
  - 设立开发计划——规划日程

- 使用popper.js来完成位置的展示

- 动态事件的添加

  - 使用v-on

- 实现外侧点击的功能

  useClickOutside

- 手动触发的功能

  组件实例实现对应方法

- 支持popper参数

  使用第三方库需要解决的问题

- 添加延时

  - 使用debouce来整合短事件内多次触发的回调
  - debouce可以使用ccancel取消

- 添加样式

  - 三角箭头的实现
  - 正方形旋转45度，加特定位移，再加边框

- 添加测试








####  Dropdown组件

![75081608470](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750816084704.png)

- 根据Tooltip二次开发的组件
- 显示/隐藏一个具体的，有多个选项的菜单
- 菜单有各种选项，用户可以自定义
  - 使用语义化结构
  - 使用js数据结构

![75081639178](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750816391786.png)

![75081651472](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750816514729.png)





####  总结

- 使用通用组件Tooltip进行二次开发
- 在vue单文件组件tomplate中渲染Vnode的方法
- 使用jsx编写组件

![75089946135](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750899461353.png)

![75089985240](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750899852400.png)

使用{}写出js表达式

![75090028809](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750900288099.png)

![75090034762](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750900347629.png)

![75090037756](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750900377566.png)

![75090045348](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750900453480.png)





####  Message组件

**功能分析**

- 在特定的行为的时候，弹出一个对应的提示（支持普通文本以及VNode）
- 提示在一段时间后消失（可控的定时器）
- 可以手动关闭（X关闭）
- 可以弹出多个（难点）
- 有多种类型（样式不同）
- 这个组件是调用函数完成的

**难点**

![75091834810](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750918348101.png)

样式

![75091855337](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750918553370.png)

将组件函数式渲染到页面上

![75091864855](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750918648558.png)

![75091869592](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750918695927.png)

这个instance不是实例，是这个组件返回的调用







- 操作

```vue
<template>
  <div
    class="fei-message"
    v-show="visible"
    role="alert"
    :class="{ ['fei-message--' + type]: type, 'is-close': showClose }"
  >
    <div class="fei-message__content">
      <slot>
        <RenderVnode v-if="message" :vNode="message"></RenderVnode>
      </slot>
    </div>
    <div class="fei-message__close" v-if="showClose">
    <Icon icon="xmark"></Icon>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { MessageProps } from "./type";
import RenderVnode from "@/hook/RenderVnode";
import Icon from "../Icon/Icon.vue";
import { ref,onMounted } from "vue";



const props = withDefaults(defineProps<MessageProps>(), {
  type: "info",
  duration: 3000,
});

const visible =ref(false);
function startTimer() {
  if(props.duration===0)return
  //禁止自动消失

  setTimeout(() => {
    visible.value = false;
  },props.duration)
}

onMounted(()=>{
  visible.value = true;
  startTimer();
})
</script>

<style></style>
```

![75099383816](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750993838168.png)

document.body.appendChild(container.firstElementChild!)

加上！表示不为空的断言操作

![75099486407](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1750994864078.png)

因为message是一个非实体的组件，是通过函数调用形成的，使用使用这种方式显示

message组件的位置

![75100770819](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751007708192.png)

第一个组件的props.offset+0就是top

第二个就是props.offset+height就是新的高度

上一个组件要保存一个值给下一个组件使用



![75107653589](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751076535890.png)

![75107668053](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751076680535.png)

剩下就是做一些锦上添花的事情

## 总结

![75110241423](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751102414230.png)

使用render挂载和销毁

- 组件动态构造并且传入属性

  ![75110246230](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751102462304.png)

  使用一个数组保存实例消息，并且添加对应函数进行处理



- 计算偏移量

  - top：lastBottomOffset（上一个实例留下的偏移）+offset
  - 为下一个实例预留bottomOffset：top+height
  - messageRef.value.getBoundingClientRect().height
  - 使用defineExpose暴露

- 在函数中获取这个偏移量

  - vnode.component.Componentinternallnsatance组件内部实例

    ![75110285457](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751102854575.png)

- 将instances数组改为响应式,使用shallowReactive([])浅层的reactive

- 添加两个通用钩子函数

  - useZIndex()
  - useEventListaner

- 添加有趣的动画

  使用transformY以及fade制作一个fade-up的效果

- 测试-函数式创建组件的测试

  - 使用rAF()等待动画执行完毕







####  vitePress简介

通用的静态站点生成器

- 基于md的文档工具

- 基于vite以及vue无缝衔接

- 支持使用md书写vue组件

- 适用很多场景

  - 普通的静态站点
  - 文档
  - 博客

- 多个vue以及周边工具文档页面的选则

  - vite
  - rollup
  - vue
  - vuejs Blog
  - element plus

  ​


![75116177203](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751161772031.png)

在vitepress的ms文件里，他可以被当作一个vue组件进行编写

并且可以不写template标签

![75116423307](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751164233074.png)

导入组件库的样式文件

![75116513134](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751165131343.png)

![75116524417](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751165244178.png)

![75116526894](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751165268947.png)

引入组件，使用内置css覆盖

![75116536611](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751165366111.png)







####  iuput组件需求分析

表单是一系列组件的实现，是组件库最有难点已经最核心的功能

- input组件
- switch组件
- select组件
- form组件

![75124480177](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751244801775.png)

![75124492715](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751244927150.png)

前后缀关系

![75124497107](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751244971073.png)

![75124501302](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751245013023.png)

![75124504724](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751245047240.png)

![75124505394](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751245053943.png)

暴露原生的元素

![75124513270](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751245132700.png)、

####  开发方式

![75124914758](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751249147584.png)

![75124921367](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751249213677.png)

![75124941962](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751249419624.png)

书写v-model的测试用例

![75125932371](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751259323714.png)



- 支持密码显示/不显示切换

  ![75135515774](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751355157749.png)



- 支持事件
- 原生属性

![75135833852](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751358338523.png)

这些都是要加上去的控件





####  Input组件开发总结

input组件开发是一个关键

使用了TDD的开发方式

- v-model的测试方法
- vue组件事件测试
- 设置表单组件

![75144144932](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751441449322.png)



####  代码框组件开发

![75144167474](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751441674744.png)





####  Switch组件开发

并不是标准的FORM组件，是手机端的延申

- 功能类似checkbox，所以内部是有个checkbox在工作
- 样式很独特，是巨大挑战

![75144195049](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751441950496.png)

![75144208888](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751442088881.png)

![75144276558](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751442765589.png)

一个外部框和一个内部框



![75154947563](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751549475638.png)

无障碍文档



####  select组件

下拉菜单

![75154975630](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751549756301.png)

![75154990139](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751549901394.png)

比较难的一个组件

麻烦

使用option传入来维护

![75155050783](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751550507834.png)

![75155054426](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751550544263.png)

 ![75164258714](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751642587147.png)



![75169842143](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751698421431.png)

![75170262577](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751702625776.png)

多用于搜索引擎的搜索功能

![75170390943](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751703909430.png)





####  FROM

![75176333317](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751763333178.png)

数据验证是这个组件的核心

![75176487111](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751764871114.png)

验证的基本流程

  ![75176522916](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751765229161.png)

验证的打包是在formitem里面制作的

![75176538185](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751765381850.png)

调用函数整体验证

![75176877128](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751768771289.png)

单个验证到整个验证

![75177005801](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751770058011.png)

![75177006466](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751770064668.png)

第三步使用了一个数组，来存储某些东西，然后在onMounted去使用





- 新功能

  输入错误的时候，回溯到最新的状态，在进行操作






####  总结表单

表单就是一个校验的功能

先校验组件的，把组件所有的状态传递到一个数组，给父组件判断

![75189616679](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896166799.png)

校验结果在这里

同时还有回调函数和清理数据

![75189620180](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896201803.png)

- scoped slot

  ![75189630393](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896303938.png)

- 父组件创建数据，子组件填充数据

  ![75189635070](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896350707.png)

  ![75189640737](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896407374.png)

  借助订阅模型来完成

![75189646034](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1751896460343.png)





####  导航栏

导航栏位于应用程序的顶部，即状态栏下方，主要用于明确页面位置、层级等，并且连接着父/子级结构页面，权重应高于当前页面的所有内容。如果用户不知道当前所在位置、如何回到上一步的时候，就说明导航栏一定存在问题。虽然在iOS系统中叫做「导航栏」、Android系统叫「顶部应用栏」，叫法不一，但所在位置以及起到的作用几乎一致。

功能

- 隐藏和显示
- 滚动缩小








####  虚拟列表

1. 在组件中创建一个List.vue组件,他是用来表示虚拟列表的
2. 把List.vue插件引入到页面中
3. 在主页模拟一万条数据,传递给子组件

item渲染的数据

size每一次渲染数据的dom节点

![75307230072](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1753072300729.png)











#  面试怎么说

![75505619364](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755056193649.png)

- 分层

  - rc-xxx，提供基础组件，unstyled component，只具备功能交互而不具备UI表现的
  - 样式体系：theming 
  - 基础组件
  - 复合组件，search，input+select，IconButtion
  - 业务组件

- 解耦

  - 对于每个组件都需要定义样式，ts类型，基础操作，工具方法

- 响应式设计

  - 媒体查询 media query Grid


 



状态管理

- 全局状态，基础配置，国际化组件，主题配置，react-》context,vue->demi
- 局部状态,表单场景,受控和非受控,状态是否跟表单值双向奔赴



样式体系

- Color Token:颜色色值系统

  ![75505695587](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755056955874.png)

  四种不同的颜色体系

- 样式模块化方案:css-in-js(弊端:运行时性能损耗,ssr支持),

  module css方案

- 样式优先级与覆盖:控制样式优先级



模块化

- 可复用性,对于props,events的设计非常重要,这就素为什么input和文本域要value和ocChange成对
- 公共方法和颜色



开发流程

**本地开发:使用现有的组件库二次开发**



组件库开发:

1. 工程架构

   ![75505779484](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755057794841.png)

2. ts

3. 构建打包:rollup,esbuild,swc

4. 测试,单元测试,vitest,jest+react-testing-libray

5. 流程化,规范化,自动化,script如何定义

   ![75505794648](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755057946481.png)





####  实现一个组件库

![75505934871](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755059348718.png)

依赖写在这里面

![75506200314](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755062003142.png)

使用文档来展示你的组件

![75506298767](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755062987678.png)

组件属性设置,组件事件设置

css-in-js的css风格![75506303383](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063033838.png)

![75506305784](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063057847.png)

![75506311779](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063117796.png)

这样做,去使用

![75506316558](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063165581.png)

**使用react去实现那个组件库**



流程

![75506335347](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063353472.png)

然后就是文档编写

再到一些分包,打包,发布



**发布**

![75506340272](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063402723.png)

使用docker进行发布

![75506344183](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063441832.png)

面试

![75506348925](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063489257.png)

![75506350653](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063506536.png)





复盘,失败不可怕

![75506357948](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063579485.png)

![75506360787](C:\Users\zxh\Desktop\前端\component-lib\UI.assets\1755063607875.png)

