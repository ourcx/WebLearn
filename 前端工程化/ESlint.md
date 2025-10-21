 #  ESlint

![75256690545](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752566905451.png)

![75256692605](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752566926052.png)

#### 你在之前的项目有用过相关的内容吗,除了eslint还有什么其他方案

做过前端工程化,规范化,自动化相关架构

####   什么是规范,为什么要

- 文件结构,monorepo

- 命名规则,自己定义

- 提交规范,commitlint

- 代码层面,

  eslint,

  oxc,

  prettier

  解决的问题,代码风格统一,强硬的手段eslint,ts->commitlint提交上去做卡点控制

- 规则的三个等级

  off,warn,error



####  跟CI/CD整个卡点结合

CI/CD的工作流控制方面增设卡点,用来执行eslint的校验

![75256849018](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752568490183.png)



####  深入项目

![75256879452](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752568794524.png)

鱿鱼须为什么要用rust重构工具链的原因,统一的ast极大提高了性能





####  eslint配置

 初始化一个包`npm init -y`

创建模块化规范文件`eslint.cofig,js`,和node_modules同级

吗

![75257009626](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752570096263.png)

console是error的规则

![75257017747](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752570177476.png)

没有用的变量提示报错

`no-dupe-keys:error`禁止重复内容

去官方文档有很多内置规则和打包好的规则还有第三方的规则

也能自己进行封装规则集

打包好的规则:@eslint/js

![75257051469](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752570514699.png)

 

####  扁平配置和多类型文件支持

![75257073211](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752570732116.png)

扁平化配置,不同文件使用不同的规范,也能针对某一个文件进行规范





####  ts项目配置和编译器理解

要先装一个@typescript-eslint/parser编译器

![75257183383](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752571833839.png)

指定语言的解析器

![75257194830](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752571948306.png)

也能配置很多相关的规则

比如semi是否需要分号

**也可以和其他项目进行结合来实现**





####  在复杂VUE应用里面使用

vue需要单独的转换器去安装`vue-eslint-parser`

> parser基本都是编译器相关的错误

![75258329954](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752583299540.png)

在eslint的文件里定义,外层是vue的解析,内层是ts的解析



####  项目推荐

![75258375840](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752583758404.png)







####  底层实现

 ![75258533916](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752585339161.png)

![75258536833](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752585368335.png)

设计模块‘

![75258551393](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752585513931.png)





####  团队专属规则设计--插件

规则集＋插件方式来扩展eslint9体系

下层很多个rule，组合为rules，rules组合为plugun

- 具体实现

插件，规则的协议

这个是插件化设计里面重要的一环

1. 插件基座
2. 插件协议
3. 插件接入协议（webpack,eslint,tapable实现插件以及生命周期钩子）

```javascript
export const noMiaomaVars={
    meta:{
        messages:{
            noMiaoVars:"禁止使用miaoma开头的变量"
        }
    },
    create(context){
        return{
            Odentifier(node){
                //准则就是匹配到某个东西就上报错误
                console.log(node.name)
                if(node.name==="miaoma"){
                    context.report({
                        node,
                        messageId:"noMiaOmaVars",
                        data:{
                            name:node.name
                        }
                    })
                }
            }
        }
    }
}
```

这里就是一个单一规则集了

```javascript
//导入规则集
import {noMiaomaVars}from "..////"

export const eslintMiaomaPlugin ={
    rules:{
        "no-miaoma-vars":noMiaomaVars,
        
    }
}
```

扁平化的好处,对于插件而言,也变成了具名插件

把刚才写完的规则导入进来

![75258794502](C:\Users\zxh\Desktop\前端\前端工程化\ESlint.assets\1752587945020.png)

miaoma就是这个插件的名字

也可以使用和rules同级的ignorePatterns来忽略某些文件



