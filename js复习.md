#  js复习

####  类

- 构造函数：是专门用于初始化新对象的函数，要使用new关键字调用，来创建新对象，因此构造函数本身只需要初始化新对象的状态，构造函数的prototype属性将会被用作新对象的原型

  ```javascript
  //用来初始化新的Range对象的构造函数
  //它不创建或者返回新对象，只是初始化this
  function Range(from,to){
      //构造函数一般大写开头，提示你
      this.from=from;
      this.to=to;
  }

  //所有Range对象都继承这个对象
  Range.prototype={
      //如果这个x在范围里面就返回true，否则返回false
      //适用于文本、日期和数值
      includes:function(x){
          return this.from<=x&&x<=this.to
      },
      
      toString:function(){
          return "("+this.from+"..."+this,to+")"
      }
  }

  let r = new Range(1,3)
  r.includers(2)//true
  r.toString()//(1...3)
  ```





- 构造函数、类标识和instanceof

  1. r instanceof Range //=>true,说明r继承了Range.prototype

     左边是要检查的对象,右边是代表某个类的构造函数

  2. ![74471345815](C:\Users\zxh\Desktop\前端\js复习.assets\1744713458157.png)

     let o = new F();//创建类F的对象o

     o.constructor===F//true,constructor属性指定类

  ​

####  class

```javascript

class Range(from,to){
    //构造函数一般大写开头，提示你
    //这些属性是不继承的,是当前对象独有的
    constructor(from,to){
    this.from=from;
    this.to=to;
    }
    
    
    //所有Range对象都继承这个对象

    //如果这个x在范围里面就返回true，否则返回false
    //适用于文本、日期和数值
    includes:function(x){
        return this.from<=x&&x<=this.to
    },
    
    toString:function(){
        return "("+this.from+"..."+this,to+")"
    }
}

//使用示例
let r = new Range(1,3)
r.includers(2)//true
r.toString()//(1...3)
```

·类是以class关键字声明的，后面跟着类名和花括号中的类体

类体包含使用对象字面量方法简写形式（示例9-1中使用过）定义的方法，因此省略了function关键字。但与对象字面量不同的是方法之间没有逗号（尽管类体与对象字面量表面上相似，但它们不是一回事。特别地，类体中不支持名/值对形式的属性定义）。

```javascript
//创建一个子类,类变量也不会提升到文件顶部
class Span extends Range{
    constructor(start,length){
        if(length>0){
            super(start,start+length)
        }else{
            super(start+length,start)
        }
    }
}
```





####  静态方法

在class体中，把static关键字放在方法声明前面可以定义静态方法。静态方法是作为构造函数而非原型对象的属性定义的。

```javascript
static parse(s){
    xxxx
}
//写在Range里面的静态方法
//通过Range进行调用
let r = Range.parse("(1....10)")

```





- 公有\私有和静态片段

  ```javascript
  class Buffer{
      #size=0;
      //私有字段
      get size(){return this.#size;}
  }
  ```



####  为已有的类添加方法

```javascript
Complex.prototype.conj=function(){
    return new COmplex(this.r,-this.r)
}

//为complex添加了一个计算共轭复数的方法

```



####  子类



