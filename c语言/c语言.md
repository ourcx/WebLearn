

# c语言

#### 初始c语言

![image.png](https://s2.loli.net/2025/03/28/XdVwIlPEBHtCGzf.png)

编辑器：编辑功能

编译器：编译功能->c/c++是编译型语言 

#### 第一个c语言

> c语言一定要有main函数,为主函数:

```c
#include <stdio.h>
//std-标准 i-input o-output
//用来告诉编译器,使用了库函数
int main()
    //main函数是程序的入口
{
    printf("hello\n")
    //printf是一个库函数
    return 0
    //如果正常运行,就返回0,不正常就返回非零数值
}
//int是整型类型的意思

//快捷键:运行->ctrl+fn+f5
```

#### 数据类型

> char字符串数据类型-一个字节
>
> short短整型 -两个字节
>
> int整形-四个字节
>
> long长整型-四个字节
>
> sizeof(long)>=sizeof(int)
>
> long long更长的整型-八个字节
>
> float单精度浮点数-四个字节
>
> double双精度浮点数-八个字节
>
> 使用类型的原因是向内存申请一个空间

- 计算机单位:
  - bit:比特位,最小,八个比特位是一个字节
  - byte:字节,1024个字节是一个kb
  - kb
  - mb
  - gb
  - tb
  - pb

```c
char ch ="w"
int weight=120
```



#### 常量和变量

常量tm就是不变的量

```c
int b =20
    //全局变量

extern int a;
//声明来自外部的符号

void text()
{
    printf("text-->a",a)
}
    
int main()
{
    int a = 100
    printf("a=%d",a)
        //局部变量,局部和全局冲突的时候,局部优先
    scanf("%d %d",&num1,&num2)
    //这个方法要输入地址,可能会报错这个函数不安全
    return 0
}
//不要将全局和局部变量的名字写出一样的
```

- 变量的作用域和生命周期

  - 局部变量的作用域是变量所在的局部范围,全局变量的作用域是整个工程

    限定整个名字可用性的是变量所在的作用域

  - 生命周期:局部变量的生命周期:进入作用域开始,出作用域结束.全局变量的生命周期:整个程序的生命周期

  - 作用域和生命周期是有点关系的

- 常量：描述不变的量

  - 字面常量

    > a = 10

  - cons定义的常变量

    > const int a = 10本质上是一个变量，但不能直接被修改，有常量的属性

  - #define定义的标识符常量

  - 枚举常量

- **字符串：**

  “sasdad”这种双引号引起来的一串字符称为字符串字面量，或者简称字符串
  字符串的结束标志是一个|0的转义字符,在计算字符串长度的时候|0是结束标志,不算字符串内容

  c语言并没有字符串类型，只有字符类型 

  ```c
  char arr[]="aedsfsdd"
       
      strlen("abc")//求长度的函数,返回数字
      strlen(arr)//8
  ```

  ![74333127453](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743331274537.png)

  arr2后面没有|0，导致在内存中继续往下打印，生成了无关的东西。\0是结束标志

- **转义字符:**

  即转变字符的意思,如:

  - \n换行

  - \0结束

  - \?在书写连续多个问好时使用,防止他们被解析成三字母词

  - \'用来表示字符常量'和\"一样

  - \\用来表示一个反斜杠,防止他被解释为一个转义序列符

  - \a警告字符,蜂鸣

  - \b退格符

    \f进纸符

  - \r回车符

  - \t水平制表符

  - \v垂直制表符

  - \ddd表示1-3八进制的数字,如:\130x(八进制是0-7)

  - \xdd dd表示2个十六进制数字,如:\x30 0

  - 把他们转义成十进制的转义字符

  - 现在已经没有三字母词了,不转义的三字母词会出现方块 

    ​

  > 打印的特殊情况
  >
  > %d-打印整型
  >
  > %c-打印字符
  >
  > %s-打印字符串
  >
  > %f-打印float类型的数据
  >
  > %lf-打印doublr类型的数据
  >
  > 只有字符串可以省略一些打印方式,其他要严格使用

  ​

  - 编码:
    - ascll编码:把字符转化成十进制数字再转换成二进制数字

- **注释:**不需要的代码删掉或者注释掉,代码太难可以写注释来解释

  - // 可以嵌套
  - /**/  不能嵌套注释

- **学习方法：**

  ![74342091325](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743420913258.png)

- **选择语句：**

  ```c
  if（条件作判断）
  {
  //相关语句
  }
  else
  {
  //相关语句
  }

  return xxx
  ```


- **循环语句:**

  - while循环

    ```c
    while(line<20000)
    {
        printf("写代码")
        line++
    }

    return 0
    ```

    ​

  - for循环

  - do....while循环

- **函数：**简化代码,复用代码

  ```c
  Add（int x,int y）{
      int z = 0;
      z=x+y;
      return z
  }

  int main()
  {
      int n1=0;
      int n2=0;
      int sum =Add(n1,n2)
          printf("%d",sum)
      return 0
  }
  ```

- **数组:**一组相同类型的元素的集合

  ```c
  int arr[10]=[1,1,1,1,1,1,1,1,1,1,1,1] 
  //数组的大小被限制了
  ```

- **操作符:**

  - 算数操作符:+ - * / %

    > %是取余,/是除
    >
    > 除号两端是整数,执行整数除法,如果两边又一个浮点数的话,就执行浮点数除法
    >
    > a+=3\a-=3

  - 移位操作符>>  <<

  - 位操作符:& ^ |

  - 赋值操作符:= += -= *= /= ^= != >>= <<=

  - 单目操作符:

    > ![74346404620](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743464046201.png)
    >
    > sizeof计算数组的话,就是计算整个数组的大小
    >
    > ++/--先把值赋给b,再去加加减减(后置)
    >
    > 前置++/--就是先加加/减减再去使用
    >
    > int a =(int)3.14强制类型转换

  - 关系操作符:> >= < <= != ==

  - 逻辑操作符:&& ||

  - 条件操作符:三元表达式x?c:d

  - 逗号表达式:a,x,x,...ssss

    是一串逗号隔开的一串表达式,从左向右依次计算 ,整个表达式的结果是最后一个表达式的结果 

- **常见关键字:**

  是c语言内置的，关键字不是自己创建的

  ![74349332492](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743493324923.png)


  变量的名字不能是关键字

![74349402689](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743494026895.png)

- typedef和static:

  > typedef就是类型重命名,即类型定义

  ```c
  typedef unsigned int uint//uint是别名

  int main{
      unsigned int num =0
          uint num2=1
          
          return 0
  }
  ```

  > static是用来修饰变量和函数
  >
  > 1. 修饰局部变量，成为静态局部函数
  >
  > 2. 修饰全局变量，成为静态全局函数
  >
  > 3. 修饰函数，成为静态函数
  >
  > 4. ![74349778744](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743497787442.png)
  >
  > 5. 这个函数在调用后不会初始化，本质上是static修饰局部变量的说话，局部变量出了作用域，不销毁的。修饰局部变量的时候，改变了变量的存储位置
  >
  > 6. ​ void text(){}表示这个函数不需要返回
  >
  > 7. 全局变量：使用extent声明一个变量(函数)再使用，被static修饰的全局变量的外部连接属性就变成了内部属性，其他源文件（.c）不能使用全局变量
  >
  >    函数也是有外部连接属性的
  >
  >    ![74356951094](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743569510942.png)
  >
  > 8.   regesiter 寄存器关键字:建议变量存放在寄存器中,这样访问会更快,编译器会决定是否将数据防在寄存器中,防止挤爆寄存器(寄存器会自动把变量放在寄存器里面)
  >
  >    电脑储存设备:寄存器,高速缓存,内存,硬盘(最大,速度较慢,造价低)

- **#define定义宏:**

  ```c
  #define NLM 1000
  //定义了一个宏,可以是变量,也能是函数,参数是无类型的,
  #define ADD(x,y)((x)+(y))//宏名,宏的参,宏体 
  int main()
  {
      printf("%d",NLM)
      int n = NLM
  }
  ```

  ​

- **指针:**

   内存会划分一个个的内存单元(一个内存单元的大小:1byte),每个内存单元丢有一个编号

  ![74389688044](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743896880442.png)

  ![74389868165](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743898681654.png)

  > 命名变量的时候,其实就是申请内存空间,使用%P打印地址,是用&变量拿到这个变量的地址
  >
  > int *p=&a,这个的\*说明p是指针变量,int说明p指向的对象是int类型
  >
  > \*p   解引用操作符,通过p存放的地址,找到p所指向的对象,\*p就是p指向的对象

  内存单元:编号->地址->地址也就是指针

  > ![74389895191](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743898951917.png)

  ​

- **结构体:** 

- Stu就是结构体,结构体让c有描述复杂类型的能力

   ![74395267847](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743952678470.png)

  ![74395288515](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1743952885158.png)


结构体指针变量->成员名,也可以直接拿到相关的值去使用



- **分支和循环语句:**

  > c语言语句可以分为以下五类:
  >
  > 1. 表达式语句
  > 2. 函数调用语句
  > 3. 控制语句
  > 4. 复合语句
  > 5. 空语句

  控制语句用于控制程序的执行流程,以实现程序的各种结构方式,它们由特定的语句定义符组成,c语言有九种控制语句

  三类:

  > 条件判断语句:if        switch
  >
  > 循环执行语句:do while语句     while语句      for语句
  >
  > 转向语句:break语句   goto语句   continue语句   return语句

  - c语言是结构化的程序设计语言  

- 条件语句：

  ```c
  int main(){
      int age =10;
          if(age<18)
           printf(xxx)
           if(age<10)
               printf(xxx)
               //else会跟最近的哪个if，而不是最外面的if 
          else{
           printf(xxx)
          }
      return 0
  }
  ```


  //书：高质量的c++/c编程，养成良好的代码风格
  ```

- while语句

  ```c
  int i =1;
  //也是一种多分支语句
  while(i<=100)
  {
      case 1：
          printf(xxxx)
          i+=2
          //多个case同时实现
      case 3:
      case 4:
      case 2:
          printf(xxxx)
          break
              //退出循环
      default:
          //匹配不到任何值的时候执行
  }
  return 0
  ```

  > case 整型常量表达式:
  >
  > 语句

- 循环语句

  while、for、do while


```c
#include <stdio.>
int main()
{
    int i=1;
    while(i<=10)
    {
        printf(xxx);
        i++;
        if(5==i)
        {
            break;
            //break是永久终止循环
            //continue是跳过本次循环后面的代码,直接删去判断部分,进行下一次 
        }
    }
    return 0;
}

```

```c
//输入函数
#include <stdio.>
int main()
{
    int ch=0;
    while((ch=getchar()!=EOF))
        //拿到输入的数据
        //scanf也是一种输入方法
    {
        putchar(ch);
        //打印输入的
        //有一个输入缓存区,putchar去拿这个输入缓存区的值
        //只要输入缓存区有值他就会去拿
        //用while((chgetchar())!='\n'){}来清理缓存区的\n
        //返回值类型是int
        scanf();
        //不能读取空格,有空格就结束
        //也是拿到相关的值
    }
}



int main(){
    int ch = 0;
    while((ch=getchar())!=EOF)
        putcher(ch);
    return 0;
    //这里的代码可以适当清理缓冲区
}



int main(){
    char ch ='\0';
    while((ch=getchar())!=EOF)
    {
        if(ch<'0'||ch>'9')
            continue;
        putchar(ch);
    }
    return 0
}
//打印字符并且跳过相关的字符
```





 

####  for循环：

语法：for(表达式1；表达式2；表达式3)

​                      循环语句

循环语句是两条以上的行数，就要加一个{}

**表达式1**

初始化循环变量

**表达式2**

条件判断语句部分，用于判断循环时候终止

**表达式3**

调整部分，用于循环条件的调整

![74506409080](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1745064090802.png)

continue是跳过这次循环，本次循环后面的代码是不会执行的，会直接把for循环跳到判断部分去了，而while循环会直接把循环跳到判断部分去了

break就是单纯的终结了





**建议：**

1. 不可以在for循环修改变量，防止for循坏失去控制

   ​

2. 建议for语句的循环控制变量的取值采用前闭合开区间的写法

![74506472041](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1745064720415.png)



**变种：**
/
1. for循环的判断部分省略意味着这部分判断会恒成立

2. 嵌套for循环

   ![74506505803](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1745065058035.png)

3. 可以直接在for循环里面创建变量

   for（int i=0；i<10;i){

   xxxxxx

   }

   注意有点编译器不支持c99语法







####  do...while循环

```c
do
{
    //循环语句
    //两条以上加大括号
}
while(表达式)
```

![74506550477](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1745065504772.png)

循环流程图,如上

特点:循坏至少执行一次,使用场景有限,用到不多





####  练习

```c
//打印1！+2！+...10!
int line(int n){
    if(n==1)
        return 1;
    return n*line(n-1)
}

int main(){
    int ret = 0;
    for(int i =0 ;i<10;i++){
        ret=ret+line(i);
    }
    printf("%d\n",ret);
        return 0;
}


//在一个有序数组中查找n
int binarySearch(int arr[], int target, int left, int right) {
    if (left > right) {
        return -1; // 未找到目标
    }
    int mid = left + (right - left) / 2; // 防止整数溢出
    
    if (arr[mid] == target) {
        return mid; // 找到目标，返回下标
    } else if (arr[mid] < target) {
        return binarySearch(arr, target, mid + 1, right); // 查找右半部分
    } else {
        return binarySearch(arr, target, left, mid - 1); // 查找左半部分
    }
}
//避免复制数组

//使用左右的left和right的值来实现二分查找
int main(){
   int arr[]={1,2,3,4,5,6,7,8,9,10};
    int k = 7;
    
    int i = 0;
    int sz = sizeof(arr)/sizeof(arr[0]);
    
    int left = 0;
    int right = sz - 1;
    int mid = (left+right)/2;
    
    whihe(){
        if(arr[mid]<k){
            left = mid +1
        }else if(arr[mid]>k){
            rigtht = mid -1
        }else{
            print();
            break
        }
    } 
    
    //不断重复这个过程,用while去查找
    
    return 0;
}

//二分查找
int flod(int arr[],int target,int felt,int right){
    if(relt>right){
        return -1;
    };
    
    
    int mid = left+(right-left)/2;
    
    if(arr[mid]==target){
        return mid
    }else if(arr[mid]<target){
        return flod(arr,target,mid+1,right)
    }else(arr[mid]>target){
        return flod(arr,target,felt,mid-1)
    }
}




//编写代码，演示多个字符从两端移动，向中间汇集 
int main(){
    char arr1[]="xxxxxxxxxxx";
    char arr2[]="nihaowohaoa";
    
    int left=0;
    int right=strlen(arr2)-1;
    
    while(left<=right){
        arr2[left]=arr1[left];
        arr2[right]=arr1[right];
        printf("%s\n",arr2);
        Sleep(1000);
        //休眠1秒
        left++;
        right--;
    }
    return 0
}



//电脑产生一个随机数
//猜数字
//猜大了
//猜小了
void menu(){
    printf("****************\n");
    printf("***1.play*******\n");
    printf("***0.exit*******\n");
}

void game(){
    //电脑生成随机数的方法
    int guess = 0;
    //时间戳: 表示时间的刻度
    let ret= rand()%+1;
    //0~32767是生成的范围
    //猜大了
    //猜小了
    //直到猜对了
    while(1){
         printf("请猜数字:>");
    scanf("%d",&guess);
    if(guess<ret){
        printf("猜小了\n");
    }
    else if(guess>ret){
        printf("才打了\n");
    }
    else{
        printf("猜对了\n");
        break
    }
    }
}


int main(){
    int input = 0;
        //设置一个生成起点
    srand((unsigned int)time(NULL));
    //NULL是一个空指针
    do{
        menu();
        printf("请选择：》");
        scanf("%d",&input);
        switch(input){
            case 1:
                game();
                printf("猜数字\n");
                break;
            case 0:
                printf("退出游戏\n");
                break;
            default:
                printf("选择错误,重新选择!\n")
                break;
        }
    }while(input);
    //写代码注意写一步测一步,防止bug干扰你的书写
    return 0;
}
```

> debug是调试版本,程序调试的版本
>
> release是发布版本,给用户使用的版本



####  函数

是一个大型程序中的某部分代码，由一个或者多个语句组成，负责完成某项特定任务，而且相较于其他代码，具备相对的独立性

一般会输入参数并且有返回值，提供封装和细节的隐藏，就是软件库

- 库函数

  一些基础库，储存了你需要使用的函数

  ![74712642670](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1747126426707.png)

  使用库函数必须包含#incude对应的头文件

```c

int main(){
    char arr1[20]={0};
    char arr2[]="hello";
    
    strcpy(arr1,arr2);
    
    print("%s\n",arr1);
    //strcpy可以将arr2复制给arr1
    
    
    
    char arr[20]="hello world";
    memset(arr,"x",5);
    printf("%s\n",arr);
    //将arr的5个字节替换为x
    //阅读文档学习函数方法
    return 0;
}
```

- 自定义函数

  自定义函数，由我们来设计

  ![74712676378](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1747126763782.png)

  ```c
  //形式参数
  //实参传递给形参的时候，是实参的一次拷贝
  //改变形参不影响实参
  //如果要改变实参，传地址进去
  get_max(int x,int y){
      return (x>y?x:y)
  }

  int main(){
      int a = 0;
      int b = 0;
      //实际参数
      scanf("%d,%d",&a,&b);
      
      int m =get_max(a,b);
      //拿到最大值
      
  }
  ```

  ​

####  嵌套调用

![74910590087](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749105900870.png)

政策的嵌套调用

 ![74910554559](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749105545596.png)

错误的调用，c语言不支持这种调用方法

 #### 链式调用

![74910587667](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749105876672.png)

最典型的链式调用

> printf("%d",ptintf("%d",ptintf("%d",43)))
>
> 结果4321,printf返回了字符打印的个数

![74910633741](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749106337412.png)

但是尽量不要这样做,写出返回值是好习惯

 写了返回类型,又不写return,会返回最后一条指令执行的结果,不适合所有编译器,也是一种不规范的行为

相似的还有,没有形式参数就不要传参数比较好

![74910664654](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749106646549.png)

声明一个void,不要传递参数

![74910674995](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1749106749958.png)

一般来说,main函数不需要参数,但是它还是可以传参数的

####  链式访问

把一个函数的返回值做另外一个函数的参数



####  函数的声明和定义

- 函数声明：告诉编译器一个函数叫什么，参数是什么，返回类型是什么，但具体是不是存在，函数决定不了
- 函数声明一般出现在函数使用之前，要满足先声明后使用
- 函数声明一般在头文件

函数的定义是指函数的具体实现，交代函数的功能实现

```c

int Add(int x,int y){
    return x+y
}

int main(){
    int a = 0;
    int b = 0;
    int sun = Add(a,b)
}
//要放在前面，而不是后面，除非先声明

int Add(int,int);
//返回值返回类型和参数

int main(){
    int a = 0;
    int b = 0;
    int sun = Add(a,b)
}

int Add(int x,int y){
    return x+y
}

//或者再去使用include "add.h",引入模块

```

![75168399674](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1751683996742.png)

可以把模块打包成静态库，再进行使用，静态库源码不会泄露

![75168554462](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1751685544620.png)



####  函数递归  

程序调用自身的编程技巧称为递归

递归做为一种算法再程序设计语言中广泛应用，一个过程或者函数在其定义或说明中的直接或者间接调用

常常把大型复杂的问题转化为小问题求解的策略

![75168792721](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1751687927214.png)



> %d是打印符号的整数
>
> %u是打印无符号的整数

内存空间

> 栈区：局部变量，函数的形参，函数的调用都会在栈区上申请空间
>
> 堆区:   
>
> 静态堆:



输入某个数字，依次打印

![75169081281](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1751690812818.png)



![75169082018](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1751690820188.png)



```c
int my_strlen(char* str){
    int count = 0;//临时变量
    while(*str:'\0'){
        count++;
        str++;//下一个字符
    }
    return count
}

//不创建临时变量
int my_strlen(char* str){
    if(*str!='\0')
        return 1+my_strlen(str+1);
    //不能用str++，会导致str的地址又被传进去
    else
        return 0
}
```

 ```c
int fac(int n){
    if(n<=1)
        return 1;
    else
        return n*fac(n-1)
}
//递归写法
int fac(int n){
    int i = 0;
    int ret = 1;
    for (i=1;i<=n;i++){
        ret*=i
    }
    return ret
}
//迭代的写法
 ```

递归并不适合所有情况,比如求斐波那契数列的时候

![75256436088](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752564360886.png)

 这里存在严重的重复,导致开销大,时间久

![75256520918](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752565209186.png)

将前面计算过的值缓存起来,但是这个大数的时候可能会溢出



> 栈溢出的现象
>
> ![75256537360](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752565373601.png)

![75256547979](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752565479799.png)

汉诺塔和青蛙跳台

`![75256564911](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752565649111.png)







####  数组

1. 一维数组

   数组是一组相同元素的集合

   > int arr[10]
   >
   > char ch[5]
   >
   > double data[20]

    ![75263904034](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752639040340.png)

   ![75263910263](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752639102633.png)

   变长数组不支持初始化

   ![75263936510](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752639365103.png)

   不同的内容初始化也不一样

   也可以不写大小

   char ch4[]={"1","2","3"}

   数组创建的时候如果不想指定大小就得初始化,数组的元素个数根据初始化的内容来确定

2. 储存位置

   > int arr[]={1,2,3,4,5,6,7,8,9}

   通过arr[4]拿到相关的值

   拿到数组大小

   int s2 = sizeof(arr)/sizeof(arr[0])

   **数组储存在一片连续空间内**

   ![75264010922](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752640109225.png)

3. 二维数组和越界

   完全初始化

   int arr2[3]\[4]={1,2,3,4,5,6,7,8,9,10}

   不完全初始化

   int arr1[3]\[4]={{1,2},{3,4},{5,6}}

   生成![75264397237](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752643972377.png)

   ![75264404259](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752644042599.png)

   二维数组的初始化

   二维数组可以看为一维数组的数组

   二维数组在储存里面也是连续的，一行接上一行，行是可以省略的，但是列不能省略

   也是会存在越界的行为，要做好越界检查

4. ![75264543108](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752645431085.png)

   ![75264580427](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752645804270.png)

   **不要在函数内部求数组大小，数组作为参数的时候，不是把整个数组传进去的**

   数组名的本质是数组首元素的地址

5. 数组名

   ![75281675414](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1752816754144.png)

   除去这两种情况,其他所有数组名都是数组首元素的地址

   一维数组数组名

   ```c
   int main(){
       int arr[10]={0}
       print("%p\n",arr);//arr就是首元素的地址
       printf("%p\n", arr+1);
       printf("---\n");
       printf("%p\n", &arr[0]); // 首元素的地址
       printf("%p\n", &arr[0]+1);
       printf("---\n");
       printf("%p\n", &arr);    // 数组的地址
       printf("%p\n", &arr+1);
   }
   ```

   二维数组数组名

   也表示首元素的地址,就是一行的地址

   总数组的大小除以一行的大小,就是行数了,列数也是同理的

   就是行的大小除以一个元素的大小,就是列数

   ​

6. 数据实例

   **三子棋**

   test.c//负责测试游戏的逻辑

   game.c//游戏逻辑的实现

   game.h//游戏代的声明,函数声明,符号定义

   ​

   在test.c

   ```c
   #include <studio.h>

   printf("*********************\n");
   printf("*** 1.play 0.exit****\n");
   printf("*********************\n");
   void game(){
       char board[ROW][COL]={0};
       InitBoard(board,ROW,COL);
       //初始化
       DispalyBoard(board.ROW,COL);
       //下棋
       while(1){
           playerMove();
           ret = IsWin（Board,ROW,COL);
           if(ret != "c"){
               break
           }
           DispalyBoard(board,ROW,COL);
           ComputerMove(board,ROW,COL);
           ret = ISWIn(board,ROW,COL);
           if(ret != "C"){
               break
           }
           DispalyBoard(board,ROW,COL);
          
       }
       if(ret="*"){
           //玩家胜利
       }else if(ret = "#"){
           //电脑胜利
       }else{
           //平局
       }
   }

   int main(){
       int input = 0;
       do{
           menu();
           printf("请选择:>");
           scanf("%d",&input);
           switch(input){
               case 1:
                   game();
                   break;
               case 0:
                   printf("退出游戏\n");
               default:
                   printf("选择错误!\n");
                   break;
           }
       }while(input);
       return 0
   }
   ```

   game.c 

```c
   void InitBoard(char board[3][3],int row,int col){
       int i = 0;
       int j = 0;
       for(i=0;i<row;j++){
           for(j=0;j<col;j++){
              board[i][j]='';
           }
       }
   }
   void DispalyBoard(char board[ROW][COL],int roww,int col){
       int i = 0;
       for(i=0;i<row;i++){
           //printf("%c | %C | %C \n",board[i][0],board[i][1],board[i][2]);
           int j = 0;
           for(j=0;j<col;j++){
               printf("%c",board[i][j]);
               if(j<col-1)
                   printf("|")
           }
           if(i<row-1){
               int j = 0;
               for(j=0;j<col;j++){
                   printf("---");
                   if(j<col-1)
                        printf("|")
               }
               printf("\n")
           }
       }
   }

   void layerMove(char board[ROW][COL],int row,int col){
       int x = 0;
       unt y = 0;
       
       while(1){
               printf("玩家下棋:>\n");
       printf("请输入坐标:>");
       scanf('%d %d',&x,&y);
       if(x>=1&&x<=row&&y>=1y<=col){
           if(board[x-1][y-1]==''){
               board[x-1][y-1]="*";
                   break;
           }else{
               printf("坐标被占用")
           }
       }else{
           printf("坐标非法\n")
       }
       } 
   }
   void ComputerMove(char board[ROW][COL], int row, int col) 
   {
       printf("电脑下棋:>\n");
       int x = 0;
       int y = 0;

       while (1) 
       {
           x = rand() % row; // 生成0到row-1的随机数
           y = rand() % col; // 生成0到col-1的随机数
           
           if (board[x][y] == ' ') // 检查该位置是否为空
           {
               board[x][y] = '#';  // 电脑下棋，用'#'标记
               break;              // 成功下棋后退出循环
           }
       }
   }
```
   在game,h

   ```c
   #pragna once

   #include <stdio.h>
   #include <stdlib.h>
   #include <time.h>

   #define ROW 3
   #define COL 3

   //初始化
   void InitBoard(char board[ROW][COL],int row,int col);

   // 打印棋盘
   void DisplayBoard(char board[ROW][COL], int row, int col);

   // 玩家下棋
   void PlayerMove(char board[ROW][COL], int row, int col);

   // 电脑下棋
   void ComputerMove(char board[ROW][COL], int row, int col);

//玩家赢*
//电脑赢#
//平局Q
//没有就继续游戏C
char IsWin(char board[ROW][COL],int row,int col)
   ```



二维数组，n*m，并且   ​数组初始化为空格





####  变量和作用域

![75914291848](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759142918481.png)

![75914323667](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759143236674.png)

![75914336018](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759143360186.png)



   







####  操作符

- 算术操作符

  +-*%/

  整数是直接除

  浮点操作符要确保两边有一个浮点数

  取模必须是整数

- ![75914388765](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759143887658.png)

  左移操作符，移动的是二进制位，整数的二进制表示有两种，源码，反码，补码

  正整数的源码，反码和补码相同

  负整数的原码，反码，补码需要计算

  ![75914452370](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759144523700.png)

  整数在内存里面储存的是补码

  移位操作符操作的是补码

  ![75914480557](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759144805579.png)

  不会改变原来的值，只是产生了一个新值，左移会有乘2的效果

  ![75914549856](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759145498562.png)

  ​

  a >> 1右移一位，不能是负数，c是未定义这种操作的

  ​

- 位操作符

  & - 按2进制位与

  | - 按2进制位或

  ^ - 按2进制位异或

  与操作:有零就是0,两个为一才是1

  或:只要有1就是1,两个为零才是0

  异或:相同为0,相异为1

  使用异或交换数据

  ![75931061564](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759310615641.png)

  a^0=a

  a^a=0

  ​

- 赋值操作符

  ![75931124102](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759311241028.png)

  ![75931135864](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759311358645.png)

  用法是

  ![75931137807](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759311378079.png)

- 单目操作符

  ![75931139605](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759311396055.png)

  单目操作符

  只有一个操作数

  +

  a+b双目操作符,＋操作符有两个操作数

  取地址操作符

  ![75931197107](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759311971070.png)

  sizeof可以计算遍历的变量所占空间大小

  ![75931209519](C:\Users\zxh\Desktop\前端\c语言\c语言.assets\1759312095197.png)

  计算数组的大小,字节
