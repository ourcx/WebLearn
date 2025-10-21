# css布局

####  grid布局---网格

![75134847318](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751348473181.png)

flex是轴线布局

1. 容器-有容器属性

2. 项目-有项目属性

   ![75134862893](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751348628938.png)

基本内容

![75134868718](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751348687187.png)

**先确定行列**

![75135009154](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751350091548.png)

![75135041641](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751350416417.png)

使用百分比进行区分

 ```css
.main{
    grid-template-colums: 1fr minmax(150px,1fr);
    //分成份和线
    grid-template-rows:100px 100px 100px 100px
}
 ```

![75135080560](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751350805600.png)

```css
.main{
    grid-template-colums: 1fr minmax(150px,1fr);
    //分成份和线
     grid-template-colums: [c1] 100px [c2] 100px 100px [c3];
    //网格线
    grid-template-rows:100px 100px 100px 100px;
    row-gap:20px;
    colums-gap:20px;
}
```

![75135112978](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751351129788.png)

```scss
.main{
    grid-template-colums: 1fr minmax(150px,1fr);
    //分成份和线
     grid-template-colums: [c1] 100px [c2] 100px 100px [c3];
    //网格线
    grid-template-rows:100px 100px 100px 100px;
    row-gap:20px;
    colums-gap:20px;
    grid-tenmplate-areas:'a b c' 'a b c' 'a b c';
    //区域的使用
    grid-auto-flow: colum;
    //方向
}
```

![75135155037](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751351550375.png)

![75135155890](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751351558907.png)

dense是将其补上的东西

```scss
.main{
    grid-template-colums: 1fr minmax(150px,1fr);
    //分成份和线
     grid-template-colums: [c1] 100px [c2] 100px 100px [c3];
    //网格线
    grid-template-rows:100px 100px 100px 100px;
    row-gap:20px;
    colums-gap:20px;
    grid-tenmplate-areas:'a b c' 'a b c' 'a b c';
    //区域的使用
    grid-auto-flow: colum;
    //方向
    jusify-items:center;
    //防止网格中间,所有的,水平方向
    align-items:center;
    //垂直的
    justify-content:center;
    //内容的居中
}
```

![75135187956](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751351879561.png)

![75135195973](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751351959731.png)

```scss
.main{
    grid-template-colums: 1fr minmax(150px,1fr);
    //分成份和线
     grid-template-colums: [c1] 100px [c2] 100px 100px [c3];
    //网格线
    grid-template-rows:100px 100px 100px 100px;
    row-gap:20px;
    colums-gap:20px;
    grid-tenmplate-areas:'a b c' 'a b c' 'a b c';
    //区域的使用
    grid-auto-flow: colum;
    //方向
    jusify-items:center;
    //防止网格中间,所有的,水平方向
    align-items:center;
    //垂直的
    justify-content:center;
    //内容的居中
    grid-auto-colimns:50px;
    grid-auto-rows:50px;
    //设置多出来项目的宽高
}
```

![75135219583](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751352195838.png)

![75135252986](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751352529863.png)

设立网格线开始与结束的地方

```scss
grid-colum: 1 / 3;
//简写方式

```



```scss
.item{
    gird-area: b;
    //设置区域
}
```

![75135282073](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751352820730.png)

```scss
item{
    grid-area:1/1/1/2;
    //简写线区域
}
```

![75135285433](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751352854330.png)



![75135297438](C:\Users\zxh\Desktop\前端\css\css布局方式.assets\1751352974383.png)

作用于自己的位置属性





####  媒体查询布局@media

根据页面变化进行布局的变化

```css
@media screen and (max-width:560px){
    .menu{
        disaply:block
    }
    .nav{
        disaply:none
    }
}
大于的时候就隐藏菜单,小于就显示,做动态变化,注意样式覆盖问题
去覆盖掉原来的样式,相当于两套

可以加动画,因为它的元素属性变化了
```





















