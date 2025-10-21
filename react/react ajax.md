#  react ajax

- react本身只关注界面，不包含ajaxj请求的代码
- 自己封装或者第三方
- 常见ajax封装：jQuery或者axios





####  使用

```react
import React,{Component}form 'react'
import axios from 'axios'

export default class App extends Component{
    getStudentData=(??????)=>{
        axios.get().then(
        response=>{},
        error=>{}
        )
    }
    
    render(){
        <div>
            <button onClick={this.getSthudentData}>点击我</button>
        </div>
    }
}
//会遇到跨域的问题
```

react脚手架配置代理的方法，在package.json里面使用

![74805770239](C:\Users\zxh\Desktop\前端\react\react ajax.assets\1748057702398.png)

代理的路径不对（本地有这个路径），会访问你本地的东西，并不是全部都转发倒代理的地址去

要配置两个，就不能在这里配置了

![74806990875](C:\Users\zxh\Desktop\前端\react\react ajax.assets\1748069908753.png)

在src建立一个setupProxy.js

![74807109993](C:\Users\zxh\Desktop\前端\react\react ajax.assets\1748071099935.png)

多个代理可用配置，但配起来比较繁琐







####  消息订阅和发布

 类似于对讲机的消息发布机制，和vue2那个很像

工具库：PubSubJS

安装：` npm i pubsub-js`

![74826005343](C:\Users\zxh\Desktop\前端\react\react ajax.assets\1748260053437.png)

监听subscribe和发布publish

下面是异步版本的监听和发布：publishSync发布和

取消监听unsubscribe