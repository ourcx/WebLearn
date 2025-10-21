# 1.Typescript的价值

js的核心是灵活，但是会加大代码的维护困难，所以需要一些类型的约束

- ts更像后端java，让js可以去开发大型应用
- 类型系统可以帮助我们在写代码时提供丰富的语法提示
- 在编写代码时进行类型检查从而避免很多线上错误

####  什么是ts

![74916904898](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749169048983.png)

1. 需要把ts代码转化为js代码才能运行

   `npm install typescript -g`安装转译模块,需要项目先有一个配置文件(用tsc --init),再运行tsc   

2. 通过vscode的插件完成ts文件的运行

   code-runner,如果是js文件,会直接采用node+文件名来执行,如果是ts,会使用ts-node来执行文件   

   `sudo npm install ts-node -g`

3. 通过代码构建工具去转化js,再去运行(webpack,rollup,esbuilde,vite),最终变成js执行 

   默认rollip,打包时会查找当前的目录,rollup.config.js

   ![74917231262](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749172312626.png)

   output要加一个sourcemap:true

####  ts的基础类型

ts关注的是类型,不是业务逻辑

> 变量名:类型=值

1. 基础类型
2. 高级类型
3. 内置类型
4. 自定义类型
5. 类型体操




ts的目的是在赋予的时候进行检查类型,只是提示作用,不是在运行的时候发生的,编译ts之后,类型就消失了(写空气)

 

ts比较偏向使用export{}来分割模块,而不是使用函数来分割

ts还有一个特点,不是所有变量都要添加类型,ts支持变量推导,可以猜测类型

 

```typescript
//基础类型,小写的命名
const name:string="jwt"
const age:number=1
const gender:boolen=true

//大写的用来描述实例类型
```

![74917527544](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749175275440.png)

类的类型,用来描述实例的

- 数组

```typescript
//类型[]Array类型,用于声明数组
let arr1:number[]=[1,2,3,2,]
let arr2:Array<number>=[1,2,3,1,3]
let arr3:(number|string)[]=["ss",1]
//联合类型
```

数组要求的是存储的格式安装特定类型来存储,不关心位置

- 元组

  你赋予的值要符合这个结构和顺序,不能无序去排列,也是不能增加额外的类型的值,只能是已有的,而且增加后无法访问

  ![74917630797](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749176307970.png)

  安全问题

  ​

- 枚举

  自带类型的对象（自己有类型,就是一个对象）

  约定一组格式我们会用枚举,状态码,权限,数据格式,标志位

  ![74943392254](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749433922540.png)

  ![74943399988](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749433999883.png)

  通过索引可以拿到枚举里面的值

   并且一般会给好这些东西一个固定的值

  ![74943415587](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749434155875.png)

   常量枚举不会额外编译成对象,更节约内存

  ![74943444037](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749434440374.png)

  异构枚举

  使用大写的命名,防止命名冲突

- null和undefiend

  也是一种基本类型,正常情况夏,只能赋予给unll和undefiend

   除非关闭相关的检测 

  void类型代表的类型是空,void用来表示函数的返回值,undefiend的范围小于void

  ​

  never是永远到不了的地方,就是never(函数无法到达结束的状态)

  ![74943557977](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749435579774.png)

   代码的完整性保护

  ![74943599357](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749435993571.png)

  不能让val到达validateCheck

  ​

  any是任何类型,能不写ant就不写,any会导致类型丧失检测,容易1导致出错,相当于没有ts加持




####  object引用类型

```typescript
function create(val:object){
    
}
//标识非原始数据类型,用object
```

![74943672983](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749436729830.png)

symbol和bigint

用的很少



> 基础类型
>
> string number boolean 数组 元组 枚举 null undefiend void never any object symbol bigint







####  ts的类型断言

![74943713487](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749437134875.png)

const声明的a1的类型是1,不可修改,而let的类型是number,范围会更大一点

**断言**

1. 指定类型再使用(非断言)

   ```typescript
   let strOrNum:string|number;
   strOrNum="1"
   strOrNum.endsWith()
   ```

2. 断言类型后再使用,as断言成某种类型(一定是联合类型中的某一个),!是非空断言,表示这个值一定不是空的

   不存在的结果自己承担,ts不管了

   ```typescript
   let strOrNum:string|number;
   (strOrNum! as string).charCodeAt(0);
   (<number>strOrNum!).toFixed(3)
   //这里有两个定义的方法
   ```

   使用! ? ??的使用方法

   !就是确定它存在

   ![74943900349](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749439003490.png)

   值 as xxx或者<xxx>值,一般用在联合类型

   ![74943919372](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749439193721.png)

   双重断言会破坏原有关系

   大类型断言成子类型

   ![74943932960](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749439329602.png)

   ​







####  函数类型

 函数的声明方法

```typescript
function sum(a,b){
    return a+b
}


const sum=function(a,b){
    return a+b
}
//一般使用第二种进行定义,因为function关键字声明的函数可以提升到作用域顶部,就是一般function用来描述外部函数,cosnt描述内部函数
```

 对于ts来说,函数关键字标注的函数,不能标注函数类型,通过表达式来声明函数,必须赋予的值满足定义的类型

![74957027039](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749570270393.png)

**可以使用type来声明关键字**

![74957058699](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749570586992.png)

如果标明函数的类型,在使用函数的时候以标明的为准

**参数**

![74957102517](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749571025173.png)

?不能和默认值一起使用

**参数this问题**

尽量不采用this,来作为函数的上下文,this的缺陷就是类型推到问题

如果要限制this的类型,需要手动指定this类型

![74957148412](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749571484124.png)

![74957205079](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749572050798.png)

ts有一个概念是重载,对于强类型语言可以写一个函数写多遍(参数不同),js实现重载考的是arguments

![74957217118](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749572171189.png)

类型重载后

![74957219976](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749572199760.png)

下面定义了父类型(不用any用重载前的那样也可以),上面的是一种具体的方案(子类型),上面可以限制下面的类型,根据用户给的值,来判断返回值

![74957231264](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749572312648.png)





####  类的类型

类可以继承

类的组成：构造函数，属性(实例属性，原型属性，静态属性)、方法（实例的方法，原型方法，静态方法），访问器，静态相关的配置

 ```typescript
class Circle{
    //给这个类来声明属性
    public x:number
    construct(x:number,y:number=200){
        this.x=x;
        this.y=y;
        this.fn=()=>{}
    }
    //public公开属性,类的实例可以在外部可以访问这个属性,类的内部也可以访问,继承的子类也可以访问
    //protected我自己能访问,子类也可以访问,外部不能访问
    //private私有的,自己才能访问
    //readonly是在初始化后(costructor之外)时候是只读的
}



//父子类的操作
class Animal{
    constructor(portected name:string){
        this.name='name'
    }
    //原型方法,就是每一个实例共享的方法,父类提供的方法,子类是可以进行方法重写的
    //void意味着是不关心函数的返回值，并不是空的意思
    changeName(value:string,age:number):void{
        this.name=value
    }
    //这个函数等同于加了一个public
    get aliasName(){
        return '$'+this.name;
    }
    //原型属性,需要通过访问器访问
    //原型是共享的,挂在protapyte上面实现
    set alisName(name:string){
        this.name=name;
    }
    static a=1;
    static getA(){
        return this.a
    }
    //静态方法声明
}


class Cat extends Animal{
    constructor(name:string,public rendonly age:number){
        super(name);
        //super在构造函数指向的是父类，在原型方法中调用的时候指向的是父类的原型
        //super在类中访问constructor\static函数中指向的都是父类,在原型方法中,属性访问器都是父类的原型
    }
    //子类重写父类方法要统一,赋予的函数要兼容父类,子类的参数可以比父类少
    //并且可以忽略void
    changeName(value:number){
         super.chageName(value )
         return 'abc'
    }
}

const tom=new Cat('tom',30)
//初始化
 ```



 不能被new的类

![74961944549](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749619445492.png)

ts有抽象类概念,abstract不存在的

抽象类可以含义非抽象的方法和属性,不会new它

![74962008164](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749620081641.png)

代码编程的时候,现在已经是慢慢脱离继承了,组合优于继承,类的装饰器redux\nest\mabx







####  接口和泛型的使用

接口：

和抽象类不同，接口必须都是抽象的，没有具体的实现

接口就是描述数据的结构或者形状，定义好结构，再去针对结构进行实现

**type和interface**

一般来说，描述对象、类，使用interface比较多,不能声明联合类型

type可以快速声明类型，比如联合类型，一些工具类型使用1type

type用的比较多，能type就用type，type不能重名,复杂类型使用type，interface可以重名

 ```typescript
//描述对象结构
interface IPerson{
    username:'abc'
}
let person:IPerson={
    username:'abc',
    age:38,
}


//子可以赋予给父亲，我们需要把一个值赋予给另外一个值，如果是声明的必须一致
let obj={
    username:'abc',
    age:38,
    address:'地址'
}
let person:IPerson=obj//赋值的时候会产生兼容性(儿子可以赋予给父亲)
 

//接口可以描述函数
interface ICounter{
    {}:number
    count:number
}
//const标识此值不能修改，let是可以修改的（如果给函数增加类型定义，函数不能被修改时只能用const）
const counter：ICount=()=>{
    return counter.count++
}
counter.count=0
 ```

//接口合并，自己的业务少

![74969126861](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749691268613.png)

会改变源码的逻辑 

也可以使用继承的逻辑来扩展接口

![74969169418](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749691694182.png)

![74969179266](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1749691792665.png)

任意类型和实例哪里的[Symbol()]:'ABC'是不用定义的



任意类型,随机的属性,描述数字索引(除了必要的属性,其他随意)

```typescript
interface IArray{
    [key:number]:any
}
let arr:IArray={
    0:1,
    1:2,
    3:3,
    4:"abc",
    5:true 
}
//或者
let arr:IArray=[1,2,3]
```





通过索引访问符,来获取内部类型

```typescript
interface ResponseData{
    username:string,
        token:string
}

interface ReturnVal{
    code:number,
        data:ResponseData
}

type ICode=ReturVal['code'];
type IUsername=ReturVal['data']['username'];
type IKey = ReturnVal[teyof ReturnVal]
//interface没有联合类型
```

接口可以实现,接口通过类实现

一个接口可以继承多个接口,接口可以用于继承类

![75024471539](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750244715390.png)

  



怎么标识自己传递了一个类,类类型不能描述类本身,描述的是实例

```typescript
class Dog{
    
}

class Cat{
    
}




function createInstance<T>(clazz:{new:()T}){
    //不能直接使用dog为类型,也不能typeof dog来作为类型
    //ts的检验类型,鸭子类型检测,有大部分相同的就是一个类型,两个空类,是一样的类型
    //所以这里要使用泛型来完成,传什么就是什么类型
    //或者写成new ()=>Dog
}

const instance = createInstance<Dog>(Dog)
//省略写法
const instance = createInstance(Dog)

```

**泛型:**类似坑位(函数的形式参数,刚开始不知道类型,通过使用才知道类型)

使用才传递东西

通过<>传到上面,上面用<>接收,可以一直传递

![75024632513](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750246325132.png)

也可以这样写

有参数,且接口或者type使用泛型的情况



泛型可以用于函数接口类这些东西上面

无法确定类型的时候使用

**实际使用例子**

![75024706093](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750247060934.png)

写辅助函数的时候,有多个泛型可以用于保存值

![75024731008](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750247310081.png)

弱推导类型不能发现数组交换了,所以要用泛型来





区分在使用函数的时候定义泛型还是在定于函数的时候使用泛型

![75024811269](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750248112693.png)

![75024846365](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750248463652.png)

声明时准备和使用时准备







####  泛型和交叉类型

```typescript
//接口的返回类型可能是统一的
//code,data,message
//泛型的默认值来解决泛型的值默认情况
interface APIResponse<T>{
    error:number,
    data:T,
    message:string,
}

    interface LoginInfo{
        username:string,
            token:string
    }
    
    function login():APIResponse<LoginInfo>{
        return{
            error:1,
            data:{
                username:"sss",
                token:'sssssssss'
            },
            message:'成功'
        }
    }
let r = login()

//在开发的时候想使用联合类型
type IUnion<T=Boolen>=T|string|number
type t1=IUnion
type t2=IUnion<string[]|number[]>

//用户随意传递
//跟上面那个类似
           
//使用泛型的时候不能做运算（T+T=T？）
       
```

```typescript
function getVal<T extend string|number>(val:T):T{
                return val
                }
                
//约束类型T是什么什么的子类型
```

function getLen<T entends {length:number}>(val:T){

​        return val.length

}

传入的东西要有length属性

![75029744487](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750297444878.png)

```typescript
//类中泛型
class MyList<T extend sting|number>{
    public arr:T[]=[]
    add(val:T){
        this.arr.push(val)
    }
    getMax():T{
        let max = this.arr[0]
        for(let i = 0;i<this.arr.length;i++){
            let cur=this.arr[i];
            cur>max?(max=cur):void 0
        }
        return max
    }
}
const list = new Mylistconst 

const list = new Mylistconst<string>
list.add(1)
list.getMax()
```

泛型使用场景:函数(参数返回值),对象坑位,类,泛型的默认值和约束







####  交叉类型

聚合类型是并集,交叉类型是交集

```typescript
interface Person1 {
handsome:string
}
interface Person2 {
high:string
}
// let person:Person1 | Person2 ={
I1high:'高
}
let person:Person1 & Person2 ={handsome:"帅"high:'高}

//使用交叉类型需要全部都使用 
//子类型可以赋予给父类型,子类型的结构要有包含父类型的结构
```





####  内置类型

- 条件类型 if/else 三元表达式 (extends 左边和右边的关系)

  子类型 extends 父类型 = true

  ```typescript
  type StatusCode<T>= T extend 200|201?'success':'fail'
  type IReturnMessage = StatusCode<300>

  type IObj<T>= T extend {name:'jw'}?'ok':'no'
  type IRerson = IObj<{name:'jw',age:30}>
  ```

- 类型级别:1. 根据结构的角度分析,2. 从类型角度来分析

  never是任何类型的子类型

  ![75047497514](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750474975143.png)

    {}看成结构，object看成复杂类型，而Object就是万物皆对象

    T4是false，其他是true

- ​


```typescript
type T7= any extends unknown?true:false
type T8= unknown extend any?true:false
  //两个都是true，不区分彼此



type T9 = any extends 1?true:false
//条件类型的分发机制，1+除了1的部分，返回ture|false
type T10<T> = T extends 1?true:false
type Temp10 = T10<T>
 //返回type Temp10 = never,出现问题
//any是自带分发机制的
  //通过泛型传递的never相关机制,会产生分发机制,返回never

  //联合类型的子类型,是联合类型里的某个类型
  type T11 = 100 extends 100|200 ?true:false

  //通过条件类型来进行类型的区分,条件语句也可以实现约束的效果
  interface Fish{
      name:'鱼'
  }

  interface Water{
      name:'天'
  }

  type GetType<T extends Fish|Bird>=T extends Fish?Water:Sky
  type A1 = GetType<Bird>
  //鱼返回水,鸟返回天空


  type GetType<T extends Fish|Bird>=[T] extends [Fish]?Water:Sky
  //分发导致的问题:什么时候会有分发
  //1.联合类型通过泛型传递,泛型左边会被分发
  //2.而且比较(extends)的时候会产生分发
  //3.类型需要满足裸类型(裸类型就是泛型,就是自己没有和别人搭配)



  //有的场景不需要分发机制,要禁用,有的场景需要分发机制做判断
  type <T>=T&{}
  type UnionAssets<T,K>=Nodistribute<T> extends K?true : false
  type U1 = UnionAssets<1|2,1|2|3>
  type U2 = UnionAssets<1|2|3，1|2>
   //判断类似是否完全一致
  type isRqual<T,K,S,F>=T extends K?K extends T ? S:F:F
  type A2 = isREqual<1|2,1|2,true,false>
                        //分发导致返回bollen
                        把extens前的改成NoDistribute<T>
```

  内置类型里面有很多基于条件类型的类型

比如Extract Exclude

```typescript
type ExtractBes = Extract<1|2|3|4,1|2>
//实现方法
type Exclude<T,U>=T extends U?never :T
```



- 对象类型

  ```typescript
  interface Person1{
      handsome:string;
  }
  interface Person2{
      high:string;
  }
  type Computr<T>={
      [key in keyof T]:T[key]
  }
  //ketof any
  //返回number|srting|symbol,使用的比较多
  //keyof unknow
  //返回never

  type Person3=Compute<Person1 & Person2>//合并类型
  ```

  ![75048299189](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750482991891.png)

  递归发现对象内部属性

  ​



####  inference的使用

- 类型推断inference

  infer关键字只能用于条件类型中，用来提取类型的某一部分的类型，放在不同的位置，就可以帮我们取不同位置的类型

  ```typescript
  function getUser(name:string,age:number){
      return {name:age,address:{}}
  }

  type ReturnType<T>=T extends (...args
  :any[])=>infer R?"R":never
  //用infer推到出返回值类型R

  type T1 = ReturnType<typeof getUser>

  ```

- ReturnType和Parameters和InstanceType是内置类型

  实现

  ![75063656066](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750636560663.png)

  函数和构造函数的区别是new了一下啊

- 推断数组

  ```typescript
  type TailToHead<T extends any[]> = T extends [...infer C,infer B]?[B,...C]:any

  type x = TailToHead<["jw",30,40]>
  
  //将元组转化为联合类型
  type ElementOf<T>=T extends Array<number>?"R":any
  type TupleToUnion = ElementOf<string,numer,boolean>
  //转化成了string|munber|bollean
  
  
  //深入推断
  type Promise<T>= T extends Promise<infer V>?V:any
  type PromiseReturnValue = Promise<Promise<number>>
  
  //深入拆包
    type Promise<T>= T extends Promise<infer V>?PromiseV<V>:T
  type PromiseReturnValue = Promise<Promise<number>>
  //infer就是推导条件中的某个部分
  ```






####  内置类型

exclude extract nonnullable infer returntype paramters instancetype

集合、条件来操作的



- 基于对象操作的映射类型

```typescript
type A1 = {name:string}
type A2 = {age:number}

type Compute<T extends object>={
    [K in ketof T]:T[K]
    //映射类型
}
type A1&A2 = Compute<A1 & A2>;
//拿到合并后的类型

//对对象进行修饰操作的内置类型(必选,可选,只读)

interface Company{
    num:number;
    name:string
}
interface Person<T>{
    name:string;
    age:number;
    company:T
}
type withCompany = Partial<Person<Company>>
//Partial让所有类型都变成是可选的,只有第一层是可选的
//实现
type Partical<T>={
    [L in keyof T]?:T[L]
}
type Required<T>={
    [L in keyof T]?:T[L]
}
//必填所有项
//多层处理
type DeepPartial<T>={
    [K in keyof T]?:T[K] extends object ?DeepPartial<T[K]>:T[K]
}
type DeepRequired<T>={
    [K in keyof T]-?:T[K] extends object ?DeepRequired<T[K]>:T[K]
}

type S = Readonly<Person<Company>>
//添加只读


//pick 挑选
//omit 去掉某些属性
//exclude  extract是对集合操作
type PIckPerson = Pick<Person, "name"|"gae">
//实现
type Pick<T,K extends keyof T>{
    [P in K]:T[P]
}

type PickPerson = Omit<Person,"name"|"age">
//实现
type Omit<T,K extends keyof any>=Pick<T,Exclude<keyof T,K>>

//Record 取代object，告诉函数，这一定是一个非空对象
```

![75064944782](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750649447828.png)

1. 根据传入的值进行类型推到，name和age会赋予给K，value会赋予给V
2. 拿到callback返回值，，使用泛型进行类型推到它的返回值，R根据返回值进行推导
3. 映射成一个新的record，由K和R组成为结果数组
4. 遍历obj，去运行回调函数





####  兼容性和类型推导

ts兼容性分为两种：

- 子extends父

- 结构来考虑

```typescript
  let str:string="abc"
  let obj!:{toString():string}
  obj=str
  //结构来考虑

  //函数兼容性
  let sum1=(a:number,b:number)=>a+b
  let sum2=(a:number)=>a

  type Sum1 = typeof sum1
  type Sum2 = typeof sum2
  type x = Sum2 extends Sum1?true:false
  //对应函数而言,它的兼容性,少的可以赋予给多的,参数少的是子类型
  //返回值要求安全,返回值要求是子类型
  //类的兼容性,也是一样的,比较的是实例

  class A{
      a=1;
  }

  class B{
      a=1
  }

  const b:B=new A{}

  //可以赋值

  //如果类中的属性,有private或者protected则有两个值不能互相赋值

  //差异化基本类型

  class AddType<K>{
      private type!:K
  }

  type withType<T,K>=T AddType<K>

  type BTC=withType<number,"BTC">

  type USDT=withType<number,"B TC">

```
  ts主要考虑安全性,安全就可以复制

```typescript
//逆变（在函数参数可以标记儿子传父亲）和协变（可以标记父亲返回儿子）
class Parent{
    car(){}
}
class Chile extends Parent{
    house(){}
}
class Gradson extends Child{
    sleep(){}
} 

function fn(callback:(ctr:Child)=>Child){
    callback(new Child())
}
fn((child:Parent):Child=>{
    return new Child()
})//逆变
//内部调用函数的时候，可以传递Child和Grandson,但是在使用属性时，只能认为最多就是Child
//fn 内部调用 callback 时，传入的实际参数是 Child 类型。
//若回调函数声明参数为 Parent（父类型），则它可以安全接受 Child（子类型），因为子类拥有父类的所有属性和方法（Parent 的要求已被满足）。
//函数的返回值，需要返回子表，因为内部代码在访问属性的时候要保证可以访问到
//协变
//当函数 A 赋值给函数 B 时，A 的返回值类型必须是 B 返回值类型的子类型（或相同）。


interface MyArray<T>{
    concat(...args:T[]):T[]
}
//这种写法不进行逆变检测，所有描述对象中的方法时全部采用这种方式

```

参数逆变，当函数 A 赋值给函数 B 时，A 的参数类型必须是 B 参数类型的**父类型**（或相同）。

返回值协变，当函数 A 赋值给函数 B 时，A 的返回值类型必须是 B 返回值类型的**子类型**（或相同）。





####  枚举

```typescript
enum E1{}
enum E2{}

let e1:E1
let e2:E2
//两个枚举不能互相赋值,不能兼容

//泛型兼容性,如果生成结果一致,类型就兼容
type II<T>={name?:T}
type X1=II<string>extends II<string>?true:false
//生成结果一致,返回true
```

![75073018135](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1750730181355.png)

```typescript
//对象兼容性,多的属性可以赋予给少的
//类型层级兼容性,never->字面量->基础类型->包装类型->any/unknown
//子extends父,满足即可赋值
```

类型推到的概念

1. 赋值推断
2. 函数时通过左边来推导右边,基于上下文类型来进行自动的推导
3. 函数返回值标记成void,赋予一个函数的时候,意味不关心







####  类型保护

 基于js＋ts（收窄）

ts很多情况下，需要使用联合类型，默认情况下只能使用公共的方法，识别类型（针对某个类型进行处理）

```typescript
function  fn(a:string|number){
    if(typeof a === "string"){
        a;
    }else{
        a;
    }
}
```

![75101423122](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751014231222.png)

收窄了

> typeof 基础类型
>
> instanceof 类类型
>
> in 可辨识类型

```typescript
interface Bird{
    kind:"鸟";
    fly:string
}
interface Fish{
    kind:"鱼";
    swim:string
}

function getAima(val:Bird|Fish){
    //基于差异化来辨别
    if("fly"in val){
        val
    }else{
        val
    }
    if(val.kind=="鸟"){
        val.fly
    }else{
        val.swim
    }
}
//通过各种判断来缩小范围,生命周期{}

```

**函数**嵌套不识别的问题使用?或者!\if都可以做到类型保护

!一定存在,?取值但不能用于赋值

![75101565822](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751015658229.png)

ts无法识别的时候,一定要用断言

> function isBird(val:Bird|Fish):val is Bird{
>
> //函数的名字和返回值无关
>
> //ts的返回值类型
>
> //true是bird还是false是bird
>
> return  "fly" in val
>
> }





####  自定义类型

unknown和any都是顶级类型,任何类型都能赋予给他

> type unionUnkown = unknow | string | true | false
>
> type sss : unkown = "ssss"可以进行赋值

如果无法确定类型，不要贸然采用any

any不校验，意味着可以任意调用和取值

但unkown是any的安全类型

如果把类型标识为unkown类型，必须先类型保护，再去使用类型

```typescript
//自己去实现一些类型
//内置类型
//基于条件类型的 Extract Exclude 集合类型
                //基于映射的类型 Paetial Requried Readonly 修饰
                //结构的 Pick Omit Record 结构处理
                //基于推断的类型 instanceType returnType Paramtes infer 模式类型匹配
//merge类型会合并在一起

```

```typescript
//对象求交集
type ObjectInter<T extend object,K extends object>=Pick<T,ketof T & keyof K>
type X1 = ObjectInter<T1,T2>

//对象求查
type ObjectOff<T extend object,K extends object>=Omit<T,keyof K>
type X2 = ObjectOff<T1,T2>

//并集就是联合类型

//两个类型合并在一起,如果都有的属性,用T2的? {...T1,...T2} 
type OverWrite<T extends object,K extends object>=ObjectDiff<T,K>&ObjectDiff<K,T>&ObjectInter<K,T>
type X3=OverWrite<T1,T2>


//merge
type MergeType<T,K>={
    [K in keyof T]:K extends keyof U?T[K]|U[K]:T[K]
}
type MergeWrite<T extends object,K extends object>=ObjectDiff<T,K>&ObjectDiff<K,T>&MergeType<K,T>
type X4=Compute<MergeWrite<T1,T2>>
```



- 命名空间----隔离同一个文件

  namespace

  import export 外部模块，就是隔离多个不同的文件

  ![75110902009](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751109020095.png)

  可以声明重复名字的变量

  给他们划分了两个对象，导出然后访问（只有导出了才能访问）

  扩展类，函数，枚举扩展的命名空间必须写在这些声明下面

  ![75110961298](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751109612982.png)

  用的不多，基本没什么用

  声明文件，装包，拆包，类型体操

  基础类型（基础类型 string number boolean null unfiened，void）

  底端类型，数组，元组，枚举

  包装类型

  any unknown

  联合类型，交叉类型，断言 as ! ?

  ​

  //函数：（签名） 参数，返回值，可选的？默认值=剩余参数。。。this的问题 函数的重载（类型的重载）

  逆变和协变

  类 ：修饰符 private protexted pulic readonly (private constructor)子类重写方法的问题

  类型兼容 原型方法和实例方法的声明。

  抽象类，类方法，构造函数类型

  ​

  接口：可选？ readonly任意类型，可索引接口 extends implements

  type 和 interface的区别

  ​

  泛型：占位置的对象，使用时传递类型

  条件类型：子extends 父 映射类型（子：子类型，子结构。父：父类型，父结构）

  分发的问题，有好有坏，注意禁用和使用

- extract exclude

  从类型 `T` 中**提取**能赋值给类型 `U` 的部分（保留交集）extract

  从类型 `T` 中**排除**能赋值给类型 `U` 的部分（取差集）exclude

- partial readonly

- pick omit record

  pick**Pick<T, K>**
  从类型 `T` 中**挑选**指定的属性键 `K`（`K` 需是 `T` 的键的子集）

  **Omit<T, K>**
  从类型 `T` 中**剔除**指定的属性键 `K`（反向 `Pick`）

  **Record<K, T>**
  构造一个键为类型 `K`，值为类型 `T` 的**新对象类型**

  infer推导的作用

- ​

- ​

  ​

  ​

  兼容性

  内置的类型推导，写了可以直接推导类型，函数提供上下文可以推导类型，返回值也可以推导

  类型保护typeof instanceof in 可辩别类型 is

  一种缩小类型的方式

  typeof看基本类型

  instanceof看某个类或者构造类型的实例

  in检测对象是否有某个属性

  ​

  重映射+自定义类型




#### 命名空间-外部模块

es6模块，需要转化成不同的模块方式 ，commonjs规范（require，exports），amd模块(define),导出amd需要配置file入口(在moduleResolution,的值是Node在tsconfig.json哪里)

支持直接进行转化

1. 如果要使用commonjs需要安装 `pnpm i --save-dev @type/node `

   commonjs规范可以转化为amd

2. target指代的是语法

   modele导出的模块规范是什么

   ​

3. ts里面,为了做到commonjs和amd的互转,发明了一种写法

   使用eport = / import = 导入,使用这种写法可以打包关联文件,我们直接用es6就好了,import基本能解决require的问题

   ![75119332479](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751193324791.png)

   一些规范的使用场景

4. 最终编写的代码,需要转为js去使用

   需要给打包的结果添加类型,就是.d.ts

5. 打包后只生产js，没有类型了，可以开启配置中自动生产类型文件

   没有ts类型的的外置包，需要添加ts类型

   或者使用cdn外链的库，也需要写声明文件

   扩展全局组件的时候，也需要写声明文件了

   特殊文件的处理也需要声明文件

   为了我们在写代码的时候不会出错，丢失类型

   声明文件的语法declare用于声明类型（全局声明），全局声明文件就算.d.ts

6. ![75119571289](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1751195712890.png)

   声明文件

   ​






####  声明文件

如果有人已经给这个模块写好了类型，我们就可以直接引入人家的操作

如果我们在使用某些包的时候，别人提供了类型文件我们可以直接安装使用







查找声明文件

先看自己是否定义过用自己的,没有就向上找到node_modules 

1. 默认找当前node_MODULES下的同名模块，看是否有此文件夹，有的话找package.json>types
2. 如果没有types默认找这个模块下的index.ts
3. 查找node_modules下的@types是否有同名模块,按照上述方式继续查找

> @type/xxx是声明文件等
>
> 默认可以指定导入模块的声明文件路径 

![75216399009](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752163990098.png)

命名空间全局化

![75216411353](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752164113534.png)

三斜杠指令

通常可以在模块扩展上使用

- 直接扩展

![75216428686](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752164286864.png)

子模块,扩展了父模块

![75216432426](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752164324262.png)

父模块

扩展后让这些模块可以变多功能

![75216447096](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752164470961.png)

使用

> 注意文件命名的关系



- 针对特定的接口扩展

  父模块

  ![75216538754](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752165387542.png)

  子模块

  ![75216540257](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752165402574.png)

  这样避免了上面那一种全局污染的可能性,针对特定接口进行扩展 

- 三斜线指令可以通过types来引入其他的声明文件（第三方）

- 自己的用path来引入

- lib可以加载内置类型




####  拆包和装包

> 命名空间namespace就是一个对象

vue3的ref他就是一个包装对象

给一个对象,对他的取值操作和设置操作进行重写了

```typescript
interface Proxy<V>{
    get():V;
    set(value:any):void;
}

type Proxify<T> ={
    [K in keyof T]:Proxy<T[K]>
}
const props = {
    nameL"jwt",
    age:30
}
function proxify<T>(obj:T):Proxify<T>{
    const result = () as Proxify<T>;
    for(let key in obj)={
        let value = obj[key]
        result[key]={
        get(){
        return value
    },
        set(val){
        value=val
    }
    }
    }
    return result
}
const proxy = proxify(props)

proxy.name.get()
proxy.name.set("abc")

export {}
```

使用泛型进行装包



拆包

```typescript
function unProxifiy<T>(proxy:Proxif<T>):T{
    let result=()as T
    for(let Key in proxy){
        result[key]=proxy[key].get();
    }
    return result
}
```

进行拆包,返回原始的那个东西

> 装包也能用于axios的封装和处理

```typescript
export type ResponseData<T = any>={
    code:number
    data?:T
    message?:string
}

class HttpRequest{
    public timeout=3000
    public loadingMaping:Record<string,string>={}
    public baseURL= import.meta.url==="development"?"https://localhost:3000:"/"
    
    public setInterceptor(instance:AxiosRequestConfig){
        return {..option,baseURL:this.baseURL,timeout:this.tomeout
    }
    
    public setInterceptor(instance:AxiosInstance){
        instance.interceptor.requset.use(
        (config)=>{
            return config
        },
        (err)=>{
            return Promise.reject()
        }
        )
    }
    public request(option:AxiosRequestConfig){
        const instance=axios.create()
        options=this.mergeConfig(options)
        this.setInterceptor(instance)
        return instance(options)
    }
    public get<T>(url:string,data:any):Promise<ResponseData<T>>{
    this.request({
        url,
        method:"get"
        data
    }).then(res=>{
        return Promise.resolve(res.data)
    }
    }
}
const http = new HttpRequest
http.get<{token:string}>("/login",{}).then((res)=>{
    res.data?.token
})
```





####  类型体操

> 可以通过索引来访问一个对象\数组对应的值

```typescript
type LengthOfTuple<T extends any[]>=T['length']
type A = LengthOfTuple<["B","F","E"]>//3
```

> 元组自带索引和长度,可以用索引和长度

```typescript
type FirstItem<T extends any[]>=T[0]
type A = FirstItem<["B","F","E"]>//B
```

![75221877155](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752218771558.png)

拿到最后一个

![75221893786](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752218937862.png)

去除第一个

![75221974469](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752219744698.png)

链表反转

```typescript
//拍平
type Flat<T extends any[]>=T extends [infer L,...infer R]?
    //每次判断一下是不是数组，是就去递归它
    [: (L extends any[]?Flat<L>:[L])]:T
type A = Flat<[1,2,3]>
```

 ![75222095927](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752220959272.png)

判断是否满足某个条件，构建数组

```typescript
type Filter<T extend any[],K,F extends any[]=[]>=T extends [infer L,...infer R]?Filter<R,K,L&{} extends K?[...F,L]:F>:F

type A=Filter<[1,2,3,true],nmuber>//[1,2,3]
```
找到相同的值,不仅要value一样,key也要

![75225143215](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752251432157.png)

![75225118298](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752251182985.png)

类型体操多在源码里面出现，了解源码再去考虑类型，不然很难看懂



####  模板字符串

我们可以基于字符串类型，来创建新的类型，模板字符串

```typescript
type name = 'js'
type handsomeName = `handsome${name}`

```

当我们传入应该联合类型的时候,也有分发机制

```typescript
type MarginOrPading='padding'|'margin'
type Dirction='felt'|'right'|'top'
type Compose=`${MarginOrpading}-${Dirction}`
```

可以基于这种方式编写一些特定类型,scss变量

//element-plus(red-100 red-200 red-300)

![75228820877](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752288208778.png)

可以传入字面量类型,也可以传入基础类型

```typescript
//可以对这些进行一些特殊的操作
type A1= Capitalize<I1>//开头大写
type A2= Uncapitalize<A2>//开头小写
type A3= Unprecase<A1>//全部大写
type A4= LowerCase<A3>//全部小写

```

![75228938221](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752289382216.png)

![75229115424](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752291154244.png)

infer的使用，可以提取第一个字符，`${infer L}`

如果有两个infer,最后面的就是所有的

如果有符号分隔,那就会分左右,没有默认左边的都是一个字符串







####  装饰器

平时用不到，是基于类的，是一个实验性语法，mbx，nestjs主要是靠装饰器实现的

后续的具体实现可能会发生变化，但目前是比较稳定的一种

本质就是应该函数，只能在类和类的成员上使用

ts装饰器：类装饰器，（静态丰富，原型方法）方法装饰器，（静态属性，静态方法装饰器）属性装饰器，（get，set）访问符号装饰器，（原型函数，构造函数）参数装饰器

> 要在tsconfig.json打开这个功能

- 类装饰器

```typescript
function ClassDecorator(target:any){
    target.type="动物"
    target.getType=function(){
        return this.type;
    }
    target.protype.eat=function(){
        console.log('eat')
    }
}


@其他装饰器,也可以给其他类加东西
@ClassDecorator
class Animal{}

console.log((Animal as any).type)

```

- 方法装饰器

```typescript
function Enum(isEnm:boolean){
    return function(target:any,key:string,descriptot:PropertyDescriptor){
        //装饰器函数
        //类的原型和类的名字
        //属性访问器
        descriptor.enumerable=isEnm
        let original=descriptor.value
        descriptor.value=function(...args:any){
            //写一些函数逻辑
            original.call(this,...args)
            //切片和声明相关的东西
        }
    }
}

class Animal{
    @Enum(true)
    //装饰器写在这里
    eat(){}
}
const animal = new Animal()
//对类和方法进行装饰是有意义的,对其他就没有什么意义了
```

- 属性装饰器

  如果编译成esNext会编译成

  ![75233520201](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752335202016.png)

  使用ES2015编译成,默认此值就在函数上

  ![75233523060](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752335230604.png)

  在new时候才会赋值,并不会在最开始即在函数上

- 访问符装饰器

  ```typescript
  class Animal {
      private _value: string = '';
      
      @ToUpperCase('PREFIX_') // 添加前缀参数
      get value() {
          return this._value;
      }
      
      set value(newVal: string) { // 注意这里应该和getter同名
          this._value = newVal;
      }
  }

  function ToUpperCase(prefix: string) {
      return function(target: any, key: string, descriptor: PropertyDescriptor) {
          // 保存原始的setter
          const originalSet = descriptor.set;
          
          // 修改setter
          if (originalSet) {
              descriptor.set = function(newVal: string) {
                  // 调用原始setter，但先处理值
                  originalSet.call(this, prefix + newVal.toUpperCase());
              };
          }
      };
  }

  // 使用示例
  const animal = new Animal();
  animal.value = 'test'; // 实际会存储为 "PREFIX_TEST"
  console.log(animal.value); // 输出 "PREFIX_TEST"
  ```

- 参数装饰器

  ![75239803739](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752398037397.png)

  ![75239805226](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752398052264.png)

  只能拿到类,key和索引

  构造函数种的参数装饰器指代的是类本身,key是undefined

- 装饰器的执行顺序

   ![75239888960](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752398889604.png)

  ![75239891991](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752398919918.png)

  生成顺序

  ![75239894407](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752398944077.png)

  对于实例来说，先走参数，再走对应的方法（构造函数），再到静态属性，最后是自己的类

  装饰器的使用主要是切片，在类的过程里面进行添加和修改删除等操作

 

**反射元数据**

反射就是修改程序执行的时候的行为

元数据：描述数据的数据

修改执行代码的行为，对数据进行描述

 ![75247141927](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752471419272.png)

安装这个包，然后使用这个函数进行相关的描述

![75247163861](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752471638615.png)

有很多相关的扩展和处理

> 命令式和声明式的

使用声明式的装饰器来写

![75247211775](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752472117751.png)

会更加清晰一点

使用了装饰器



- 控制反转,失去了控制权

- 控制正转，整个控制的过程都是我自己完成的

- 依赖注入

  通过依赖注入让代码变得灵活，而不是四班，他是oci的具体实现。

```typescript
  // 首先定义必要的装饰器和容器
  const container = new Map();

  // Provide 装饰器 - 用于标记可被注入的类
  function Provide(identifier?: string): ClassDecorator {
    return (target: any) => {
      const id = identifier || target.name;
      container.set(id, target);
    };
  }

  // Inject 装饰器 - 用于标记需要注入的属性
  function Inject(identifier: string): PropertyDecorator {
    return (target: any, propertyKey: string | symbol) => {
      const dependency = container.get(identifier);
      if (dependency) {
        target[propertyKey] = new dependency();
      }
    };
  }
  //装饰器实现依赖注入 



  // 定义接口
  interface Monitor {
    display(): void;
  }

  interface Host {
    start(): void;
  }

  // 实现类
  @Provide("Monitor")
  class Monitor23inch implements Monitor {
    display() {
      console.log("23英寸显示器已开启");
    }
  }

  @Provide("Host")
  class AppleHost implements Host {
    start() {
      console.log("苹果主机启动");
    }
  }

  // 计算机类使用依赖注入
  class Computer {
    @Inject("Monitor")
    public monitor: Monitor;

    @Inject("Host")
    public host: Host;
    
    bootstrap() {
      this.monitor.display();
      this.host.start();
    }
  }

  // 使用示例
  const myComputer = new Computer();
  myComputer.bootstrap();

  1. **解耦**：类不直接依赖具体实现，而是依赖抽象
  2. **可测试性**：可以轻松替换依赖进行测试
  3. **可维护性**：修改实现时不需要修改使用方代码
  4. **可扩展性**：方便添加新的实现

  这个容器可以描述成

  ![75273835513](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1752738355133.png)

  可以绑定实例,自动解析依赖

  可以用Map表替代
```


####  tsconfig.json

```json
{
    'target':"es5",//根据最后的打包文件环境来设置,停入对应的声明文件
    "lib":[],//打包的时候引入的类型声明信息
    "jsx":"react",//jsx环境的配置,preserve是jsx不转jsx,react-jsx是新版默认值
    "experimentaDecorators":true,//启用装饰器语法
    "enitDecoratorMetadata":true,//发射元数据信息,默认使用装饰器都i会开启这个值
    "jsxFactory":"h",//为了preact提供
    "jsFragmentFatory":"Fragment",//针对Fragement实现不同的功能
    "jsxImportSource":"",//更改导入路径
    "moduleDetection":"force"//当前文件如果有import export就是模块,如果没有就是全局,强制实现
}
```





####  模块相关配置

ts声明文件

“module”可以选择CommonJS或者es6、es2020、esnext、nodenext、AMD、UMD、systemjs（微前端）

 ![75332151375](C:\Users\zxh\Desktop\前端\typescript\ts.assets\1753321513751.png)

> "moduleResolution":"Classic"经典的解决方法,这里也可以指定node版本,导入模块的时候,有package.json(main,module,exports)
>
> 使用更高级的node来支持更多新语法
>
> 
>
> 

 



















































