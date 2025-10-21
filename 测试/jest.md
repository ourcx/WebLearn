# jest单元测试

![75249363998](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752493639980.png)

瀑布流开发风险会随着开发而增加

![75249377536](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752493775364.png)

![75249387643](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752493876430.png)

代码不测试，就越烂，越烂bug就越多，所以进行单元测试是有必要的



####  单元测试和自动化的关系

业务不用频繁上线，并且可以整个团队有足够多的人力来进行覆盖手工测试，可以不用单元测试，或者你不在意也不用

 

- 流水线：持续集成我们的产品代码和测试代码，持续不断的验证每一次提交代码是否符合预期，自动化可以帮我们快速完成这个步骤

- ![75249449226](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752494492260.png)

- 测试最终目的就是减少推向市场的时间，快速响应市场的变化，开发高质量的代码

- BUG的破窗理论

  BUG会累积起来，导致越来越大的破坏

![75249471339](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752494713397.png)

前端单元测试

e2e-》集成测试-》单元测试

越往上成本时间越多，单元测试离开发者最近

![75249481966](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752494819669.png)

####  jest单元测试入门

![75249508436](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752495084369.png)

一个js文件

```javascript
const sum = (a,b)=>a+b
module.exports={sum}
```

一个测试文件

```javascript
const {sum}=require('./math')

describe("math module",()=>{
    test('shuld return sum result when two number pus',()=>{
        const number  =1;
        const anotherNumber=2;
        const result = sum(number,anotherNumber)
        
        expect(result).toEqual(3)
    })
})
```

然后在命令行运行`jest-demo yarn test`

![75249565847](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752495658474.png)

3A的测试模式

- 断言语句

- > toBe(value)判断是否符合某个值
  >
  > toEqual(value)可以看这个值的对象,不只是引用等
  >
  > toBeFalsy()可以判断布尔值
  >
  > toHaveLength(number)来数组的长度
  >
  > toHaveBeenCalled()方法是否被调用
  >
  > toHaveBeenCalledTimes(number)调用的次数
  >
  > toThrow(error?)是否有异常
  >
  > toMatchSnapshot(propertyMatchers?,hint?)这次的快照是否等于上一次
  >
  > toMatchinlineSnapshot(propertyMatchers?,inlineSnapshot)
  >
  > expect.extend(matchers)扩展属于自己的断言器



####  jest测试的模块依赖

![75249729293](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497292938.png)

 模块间依赖的处理方法,fake,stub,mock,spy

![75249770432](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497704320.png)

- mock用于隔离整个模块

  ![75249778071](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497780716.png)

  没有被mock的部分不会被返回.分隔了整个模块

- Stub用于模拟特定行为

  ![75249784631](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497846310.png)

  调用这个方法的时候始终返回true

- spy监听模块行为

  ![75249788244](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497882448.png)

  使用spy监听video.play()

- ![75249794644](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752497946448.png)

  模块的独立性,不会mock很多次

  ```javascript
  //services是其他模块引入进来的东西
  const searchNames = (term) => {
      const matches = services.getNames().filter((name) => {
          return name.includes(term);
      });
      return matches.length > 3 ? matches.slice(0, 3) : matches;
  };
  ```

  测试文件

  Mock（模拟）主要用于单元测试中 **隔离外部依赖**，确保测试只关注当前函数的逻辑，而不受其他模块（如 API、数据库、文件系统等）的影响。

  - **services.getNames() 是一个外部依赖**，可能从 API、数据库或文件读取数据。
  - **测试时不应依赖真实数据**，否则：
    - 测试可能变慢（如网络请求）。
    - 测试可能不稳定（如数据库数据变化导致测试失败）。
    - 无法覆盖所有边界情况（如空列表、异常情况）。

```javascript
  import searchNames from './searchNames';

  jest.mock('./services',()=>{
      getNames:jest.fn(()=>{
          ["John","Paul"."George","Ringo"]
      })
  })

  test('should return empty result when no matches found', () => {
      // given
      const keyword = 'Frank';
      getNames.mockImplementation(()=>[])
      // when
      const result = searchNames(keyword);
      
      // then
      expect(result).toEqual([]);
  });

  test('should return target result when found search', () => {
      // given
      const keyword = "John";
      getNames.mockImplementation(()=>["John","paul","George","Ringo"])
      // when
      const result = searchNames(keyword);
      
      // then
      expect(result).toEqual(['John', 'Johnny']); // 假设只返回前 3 个匹配项
  });
```

  

- 其他测试模式

  当运行 `jest --watch` 或 `npm test -- --watch` 时，Jest 会进入 **交互式监控模式**，提供以下功能：

  Watch 模式下的快捷键

  | 快捷键    | 功能                                                         |
  | --------- | ------------------------------------------------------------ |
  | **f**     | 只运行 **失败的测试**（"failed tests only"）                 |
  | **o**     | 只运行 **与修改文件相关的测试**（需 Git 支持）               |
  | **p**     | 按 **文件名** 过滤测试（输入正则表达式匹配文件名）           |
  | **t**     | 按 **测试名称** 过滤测试（输入关键字匹配 `test('name', ...)`） |
  | **q**     | 退出 Watch 模式                                              |
  | **Enter** | 重新运行所有测试                                             |

  可以自动运行等,热更新


  `jest-demo yarn test --coverage`

  得到一些数据,可以知道要被测试的东西

  可以得到一个报告,有测试覆盖率等东西

  

  ![75250003786](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752500037862.png)

  快照测试,直接用结果取代这个地方的值,自动将测试的那一行进行相关的修改

  使用这个快照比较方便,使用另外一个会生成一个文件,可读性比较差





测试原则

![75250013439](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752500134398.png)

四大部分的测试

1. 模拟用户的行为来验证功能
2. 集成测试
3. 单元测试,看模块是否符合预期
4. 静态检查,在写代码的时候检查

你的测试和你的软件使用方式越相近,一切就越好

![75250022402](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752500224027.png)

####  组件化UI测试--Library

组件化-是单元测试流行起来

UI=f(data)

组件可以被这样进行表示

 ![75254147493](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752541474931.png)

例子

![75254149951](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752541499514.png)

查询元素的方法

![75254152284](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752541522842.png)

getBy只能找到一个，但getAllBy可以找到多个,使用queryBy会报错,但getBy只会判空

findBy会等待异步的情况，queeryBy会有一个兼容

![75254189103](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752541891036.png)

UI行为

![75254193003](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752541930036.png)

找到相关的输入框，然后输入15，再去验证整个结果



例子

```javascript
import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', () => {
    // given
    render(<App />);
    // when
    const linkElement = screen.getByText(/learn react/i);
    // then
    expect(linkElement).toBeInTheDocument();
});
```





####  测试异步代码

![75254690267](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752546902670.png)

![75254692287](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752546922879.png)

![75254695286](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752546952864.png)

使用async会更清晰好读

![75254699627](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752546996270.png)



####  Cypress组件级集成测试

使用浏览器的测试框架

安装

![75254727379](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752547273795.png)

因为cyoress以及相关的react需要单独运行，需要安装webpack DEV server

再运行npx cypress open

```javascript
import {mount} from '@cypress/react'
import {Todo} from './Todo'

describe('Todo',()=>{
    it('should renders new item',()=>{
        const todo = {
            id:1,
            text:"xxxx",
            isCompleted:false,
        }
        mount(<Todo todo={todo.id} toggleTode={cy.stub()}/>)
       cy.contains('.todo button','complete')
    })
})
```

使用npx cypress open-ct打开浏览器页面

在mount的第二个参数添加css样式

![75254818286](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752548182864.png)

或者使用mountWithStyle这个函数

![75254832847](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752548328476.png)



事件测试

```javascript
it("should delete a todo", () => {
    const todo = {
        id: 1,
        text: "test todo",
        isCompleted: true,
    };
    const removeTodo = cy.stub().as('remove');
    
    mountWithStyle(
        <Todo
            todo={todo}
            index={todo.id}
            toggleTodo={cy.stub()}
            removeTodo={removeTodo}
        />
    );
    
    cy.contains(".todo", "test todo")
        .find("[data-testid='remove-todo']")
        .click();
    
    cy.get('@remove').should('be.calledWith', todo.id);
});
```

这个测试的时间和成本是更高的





####  状态数据流测试

![75254915645](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752549156453.png)

redux的状态管理

![75254943786](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752549437868.png)

职责分离,修改和查询不能混在一起

![75254958781](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752549587815.png)

所以redux抽象出了这些规则







#  vitest

![75264817778](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752648177785.png)

 文件命名xxxx.spec.ts或者xxxx.test.js

```typescript
import {test,expect} from 'vitest'

test("first",()=>{
    expect(1+1).toBe(2)
})
```

运行

![75264823388](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752648233889.png)

pnpm test

支持热更新



**配置文件**

在vite.config.ts配置

![75264877007](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752648770072.png)

支持全局使用vitest

配置vitest.config.ts的话就去看官方文档

 



####  测试组件

测试组件要安装jsdom

![75265046097](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752650460974.png)

解析jsx的话,就还需要配jsx

配置完成后再测试文件里面使用

![75265049284](C:\Users\zxh\Desktop\前端\测试\jest.assets\1752650492848.png)

使用mount挂载组件

   