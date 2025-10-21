#  react

![75465729123](C:\Users\zxh\Desktop\前端\react\react.assets\1754657291231.png)

![75465741622](C:\Users\zxh\Desktop\前端\react\react.assets\1754657416223.png)

 ![75466561296](C:\Users\zxh\Desktop\前端\react\react.assets\1754665612962.png)

  ![75466686665](C:\Users\zxh\Desktop\前端\react\react.assets\1754666866655.png)

声明式：小明，出去操场跑五圈，有很多封装好的类

命令式：小明，起来，滚出去，往东50米，到操场，跑一圈，再跑五圈，常见于数据库

是两种编码风格



![75466733640](C:\Users\zxh\Desktop\前端\react\react.assets\1754667336403.png)



单项数据流

双向数据流









####    开始去使用

![75466757254](C:\Users\zxh\Desktop\前端\react\react.assets\1754667572542.png)

直接引入的方式

react也是分react-dom网页开发，react-native的app开发

![75467118751](C:\Users\zxh\Desktop\前端\react\react.assets\1754671187513.png)

jsx完成了上面那个东西的功能，进一步进行了封装

 语法规则：

1. 所有标签必须闭合，最外层必须只有一个根元素，遇到与js关键字同名的属性要特殊处理，找不到其他标签就用空标签<></>来包裹

   ![75470208463](C:\Users\zxh\Desktop\前端\react\react.assets\1754702084633.png)

2.  ![75470220064](C:\Users\zxh\Desktop\前端\react\react.assets\1754702200648.png)

   js代码相关的

3. 注释的写法

   {/\*这里写你的注释*/}

   或者使用这样的注释

    {

   //写你的注释

   }

4. 渲染数据

   没有返回值的数据不能写进jsx里面,三元表达式可以替代if,就是用逻辑运算符实现简单判断

   大部分对象不能嵌入jsx里面,匿名函数或者函数名,无法直接嵌入jsx里面

   值为数组的对象,如果子成员不是对象,就可以渲染

   map来循环数组即可

   ![75470363338](C:\Users\zxh\Desktop\前端\react\react.assets\1754703633385.png)

   ​

5. 样式

    ![75471812157](C:\Users\zxh\Desktop\前端\react\react.assets\1754718121578.png)

   行间样式

   ![75471837997](C:\Users\zxh\Desktop\前端\react\react.assets\1754718379972.png)

   官方推荐行间样式，但是大多数人都是采用目录的方式来报错不同组件的代码，往往一个页面就是一个组件，而这个页面的css样式文件和js组件文件则会保存当前目录下

6. 代码转义

   ![75471871667](C:\Users\zxh\Desktop\前端\react\react.assets\1754718716675.png)

   直接使用html是行不通的，防范了xss攻击








####  组件化

1. 脚手架

   ![75471976019](C:\Users\zxh\Desktop\前端\react\react.assets\1754719760196.png)

    或者用yarn create vite

   或者直接npm init vite

   ![75472076108](C:\Users\zxh\Desktop\前端\react\react.assets\1754720761083.png)

    React.StractMode开启React语法/兼容性的严格模式，如果在项目中使用了废弃或者已经过时的接口或者语法时，会出现警告

2. 组件写法

   ![75472620414](C:\Users\zxh\Desktop\前端\react\react.assets\1754726204141.png)

   ```javascript
   function App()
   //函数式组件,函数名就是组件名,一般和文件名相同,必须首字母大写,返回jsx视图代码
   export default App
   //在低版本下使用组件的jsx代码,都会被babel转码成react.createElement()
   //所以在低版本React里,要注意给组件导入React

   class App extends React.Component{
       constructor{
           super();
           this.mag=`ssss`
       },
           render(){
               return(
               <div className='App'>{this.msg}</div>
               )
           }
   }
   //类组件,类名就是组件,组件名一般和文件名同名,必须首字母大写
   //类组件,必须直接或者间接集成于React.Component类,并且jsx代码必须通过render方法return返回
   export default App
   ```

    ![75472771609](C:\Users\zxh\Desktop\前端\react\react.assets\1754727716097.png)

   无状态就是函数式组件在那时做不到修改数据的办法

   不能进行试图界面的修改





####  组件的嵌套

![75473035751](C:\Users\zxh\Desktop\前端\react\react.assets\1754730357518.png)

就是这样的

快速生成相关代码的快捷键

rss,rsf,rsc

![75473067756](C:\Users\zxh\Desktop\前端\react\react.assets\1754730677561.png)







####  react的事件处理

 类组件的事件调用

![75478836577](C:\Users\zxh\Desktop\前端\react\react.assets\1754788365772.png)

  建议使用下面那种格式来写，绑定事件处理

![75478902360](C:\Users\zxh\Desktop\前端\react\react.assets\1754789023605.png)

如果要在·第一种里面使用this，可以使用bind将this传入

![75478982776](C:\Users\zxh\Desktop\前端\react\react.assets\1754789827767.png)





深入了解React事件机制

**![75479016269](C:\Users\zxh\Desktop\前端\react\react.assets\1754790162690.png)**

合成事件，把原生的浏览器事件进行包装，方便跨浏览器使用

 ![75479063889](C:\Users\zxh\Desktop\前端\react\react.assets\1754790638898.png)



![75479079354](C:\Users\zxh\Desktop\前端\react\react.assets\1754790793544.png)

阻止事件冒泡

因为react的合成事件，是在事件冒泡阶段执行的，所以会比js原生事件要慢



####   react组件ref

![75479125006](C:\Users\zxh\Desktop\前端\react\react.assets\1754791250069.png)

三大属性

```javascript
import React from "react";

class App extends React.Component {
    // 设置一个ref属性
    input = React.createRef();

    // 添加获取内容的方法
    getContent = () => {
        // 通过ref获取input元素的值
        const value = this.input.current.value;
        console.log("输入框内容:", value);
        // 这里可以添加其他处理逻辑
        alert(`输入内容: ${value}`);
    };

    render() {
        return (
            <div>
                {/* ref属性绑定DOM节点 */}
                <input type="text" ref={this.input} />
                <button onClick={this.getContent}>检查</button>
            </div>
        );
    }
}

export default App;
```





####  state

![75484059576](C:\Users\zxh\Desktop\前端\react\react.assets\1754840595767.png)

一个自动响应的东西

 ![75484108217](C:\Users\zxh\Desktop\前端\react\react.assets\1754841082179.png)

修改数据要使用setState

![75484110807](C:\Users\zxh\Desktop\前端\react\react.assets\1754841108071.png)







####  props

![75484151846](C:\Users\zxh\Desktop\前端\react\react.assets\1754841518463.png)

 ![75484181018](C:\Users\zxh\Desktop\前端\react\react.assets\1754841810187.png)

父组件

![75484184605](C:\Users\zxh\Desktop\前端\react\react.assets\1754841846057.png)

![75484195980](C:\Users\zxh\Desktop\前端\react\react.assets\1754841959800.png)

子组件也可以加一个验证器，验证格式什么的，对外界传递的格式设置约束条件，不过需要安装prop-types这个插件，

子组件接收数据

 

####  子传父

和vue一样都是利用了函数来完成

![75491733849](C:\Users\zxh\Desktop\前端\react\react.assets\1754917338498.png)

 ![75491781128](C:\Users\zxh\Desktop\前端\react\react.assets\1754917811280.png)

在子组件进行类型约束

![75491805080](C:\Users\zxh\Desktop\前端\react\react.assets\1754918050804.png)

函数里面使用

![75491849425](C:\Users\zxh\Desktop\前端\react\react.assets\1754918494254.png)









####  # 进阶使用

####  受控和非受控组件

 根据是否受外界数据控制可分为受控和非受控组件

![75500425157](C:\Users\zxh\Desktop\前端\react\react.assets\1755004251575.png)



注意

![75500422879](C:\Users\zxh\Desktop\前端\react\react.assets\1755004228799.png)

在state里面的那个数据，就是父组件不能影响的，改变props也不行

函数组件默认是无状态组件，类是有状态的，就是有没有state的意思





![75500555298](C:\Users\zxh\Desktop\前端\react\react.assets\1755005552987.png)

通过ref拿到里面的数据





![75500576493](C:\Users\zxh\Desktop\前端\react\react.assets\1755005764931.png)

通过state和value来进行





####  组件间通讯

1. 状态提升：

   不同父组件的组件，把状态向上提升，直到找到共同的父组件

   另外一个组件修改数据，父组件数据变化，另外一个组件的数据也发生变化

   **但有的数据存放的地方过于高了，可能碰不到**

2. 消息订阅模型

   ![75500757105](C:\Users\zxh\Desktop\前端\react\react.assets\1755007571051.png)

   ![75500798102](C:\Users\zxh\Desktop\前端\react\react.assets\1755007981028.png)

    先订阅在发布

   订阅代码一般写在挂载之后

   ![75500911725](C:\Users\zxh\Desktop\前端\react\react.assets\1755009117257.png)

   写在这个生命周期，来订阅 

3. 中央文件控制数据，给组件分发,上下文的方式






####  非父组件的通信-context

![75508596173](C:\Users\zxh\Desktop\前端\react\react.assets\1755085961732.png)

执行上下文

![75508600394](C:\Users\zxh\Desktop\前端\react\react.assets\1755086003942.png)

生成那个文件

![75508979034](C:\Users\zxh\Desktop\前端\react\react.assets\1755089790340.png)

这个GlobalContext.Provider导出来的是一个组件

在父组件共享了num

![75509016718](C:\Users\zxh\Desktop\前端\react\react.assets\1755090167186.png)

 在子组件使用GlobalContext.Consunmer,用回调的方式,进行提取上下文\



接收其他组件的数据也同理

![75509030018](C:\Users\zxh\Desktop\前端\react\react.assets\1755090300184.png)

使用函数









####  插槽

和vue是类似的,所有传进去的html在this.props.children里面,这个东西是一个数组,你可以更灵活操作

![75509208640](C:\Users\zxh\Desktop\前端\react\react.assets\1755092086400.png)







####  生命周期

constructor是最早执行的生命周期，一般用来设置初始化state

 ![75513310339](C:\Users\zxh\Desktop\前端\react\react.assets\1755133103393.png)

  ![75513331002](C:\Users\zxh\Desktop\前端\react\react.assets\1755133310023.png)

 在render里面不要进行修改dom或者state的操作，会引发错误，死循环，因为修改dom会重新执行render



componetDidMount一般是页面渲染完成的初始化操作，完成请求数据，监听，定时等任务

**更新阶段**

![75513462528](C:\Users\zxh\Desktop\前端\react\react.assets\1755134625283.png)

![75513638889](C:\Users\zxh\Desktop\前端\react\react.assets\1755136388895.png)

要有返回值，是布尔，react他提供的一种优化方法

![75513692898](C:\Users\zxh\Desktop\前端\react\react.assets\1755136928984.png)





**static getSnapshotBeforeUpdate这个方法替代了componetnWillUpdate这个生命周期**

![75513806574](C:\Users\zxh\Desktop\前端\react\react.assets\1755138065746.png)

必须要有返回值

 

![75513712506](C:\Users\zxh\Desktop\前端\react\react.assets\1755137125067.png)

这个是完成一些收尾的工作的，比如向服务端保存数据什么的，还有扬子鳄事件的解绑，定时器什么的



> 函数式组件没有生命周期这个东西
>
> 函数使用生命周期要通过hook来进行



**实战**

![75532697489](C:\Users\zxh\Desktop\前端\react\react.assets\1755326974895.png)

![75532748897](C:\Users\zxh\Desktop\前端\react\react.assets\1755327488973.png)

在生命周期componentDiMount去执行axios这个东西，注意封装axios





####  hooks

![75532858245](C:\Users\zxh\Desktop\前端\react\react.assets\1755328582453.png)

![75532859468](C:\Users\zxh\Desktop\前端\react\react.assets\1755328594681.png)



- 基础hook

  1. useState:用于函数式组件实现有状态

     ![75532932056](C:\Users\zxh\Desktop\前端\react\react.assets\1755329320565.png)

     函数体相当于render

     ```jsx
     function App() { // 函数式视图的函数体就相当于类组件的render
         // const [状态变量名，修改状态的唯一函数] = useState(状态初始值)
         const [count, setCount] = useState(0);
         return (
             <div onClick={()=>setCount(count+1)}>{count}</div>
         );
     }

     export default App;
     ```

  2. useEffect

     ![75533009896](C:\Users\zxh\Desktop\前端\react\react.assets\1755330098961.png)

     副作用是:异步的axios在渲染中出现死循环的情况

     所以异步代码里修改属性/states的操作,属于副作用操作

     ```jsx
     import React, { useState, useEffect } from 'react';
     import axios from 'axios';

     function App() { // render
         const [data, setData] = useState([]); // 参数是默认值

         // 参数1，就是副作用代码，
         // 参数2，即副作用依赖，
         // 如果参数2是一个空数组，则表示当前代码是没有任何外部依赖的副作用代码，只需要执行一次。
         useEffect(() => {
             axios.get("/data.json").then(response => {
                 setData(response.data);
             });
         }, []);

         // 返回JSX内容
         return (
             <div>
                 {/* 渲染内容 */}
             </div>
         );
     }

     export default App;
     ```

     依赖状态的副作用代码

     ```jsx
     import React, { useState, useEffect } from 'react';
     import axios from "axios";

     function App(props) {
         const [city, setCity] = useState("北京"); // 参数是默认值
         const [weatherData, setWeatherData] = useState([]); // 初始值应为空数组

         useEffect(() => {
             axios.get(`https://v0.tianqiapi.com/?version=day&unit=m&language=zh&query=${city}&appid=43656176&appsecret=142og6Lm`)
                 .then(response => {
                     console.log(response.data.month);
                     setWeatherData(response.data.month);
                 });
         }, [city]);

         return (
             <div>
                 <input 
                     type="text" 
                     value={city} 
                     onChange={(event) => setCity(event.target.value)}
                 />
                 <ul>
                     {weatherData.map((item, key) => (
                         <li key={key}>
                             {item.date} {item.dateOfWeek} {item.day.narrative}
                         </li>
                     ))}
                 </ul>
             </div>
         );
     }

     export default App;
     ```

     useEffect的卸载阶段

     ![75533355033](C:\Users\zxh\Desktop\前端\react\react.assets\1755333550338.png)

     类似与生命周期componentWillunmont的作用

  3. useLayoutEffect

     ![75533363942](C:\Users\zxh\Desktop\前端\react\react.assets\1755333639426.png)

     EOM更新以后直接用,和useEffect是在页面渲染后才用

  4. useMemo

     ![75533384652](C:\Users\zxh\Desktop\前端\react\react.assets\1755333846527.png)

      常见于父组件数据变化,引发子组件变化的情景

      为了减少这些重复计算,提供了useState,useCallback,useMemo这些钩子让我们可以在函数式组件中,对变量,方法,数据结果进行缓存

     ![75533512931](C:\Users\zxh\Desktop\前端\react\react.assets\1755335129318.png)

     缓存了一段代码的执行结果,类似于计算属性.

     参数1:是函数,函数会在组件挂载以后useMemo被运行一次并得到应该缓存结果作为返回值

     参数2:数组,当数组中记录的状态发生变化是,参数一里面的函数会重新执行,得到一个新的缓存结果,如果参数2是一个空数组,参数1会在每次组件渲染的时候重新执行声明

  5. useCallback

     可以认为是useMemo的语法糖,底层是一样的,主要区别是React.useMemo将调用fn函数并返回其结果,而React.useCallbak将返回fn函数而不调用它,因此开发中,使用useMemo缓存变量或组件,而使用useCallback缓存函数

     ![75533700656](C:\Users\zxh\Desktop\前端\react\react.assets\1755337006560.png)
     ![75533729581](C:\Users\zxh\Desktop\前端\react\react.assets\1755337295813.png)

     两者声明的比较

     使用的时候也有所不同

     ![75533735483](C:\Users\zxh\Desktop\前端\react\react.assets\1755337354831.png)

  6. useRef

     是三大属性里面的基础,就是获得dom对象

  7. userContext，以前学过的Context在函数式组件的用法与类组件的用法是一致的

     共同文件

     ```jsx
     // 创建Context对象（两种组件通用）
     import React from 'react';

     const ThemeContext = React.createContext('light'); // 默认值

     // 父组件
     function App() {
       return (
         <ThemeContext.Provider value="dark"> {/* 覆盖默认值 */}
           <FunctionComponent />
           <ClassComponent />
         </ThemeContext.Provider>
       );
     }
     //提供值的
     ```

     ```jsx
     //函数式里面
     import React, { useContext } from 'react';

     // 子组件
     function FunctionComponent() {
       // 直接获取Context值
       const theme = useContext(ThemeContext);

       return (
         <div style={{ 
           background: theme === 'dark' ? '#333' : '#fff',
           color: theme === 'dark' ? 'white' : 'black'
         }}>
           函数组件当前主题: {theme}
         </div>
       );
     }
     ```

     ```jsx
     //类组件里面
     class ClassComponent extends React.Component {
       render() {
         return (
           // 通过Consumer包裹
           <ThemeContext.Consumer>
             {theme => (
               <div style={{ 
                 background: theme === 'dark' ? '#333' : '#fff',
                 color: theme === 'dark' ? 'white' : 'black'
               }}>
                 类组件当前主题: {theme}
               </div>
             )}
           </ThemeContext.Consumer>
         );
       }
     }
     ```


     //Context提供的provider生产者和consumer消费者组件在函数式组件的用法是一模一样的,但是上面的嵌套层级太深了,所以在函数式组件就有了useContext这个钩子
     ```

  8. useReducer 

     相当于useState的升级版,useState的替代方案,功能上类似useState,都是保存状态,不同点在于可以定义一个reducer函数,用来处理复杂数据,目的是为了解决函数式组件中逻辑代码与试图代码耦合性太高的情况

     ![75541441247](C:\Users\zxh\Desktop\前端\react\react.assets\1755414412471.png)

     > 是redux的影响

     ![75540972392](C:\Users\zxh\Desktop\前端\react\react.assets\1755409723921.png)

     相关的参数

     ```jsx
     //参数具体写法
     import React, { useReducer } from 'react';

     // 1. 定义reducer函数 - 处理状态更新的纯函数
     function reducer(state, action) {
       switch (action.type) {
         case 'INCREMENT':
           return { ...state, hum: state.hum + 1 };
         case 'DECREMENT':
           return { ...state, hum: state.hum - 1 };
         case 'RESET':
           return { ...state, hum: state.initHum }; // 使用初始值重置
         default:
           throw new Error(`未知的action类型: ${action.type}`);
       }
     }

     // 2. 初始状态
     const initialState = {
       hum: 0,  // 初始值
       initHum: 0 // 保存初始值用于重置
     };

     // 3. init函数 - 对初始状态进行转换
     function init(initialState) {
       // 这里可以添加初始化逻辑，例如从localStorage读取数据
       return {
         ...initialState,
         hum: initialState.hum + 10 // 示例：将初始值增加10
       };
     }
     ```

  9. 自定义hook

     自定义hook是`1一个函数,名字以use开头,自定义hook内部可以调用其他hook,还可以将某些组件逻辑取到可重用的函数中,让代码结构更清晰,易维护,比如把axios进行封装

     ​





####  路由

![75541413245](C:\Users\zxh\Desktop\前端\react\react.assets\1755414132456.png)

`npm i react router-dom@next`

使用

![75541466210](C:\Users\zxh\Desktop\前端\react\react.assets\1755414662106.png)

 在app.jsx组件里面

![75541505462](C:\Users\zxh\Desktop\前端\react\react.assets\1755415054622.png)

![75541531881](C:\Users\zxh\Desktop\前端\react\react.assets\1755415318810.png)

实际开发不是这样写的,要你去准备路由表

所以有routes.jsx文件

```jsx
// src/routes/index.js
import { createBrowserRouter } from "react-router-dom";
import Layout from "@/layout"; // 布局组件
import Home from "@/pages/Home"; // 首页
import About from "@/pages/About"; // 关于页
import UserProfile from "@/pages/User/Profile"; // 用户详情页
import NotFound from "@/pages/Error/404"; // 404页面

// 路由配置表
const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />, // 公共布局
    errorElement: <NotFound />, // 错误处理
    children: [
      {
        index: true, // 默认子路由
        element: <Home />,
      },
      {
        path: "about",
        element: <About />,
      },
      {
        path: "user/:id", // 动态路由
        element: <UserProfile />,
        loader: ({ params }) => {
          // 路由加载器（数据预加载）
          return fetchUserData(params.id);
        },
      },
      {
        path: "*", // 404 匹配
        element: <NotFound />,
      },
    ],
  },
]);

export default router;
```

在app.jsx里面可以写成

![75541574384](C:\Users\zxh\Desktop\前端\react\react.assets\1755415743842.png)

导入编好的路由映射表

官方推荐的主旨是低耦合,易复用





页面跳转功能

![75558918175](C:\Users\zxh\Desktop\前端\react\react.assets\1755589181754.png)

 这两个的区别在于是否高亮

也可以使用编程式的跳转

![75558935620](C:\Users\zxh\Desktop\前端\react\react.assets\1755589356202.png)

路径参数

![75558957614](C:\Users\zxh\Desktop\前端\react\react.assets\1755589576147.png)

![75558990970](C:\Users\zxh\Desktop\前端\react\react.assets\1755589909704.png)

在对应的组件下拿到数据，使用useParams这个hook来接受路径参数

查询参数和这个是类似的

 

**嵌套路由**![75559135675](C:\Users\zxh\Desktop\前端\react\react.assets\1755591356757.png)

注意子路由的开头问题

**占位符路由**

表示路由映射表中匹配的子路由对应的组件将在此处展示

![75559168640](C:\Users\zxh\Desktop\前端\react\react.assets\1755591686400.png)

Outlet这个就是占位符路由







####  redux

![75568204205](C:\Users\zxh\Desktop\前端\react\react.assets\1755682042053.png)

redux不止可以在react使用

![75568366529](C:\Users\zxh\Desktop\前端\react\react.assets\1755683665297.png)

安装：去看官网安装

![75568385722](C:\Users\zxh\Desktop\前端\react\react.assets\1755683857225.png)





使用

![75568388665](C:\Users\zxh\Desktop\前端\react\react.assets\1755683886658.png)

 

![75568423232](C:\Users\zxh\Desktop\前端\react\react.assets\1755684232323.png)



相关的回调

![75568607835](C:\Users\zxh\Desktop\前端\react\react.assets\1755686078352.png)

执行unsubscribe()可以解绑这个监听





#### 集成ANTD组件库

下载安装

1. 在main.js里面引入

   ![75574408772](C:\Users\zxh\Desktop\前端\react\react.assets\1755744087722.png)

   样式库

2.  引入

   ![75574457290](C:\Users\zxh\Desktop\前端\react\react.assets\1755744572909.png)

   开箱即用

   ![75574526012](C:\Users\zxh\Desktop\前端\react\react.assets\1755745260125.png)





####  redux的补充

![75574548251](C:\Users\zxh\Desktop\前端\react\react.assets\1755745482514.png)

 ![75574555767](C:\Users\zxh\Desktop\前端\react\react.assets\1755745557679.png)

使用provide这个组件

在App.jsx里面有其他操作,引入connect的,需要你进行其他操作

![75574611072](C:\Users\zxh\Desktop\前端\react\react.assets\1755746110723.png)



Antd相当于element,很经典的一个组件库