#  tailwindCss

![75440889944](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754408899449.png)



主流css样式实现方案

本质上是在可维护性，开发效率，性能去寻找平衡

1. css和预处理器

   是最直接的处理方式，但是容易冗余和缺乏逻辑是大缺点

2. cssinjs：emotion这个方案

   ![75441022297](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754410222973.png)

3. 原子化css

   ![75441038225](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754410382251.png)

   使用编译的思想完成样式的自动化 

   优势在于不要用思考类名，不用切换文件，性能好，约束与一致性，所有预设样式来自一个配置文件



安装tailwind

![75441092963](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754410929632.png)

然后在css引入

![75441094385](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754410943854.png)

tailwind4和tailwind3差距非常大

![75441100453](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754411004534.png)

还要去安装一个插件用于解析tailwind

在tailwind里面是不允许写其他的

一个vscode插件，提示相关的

![75441192607](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754411926076.png)





使用tailwind有一些核心概念需要学习

![75441206230](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754412062307.png)

颜色体系

![75441296272](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754412962722.png)





自定义颜色相关

![75441299844](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754412998446.png)

![75441301977](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754413019774.png)

不建议过渡使用@apply这个组件

然后直接在类名可以使用btn-primary这个类名

![75444846649](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754448466495.png)

使用@theme和@utility的方式引入进来

![75444876602](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754448766027.png)

 ![75444927220](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754449272207.png)

![75444931272](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754449312723.png)

![75444931974](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754449319749.png)

![75444947662](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754449476625.png)

白色模式就是用黑色，黑暗模式用白色、







####  架构下的优势

![75445034602](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754450346020.png)

![75445036522](C:\Users\zxh\Desktop\前端\css\tailwindcss.assets\1754450365221.png)







