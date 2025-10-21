#  M Gws'q'aCanvas

canvas类似于网页上的画板

> 需要准备的东西
>
> id:标识元素的唯一性
>
> width:画布的宽度
>
> height:画布的高度

```html
<canvas id="cl" width="600" height="400" style="width:600px;height:400px">
</canvas>
```

这两个宽高的设置是有区别的

**设置画布只之后,就要进行相关的绘制V**

```javascript
// 1. 获取canvas元素
var cl = document.getElementById("cl");
// 2. 获取上下文对象
var ctx = cl.getContext("2d");  // 修正大小写
// 3. 绘制矩形
ctx.fillRect(0, 200, 300, 300);  // 移除非数组格式
```





####  canvas上下文对象与浏览器对象

普通元素是没有getContext的,可以用来做判断浏览器配置,没有这个的,基本都是十年前的浏览器了->浏览器兼容问题

![75380549232](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753805492324.png)

画笔定义,可以看文档来认识这些东西





####  canvas路径和填充

 矩形绘制

1. 填充模式

   > ctx.fillRect(0, 200, 300, 300); 

2. 路径模式

   > ctx.strokeRect(100,100.200,200)
   >
   > 路径绘制矩形strokeRect（x1,y1,矩形宽度，矩形高度）

3. 清除模式

   > ctx.fillReact(100,100,200,300)
   >
   > ctx.clearRect(0,0,c1.clientWidth,c1.clientHeight)

拆开绘制

> ctx.beginPath()
>
> 画笔抬起
>
> ctx.rect(100,100,100,100)
>
> ctx.strock()
>
> ctx.closePath()
>
> 画笔落下
>
> 拆开绘制路径

形成分段的绘画,使用beginPath和closePath完成画笔的分段





####  绘制笑脸和圆弧

![75384855591](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753848555917.png)

绘制直线使用lineTO（）

 绘制圆

> ctx.arc(圆心坐标,半径)
>
> ctx.arc(300,200,10)
>
> ctx.fill()

 绘制圆弧

> ctx.arc(圆心x,圆心y,半径,开始的角度,结束的角度,逆时针还是顺时针默认是顺时针false设置顺时针为true)
>
> ctx.arc(300,200,50,0,Math.PI/2)

![75385134021](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753851340215.png)

一种绘制方法

另外一种绘制方法是移动画笔

![75385164352](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753851643521.png)





####  绘制折线

使用lineTo（）绘制直线

 ![75396354296](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753963542965.png)





####  其他绘制圆弧的方法

使用两条线和半径绘制

![75396406568](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753964065683.png)

```javascript
// 1. 获取画布元素
var cl = document.getElementById("cl");

// 判断浏览器是否支持canvas
if (!cl.getContext) {
    console.log("当前浏览器不支持canvas，请下载最新的浏览器");
}

// 2. 获取2D上下文
var ctx = cl.getContext("2d");

// 开始绘制路径
ctx.beginPath();
ctx.moveTo(300, 200);             // 起点
ctx.arcTo(300, 250, 250, 250, 25); // 绘制圆弧
ctx.stroke();                      // 描边
ctx.closePath();                   // 关闭路径
```

- 使用 `arcTo()` 方法绘制圆弧：
  - 控制点1: (300, 250)
  - 控制点2: (250, 250)
  - 半径: 25px





####  多次贝塞尔曲线

![75396492519](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753964925193.png)

是前端很重要的一种曲线

二次贝塞尔曲线需要三种点:起点,控制点,终点

![75396616721](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1753966167216.png)



**例子:绘制气泡**

```javascript
// 1. 获取画布元素
var cl = document.getElementById("cl");

// 判断浏览器是否支持canvas
if (!cl.getContext) {
    console.log("当前浏览器不支持canvas，请下载最新的浏览器");
}

// 2. 获取2D上下文
var ctx = cl.getContext("2d");

// 开始绘制路径
ctx.beginPath();
ctx.moveTo(200,300)
ctx.quadraticCurveTo(150,300,150,200)
ctx.quadraticCurveTo(150,100,300,100)
ctx.quadraticCurveTo(450,100,450,200)
ctx.quadraticCurveTo(450,300,250,300)
ctx.quadraticCurveTo(250,350,250,350)
ctx.quadraticCurveTo(150,350,150,300)
ctx.closePath();                   // 关闭路径
```







####  path2D封装

```javascript
// 1. 获取画布元素
var cl = document.getElementById("cl");

// 判断浏览器是否支持canvas
if (!cl.getContext) {
    console.log("当前浏览器不支持canvas，请下载最新的浏览器");
}

// 2. 获取2D上下文
var ctx = cl.getContext("2d");

var heartPath = new Path2D();

// 开始绘制路径
heartPath.moveTo(200,300)
heartPath.quadraticCurveTo(150,300,150,200)
heartPath.quadraticCurveTo(150,100,300,100)
heartPath.quadraticCurveTo(450,100,450,200)
heartPath.quadraticCurveTo(450,300,250,300)
heartPath.quadraticCurveTo(250,350,250,350)
heartPath.quadraticCurveTo(150,350,150,300)
ctx.closePath(heartPath);                   // 关闭路径

ctx.full(heartPath)
//直接填充颜色

//完成了heartPath的封装
```

Path2D也有一些特别的配置,看文档





####  颜色样式的控制

Colors色彩

> fullStyle = color
>
> 设置填充的样式
>
> strockStyle = color
>
> 设置图像轮廓的样式颜色

color可以是css颜色的字符串,渐变对象或者图案对象

> 一旦设置了strokeStyle或者fillStyle的值,那么这个新值就会成为新绘制的图形的默认值,如果你要给每个图形设置新的颜色值,就需要重新设置fullStyle或者StrokeStyle



####  线性渐变和径向渐变

1. 线性渐变

   ![75585737606](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755857376061.png)

   可以任意去插颜色进去

2.  线性渐变

   对于圆形进行作用

   ![75585797428](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755857974283.png)

3. 径向渐变模拟3D球

   ![75585864030](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755858640302.png)

4. 圆锥渐变

   一个圆

   ![75585894744](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755858947446.png)

   ![75585896357](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755858963578.png)

   注意是createConicGradient

   这个中间也是能插很多颜色的

   ​



####  加载印章

绘制图案

![75585926402](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755859264023.png)

createPattern创建印章



####  线条样式

1. lineWidth

   设置线条的粗细

   ![75586187226](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755861872267.png)

   > ctx.lineWidth=number
   >
   > //设置端口的样式
   >
   > ctx.lineCap='butt'|'round'|'square'
   >
   > //平齐,半圆,正方形

2. lineJion

   设置线条链接的时候的样式

   ![75586221181](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755862211815.png)

   > ctx.lineJion = 'mitter'|'round'|'bevel'
   >
   > mitter相连的角,round圆润的角,bevel磨平

3. miterLimit()

   设置拐角处的样式

   ![75586224103](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755862241036.png)

   > ctx.miterLimit=number
   >
   > 变尖尖或者磨平

4. setLineDash

   设置虚线

   > ctx.setLineDash([10,2])
   >
   > 这样会设置折线长度和空白区域
   >
   > 传入一个数组做为参数

    ![75586351594](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1755863515948.png)

   会动的线条,注意清除线条

   ​




####  阴影设置

```javascript
ctx.shadowOffsetX = 10;
ctx.shadowOffsetY = 10;
ctx.shadowBlur = 5;
ctx.shadowColor = `rgba(255)`
```

先设置,然后画出来的就会有阴影效果

 



####  设置图像或者视频

```javascript
var ctx = cl.getContext('2d')

let img = new Image();
img.src = "./imgs/girl.webp"
img.onload = function(){
    // ctx.drawImage(img)
    //直接绘制图片
    //ctx.drewImage(img,0,0,600,400)
    //这个就指定了图片的大小600*400
    // ctx.drawImage(img,640,0,1280,720,0,0,580,400)
    //把前面四个数字的裁剪成的图片放到后面四个数字的位置
    
}
```



####  绘制视频和水印

![75818323693](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758183236930.png)

不断调用请求函数,,然后在画板上绘制这个东西

![75818354694](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758183546942.png)

在把水印图片也加载上去

就完成了水印的绘制



####  绘制文字在canvas

文字有大小,有字体

```javascript
ctx.font = "100px microsoft YaHei";
ctx.fillText("老陈",300,200)
//文本,文本起点的x,文本起点的y,绘制文字的最大宽度.这些都是那些参数
ctx.strokeText("老陈",300,200)
//这个是描边绘制
//也能去设置颜色

//文本对齐选项
ctx.textAlign = "center"
//还有start,end和left和right,文本对齐
ctx.textHasetLine = "middle"
//文本基线对齐,textBaseLine,top,bottom,alphabatic


ctx.direction = 'rtl'
//文本方向

let text = ctx.measureText('hello')
//预测量文本的宽度,得到一个宽度数据
```







####  位移缩放旋转

 ![75860846267](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758608462679.png)

拉伸  

旋转是：ctx.route(Math.PI/6 ),一个基于原点的旋转

使用translate移动旋转点

![75860913112](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758609131128.png)



####  变形，transforms

![75860927066](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758609270661.png)

旋转·缩放·位移

矩阵来控制变化

![75860985510](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758609855102.png)

水平位移

![75861001256](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758610012566.png)

旋转45°







####  合成

![75861022162](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758610221628.png)

![75861028001](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758610280016.png)



很多很多滤镜

 ![75861121175](C:\Users\zxh\Desktop\前端\canvas\Canvas.assets\1758611211758.png)

第二张图写的时候才进行修改模式







