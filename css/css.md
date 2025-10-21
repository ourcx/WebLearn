# Css动画

####  浏览器

![74446521595](C:\Users\zxh\Desktop\前端\css\css.assets\1744465215954.png)

   ####  圆角

![74446593537](C:\Users\zxh\Desktop\前端\css\css.assets\1744465935378.png)

```css
#box1{
    border-radious:10px 20px 30px 40px / 10px 20px 30px 40px;
    /*每个角的水平到每个角的垂直*/
    圆角
    1.从指定的角的顶端，向内部引出垂直半径和水平半径
    2.将水平半径和垂直半径相较于元素内部的一点（圆心点）
    3.以该点为圆心，指定的垂直半径和水平半径画圆
    4.与边框相交的部分就是圆角部分
    border-redius:20px 60px;
    //20px代表左上角和右下角的水平和垂直，60px代表左下角和右上角的水平和垂直 
    border-redius：20px
    //圆角
}
```





####  阴影

1. text-shadow:1px 2px 3px red，6px 3px 2px orange

   1px是水平位置，2px是垂直位置。3px是扩散大小(模糊度)

2. box-shadow:-5px 5px 2px orange

   和文字阴影类似,但有一些增加

   box-sahdow:5px 6px 2px 3px red

   3px是外延值,可以让阴影往外延申,不加颜色是默认阴影,加上inset是内置阴影,多个阴影用逗号分割

   > 阴影的格式: 
   >
   > 水平偏移位置-垂直偏移位置-模糊度-外颜值-颜色-(内置阴影)



####  渐变

![74469759222](C:\Users\zxh\Desktop\前端\css\css.assets\1744697592223.png)

渐变属于背景图片

1. 线性渐变:direction取值有to top\或者使用度数

![74469845902](C:\Users\zxh\Desktop\前端\css\css.assets\1744698459026.png)

2. 环形渐变:

   中心:语法是at x y 用左上角原点为参考点,可以是像素,百分比

   形状:elipse椭圆 默认值,cricle圆

   大小:最近边:closest-side,最远边farthest-side,最近角 closest-corner,最远角farthest-corner







####  transform(2D)

语法：transform：函数名(x轴值,y轴值)

> translate()    scale()       rouate()      skew()

1. translate():

   根据左x轴和顶部y轴位置给定的参数,从当前元素位置移动,接受px

2. ![74478320258](C:\Users\zxh\Desktop\前端\css\css.assets\1744783202582.png)

   transform:rotate(20deg)

3. scale(0.5)缩小一倍,进行减小

   可以有两个参数,表示水平和垂直方向的缩放 

4. skew()根据水平轴和垂直轴反转,接受两个或者一个值,两个值时前面为水平,后面为垂直的角度,一个值只是水平轴的角度,此函数是指元素的倾斜角度



####  transform(3D)

要设置好透视和3D效果

 ![74478730655](C:\Users\zxh\Desktop\前端\css\css.assets\1744787306554.png)

z轴变化

transform: perspective:400px  translate3d(x,y,z)

perspective也可以写在transform里面

在2d那几个属性加一个3d就是3d效果了







####  过渡属性

使css属性值在一段时间内平滑过渡

比如鼠标悬停后,背景色在1s内,由白色平滑的过渡到红色

指定的四个要素:

1. 过渡属性:background\color等

   transition-property:  ;


2. 过渡的时间:

   transition-duration:s|ms

   默认是0,意味不会有效果,所以是必须要设置的

3. 过渡函数:

   transition-timing-function:   ;

   ease:默认,慢到快,然后再慢

   linear:匀速

   ease-in:规定以慢速开始,加速效果

   ease-out:规定以慢速结束,减速效果

   ease-in-out:规定以慢速开始和结束,先加速后减速 

4. 过渡延迟时间

   transition-deley:s|ms

触发过渡:点击\悬停\聚焦









####  animation动画

可以制作像flash类似的动画,通过关键帧去制作动画

1. @keyframes

   作用:用于声明动画,指定关键帧

   每个帧代表某个时间点

   定义每个帧上的动作


![74485732220](C:\Users\zxh\Desktop\前端\css\css.assets\1744857322206.png)

2. animation属性

   animation属性用于控制动画,调用由@keyfrom定义的动画,设置动画属性,如时间\次数等,animation属性是一个简写属性

   ![74485777893](C:\Users\zxh\Desktop\前端\css\css.assets\1744857778935.png)

3. 子属性

   ![74485784296](C:\Users\zxh\Desktop\前端\css\css.assets\1744857842967.png)

   ![74485785404](C:\Users\zxh\Desktop\前端\css\css.assets\1744857854045.png)

   ​


#  特效

####  页面背景动态

![74669153118](C:\Users\zxh\Desktop\前端\css\css.assets\1746691531184.png)

html如上

```css
*{
    margin:0;
    padding:0;
    box-sizing:border-box;
    font-family: '微软雅黑',serif;
}
body,html{
     height:100vh;
     width:100%;
}
.container{
    height:100vh;
    width:100%;
}

.context{
   position:absolute;
   width:100%;
   height:100%;
   display:flex;
   justify-content:center;
   align-items:center;
}

.context h1{
   font-size:3.5rem;
   color:#fff;
}
.area{
   background:#4354c8;
   height:100%;
   width:100%;
}

.circles{
   position:absolute;
   top:0;
   left:0;
   width:100%;
   height:100%;
   overflow:hidden;
   z-index:1;
}

.circles li{
   position:absolute;
   display:block;
   list-style:none;
   width:20px;
   height:20px;
   background:#ffffff33;
   bottom:-150px;
   animation: animate 25s linear infinite;
   /* 线性无限 */
}

.circles li:nth-child(1){
   left:25%;
   width:80px;
   height:80px;
   animation-delay:0s;
}

/* 在复制分别处理剩下十个，delay要设置多一点，还有设置大小和duration不同 */

@keyframes animate {
   0%{
      transform: translateY(0) rotate(0deg);
      opacity:1;
      border-radius:0;

   }
   100%{
      transform: translateY(-1000px) rotate(720deg);
      opacity:0;
      border-radius:50%;
   }
}


```

![74669231209](C:\Users\zxh\Desktop\前端\css\css.assets\1746692312091.png)



####  滑动效果

html放入三张图片

![74669253093](C:\Users\zxh\Desktop\前端\css\css.assets\1746692530938.png)

![74669254525](C:\Users\zxh\Desktop\前端\css\css.assets\1746692545252.png)

关键属性是background-attachment：fixed

页面滚动时背景图固定

![74669259531](C:\Users\zxh\Desktop\前端\css\css.assets\1746692595313.png)

可交互的视差网页效果，上下遮挡和一些隐藏





####  elementplus主题切换

使用了一个api：view Transitions API

![74669281244](C:\Users\zxh\Desktop\前端\css\css.assets\1746692812440.png)

![74669282929](C:\Users\zxh\Desktop\前端\css\css.assets\1746692829296.png)

代码

```css
:root{
    --background-color:#fff;
    background-color:var(--background-color)
}

:root .dark{
    --background-color:#000000
}

::view-transition-new(root),
::view-transition-old(root){
    animation:none
}
//关闭默认动画
```

```javascript
const btn = document.qurySelector("#btn")
btn.addEventListener("click",(e)=>{
    const transition=document.startViewTransition(()=>{
        document.documnertElement.classList.toggle("dark")
    })
    
    const x=e.clientX
    const y=e.clientY
    
    //半径
    const tragetRadius=Math.hypot(Math.max(window.innerWidth-x),Math.max(window.innerHeight-y))
    
    
    transition.ready.then(()=>{
        docunment.documentElement.animate({
            clipPath:["circle(0%at${x}px ${y}px)","circle(${tragetRadius}at${x}px ${y}px)"]
        },{
            duration:300,
            pseudoElement:"::view-transition-new(root)"
        })
    })
})
//切换颜色




```

####  视频背景

```html
<!doctype html>
<html>
<head>
</head>
 
<body style="height: 100%;width: 100%;position=relative">
	<video src="http://www.longsong.online/wp-content/uploads/2022/03/mulai.mp4" style="width: 100%;height: 100%;object-fit: cover;position: absolute;top: 0;left: 0;" autoplay="autoplay" loop="loop" muted="muted"></video>
</body>
 
 
</html>
```

关键是object-fit: cover，加上后面的静音和循环播放

####  最简单水平居中

```css
<div class="g-container">
    <div class="g-box"></div>
</div>

.g-container {
    width: 100vw;
    height: 100vh;
    
    display: flex;
    // display: grid;
    // display: inline-flex;
    // display: inline-grid;

}

.g-box {
    width: 40vmin;
    height: 40vmin;
    background: #000;
    margin: auto;
}
```

简单翻译一下，大意是在 **flex 格式化上下文**中，设置了 `margin: auto` 的元素，在通过 `justify-content` 和 `align-self` 进行对齐之前，任何正处于空闲的空间都会分配到该方向的自动 margin 中去,但是局限性太大，仅仅是能让文字水平垂直。如果给一个子元素加一个宽高或者任何其他属性就不行了



####  [grid 布局配合 clip-path 实现 GTA5 封面](https://chokcoco.github.io/CSS-Inspiration/#/./layout/grid-clip-path-gta5-poster?id=grid-%e5%b8%83%e5%b1%80%e9%85%8d%e5%90%88-clip-path-%e5%ae%9e%e7%8e%b0-gta5-%e5%b0%81%e9%9d%a2)

- grid 实现不规则的网格布局
- clip-path 控制每个格子的形状

```scss
<div class="parent">
  <div class="child" >
    <img src="https://i.pinimg.com/originals/0d/67/72/0d677237854ed19dcfe69f0f9a4065ee.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/736x/26/db/84/26db84b2bf348f79792f7c5f0f9bd5ef.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/736x/45/0d/1c/450d1c87ce61bc0c2429701ed3ea631a.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/564x/94/76/dd/9476dd3d346a3d697362da94b9aa2dc2.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://www.sitedogta.com.br/gta5/imagens/personagens/Trevor%20GTA%20V.JPG" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/564x/3b/3b/56/3b3b56745376625aa66d5943b3db0275.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/originals/c8/9c/6b/c89c6b8f2165cfbe5ecccfebace1042d.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://i.pinimg.com/736x/ea/e7/b5/eae7b513060702e86bdd51d4d5cfc5ae.jpg" alt="">
  </div>
  <div class="child">
    <img src="https://cdn.hipwallpaper.com/i/94/92/Fk0l6I.jpg" alt="">
  </div>
</div>  


*{
  box-sizing: border-box;
}
body{
  padding:0; 
  margin: 0;
  background: #23232a;
}
img{
  width:100%;
  height: 100%;
  object-fit: cover;
  object-position: 40% 0;
}

.parent{
  padding: .8rem;
  background: black;
  height: 95vh;
  min-height: 500px;
  width: 100%;
  max-width: 600px;
  margin: auto;
  margin-top: 2.5vh;
  border: 1px solid #c9b473;
  overflow: hidden;
  display: grid;

  grid-template-columns: 1fr .7fr .3fr 1fr;
  grid-template-rows: 20% 40% 20% 20%;
  grid-template-areas: 'one two two three'
    'four five five five'
    'six five five five'
    'six seven eight eight';
}


// For clipping I used mozilla's inspect element. 
.child{

  &:nth-child(1),
  &:nth-child(2),
  &:nth-child(3){
    img{
      width:120%;
      height: 120%;
    }
  }
  &:first-child{
    grid-area: one;
    clip-path: polygon(0% 0%, 93.24% 0%, 105.04% 110.16%, 0% 90%);
  }
  &:nth-child(2){
    grid-area: two;
    clip-path: polygon(0% 0%, 108.28% 0%, 96.45% 110.13%, 10.55% 110.93%);
  }
  &:nth-child(3){
    grid-area: three;
    clip-path:polygon(15.05% 0%, 100% 0%, 99.35% 91.7%, 3.08% 108.48%);
  }
  &:nth-child(4){
    grid-area: four;
    clip-path: polygon(0% -0.85%, 106.34% 9.98%, 121.32% 65.63%, 99.66% 109.89%, 1.86% 124.41%);

    img{
      width: 135%;
      height: 135%;
    }
  }
  &:nth-child(5){
    grid-area: five;
    clip-path: polygon(6.4% 6.48%, 47.24% 5.89%, 100% 0%, 98.41% 96.85%, 53.37% 100%, 53% 63.21%, 3.23% 73.02%, 14.30% 44.04%);
  }
  &:nth-child(6){
    grid-area: six;
    clip-path:  polygon(2.14% 29.3%, 99.34% 15.42%, 98.14% 100.82%, 1.57% 101.2%);
  }
  &:nth-child(7){
    grid-area: seven;
    clip-path: polygon(7.92% 33.47%, 96.31% 23.39%, 95.38% 100%, 5.30% 100.85%);
  }
  &:nth-child(8){
    grid-area: eight;
    clip-path: polygon(2.5% 22.35%, 100% 0%, 100% 100%, 1.55% 100%);
  }
  &:nth-child(9){
    grid-row-start: 3;
    grid-row-end: 4;
    grid-column-start: 2;
    grid-column-end: 4; 

    clip-path:polygon(5.94% 28.66%, 100.61% -0.67%, 101.1% 108.57%, 5.4% 126.28%);

    img{
      object-position: 30% 50%;
      height: 135%;
    }
  }
}


```



####  css实现瀑布流布局（column-count）

关键点，

- column-count: 元素内容将被划分的最佳列数
- break-inside: 避免在元素内部插入分页符

```scss
// pug 模板引擎
div.g-container
    -for(var j = 0; j<32; j++)
        div.g-item



//scss
$count: 32;

@function randomNum($max, $min: 0, $u: 1) {
    @return ($min + random($max)) * $u;
}
//scss随机数函数
@function randomColor() {
    @return rgb(randomNum(255), randomNum(255), randomNum(255));
}

.g-container {
    column-count: 4;
    column-gap: .5vw;
    padding-top: .5vw;
}

.g-item {
    position: relative;
    width: 24vw;
    margin-bottom: 1vw;
    break-inside: avoid;
    //避免插入分页符
}

//生成随机的页标内容
@for $i from 1 to $count+1 {
    .g-item:nth-child(#{$i}) {
        height: #{randomNum(300, 50)}px;
        background: randomColor();

        &::after {
            content: "#{$i}";
            position: absolute;
            color: #fff;
            font-size: 2vw;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
        }
    }
}//生成随机的页标内容
```



####  box-shadow实现圆环进度条动画

- 圆环进度条的移动本质上是阴影顺序延时移动的结果。

```scss
<div class="container">
    <div class="shadow">Hover Me</div>
</div>



$color: #e91e63;

body {
    background: #000;
}

.container {
    position: relative;
    overflow: hidden;
    width: 124px;
    height: 124px;
    overflow: hidden;
    margin: 100px auto;
    border-radius: 50%;
}

.shadow {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 120px;
    height: 120px;
    line-height: 120px;
    border-radius: 50%;
    color: #fff;
    font-size: 20px;
    cursor: pointer;
    box-shadow: 60px -60px 0 2px $color, -60px -60px 0 2px $color,
        -60px 60px 0 2px $color, 60px 60px 0 2px $color;
    text-align: center;
    
    &:hover {
        animation: border .5s ease forwards;
    }
}

//核心的动画效果
@keyframes border{
  0% {
    box-shadow: 60px -60px 0 2px $color, -60px -60px 0 2px $color, -60px 60px 0 2px $color, 60px 60px 0 2px $color, 0 0 0 2px transparent;
  }
  25% {
    box-shadow: 0 -125px 0 2px $color, -60px -60px 0 2px $color, -60px 60px 0 2px $color, 60px 60px 0 2px $color, 0 0 0 2px #fff;
  }
  50% {
    box-shadow: 0 -125px 0 2px $color, -125px 0px 0 2px $color, -60px 60px 0 2px $color, 60px 60px 0 2px $color, 0 0 0 2px #fff;
  }
  75% {
    box-shadow: 0 -125px 0 2px $color, -125px 0px 0 2px $color, 0px 125px 0 2px $color, 60px 60px 0 2px $color, 0 0 0 2px #fff;
  }
  100% {
    box-shadow: 0 -125px 0 2px $color, -125px 0px 0 2px $color, 0px 125px 0 2px $color, 120px 40px 0 2px $color, 0 0 0 2px #fff;
  } 
}
```





####  纯css实现跑马灯效果

```html
<style>
/* 容器样式 */
.container {
  --s: 150px;   /* 控制logo大小 */
  --d: 8s;      /* 控制动画速度 */
  --n: 4;       /* 控制可见logo数量 */
  
  display: flex;
  overflow: hidden;
}

/* Logo图片样式 */
.container img {
  width: var(--s);
  offset-path: path("M0,0 L1000,0"); /* 水平运动路径 */
  animation: slide var(--d) linear infinite;
  animation-delay: calc(-1 * (var(--i)) * (var(--d) / 8));
    /*这里使用自定义属性--i为每个图片设置不同的延迟，形成连续流动效果。*/
}

/* 动画定义 */
@keyframes slide {
  0% {
    offset-distance: 0%;
  }
  100% {
    offset-distance: 100%;
  }
}
</style>
</head>
<body>
<div class="container">
  <img src="https://s2.loli.net/2025/09/07/D5hwT7Pen4fHF9N.jpg">
  <img src="https://s2.loli.net/2025/08/19/GP4Lma8gVe5ItlE.jpg">
  <!-- as many images as you want -->
</div>
</body>
```







####  由纯css实现滚动监视

```html


#table-of-contents a {
  transition: color 0.5s ease;
  
  &:target-current {
    color: red;
  }
}

.scroll-parent {
  scroll-behavior: smooth;
}<ol>
  <li><a href="#intro">Introduction</a></li>
  <li><a href="#ch1">Chapter 1</a></li>
  <li><a href="#ch2">Chapter 2</a></li>
</ol>

<div id="intro" class="chapter">Introduction content</div>
<div id="ch1" class="chapter">Chapter 1 content</div>
<div id="ch2" class="chapter">Chapter 2 content</div>


ol {
  background-color: gray;
  right: 10px;
  top: 10px;
  position: fixed;
  scroll-target-group: auto;
}

a:target-current {
  color: red;
}

.chapter {
  background: lightgray;
  height: 60vh;
  margin: 10px;
}

#table-of-contents a {
  transition: color 0.5s ease;
  
  &:target-current {
    color: red;
  }
}

.scroll-parent {
  scroll-behavior: smooth;
}

#table-of-contents {
  display: none;

  /* Display table of contents if scroll-target-group is supported */
  @supports(scroll-target-group: auto) {
    display: block;
    position: fixed;
  }
}
/*看是否符合浏览器支持 */
```

这里相关的代码行是：在包含“章节”的 `ol` 上使用 `scroll-target-group: auto;`，以及 [`:target-current`](https://developer.mozilla.org/en-US/docs/Web/CSS/:target-current) 伪类，你可能从 [CSS 轮播 API](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_overflow/CSS_carousels) 中认识它。这基本上告诉浏览器，`scroll-target-group` 内的链接列表应该像 `::scroll-marker` 一样处理（这也是来自 [CSS 轮播 API](https://developer.mozilla.org/en-US/docs/Web/CSS/::scroll-marker) 的——你能看出趋势吗？）。所以当链接的 `id` 作为与该内容区域的“滚动标记”相关时，你就可以用 `:target-current` 来样式化它。

`scroll-target-group` 丰富了 HTML 锚元素的函数，使其与 `::scroll-marker` 功能相匹配。正因如此，你现在对创建自定义标记有了更多的控制权（因为它们不再受伪元素限制）。这个演示只是 `scroll-target-group` 创意用法的一个例子，当然你也可以用它来创建更传统风格的轮播图。

“scroll-target-group 属性允许通过将 HTML 锚元素设为‘滚动标记’来克服这些限制。通过指定片段标识符，作者可以获得 ::scroll-markers 的‘滚动目标进入视图’功能，但无法获得‘跟踪当前滚动标记’的功能。使用 scroll-target-group 属性时，浏览器会运行算法来确定‘当前滚动标记’，作者可以用 :target-current 伪类来设置该锚元素。” -scroll-target-group 解释。''













