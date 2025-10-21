####  webgl

浏览器支持的一个3d引擎

![76101418460](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014184603.png)

资料

![76101438405](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014384050.png)





####  基础

1. 坐标系

   ![76101454215](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014542150.png)

   canvas的坐标系

   ![76101459162](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014591628.png)

   ![76101480136](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014801369.png)

2. webgl的渲染管线

   一个工厂模式的渲染器、

   ![76101482981](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761014829811.png)

   点数据给顶点着色器，然后给图源装配，告诉电脑要画什么东西，通过公式进行处理，然后使用光栅器转化成像素图形，把每个东西都弄成一个小块，再对小片进行着色，加颜色纹理

   最后处理把它绘制出来

3. webgl的关键名词

   顶点着色器

   ![76101537112](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761015371121.png)

   图元装配

   ![76101548568](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761015485686.png)

   光栅化

   ![76101551544](C:\Users\zxh\Desktop\前端\threejs\webgl.assets\1761015515441.png)

4. 例子

   ```html
   <script src="引入一个坐标系处理文件"></script>
   <script>
       function init(){
           //入口函数
       }
       function initWebgl(){
           //初始化函数
       }
       function initShader(){
           //shader初始化函数
       }
       function initBuffer(){
           //缓存区函数
       }
       function draw(){
           
       }
   </script>
   <body>
       <canvas id="webglCanvas" width="500" height="500">
       </canvas>
   </body>
   ```

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​

   ​