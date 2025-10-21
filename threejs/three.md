#  threejs

####  引入

1. js引入
2. cdn链接
3. npm




####  最简单的three

场景,物体,灯光,摄像机,渲染器组成一个最简单的场景

```javascript
import * as THREE from 'three'

const w=window.innerWidth
const h=window.innerHeight
//房间-3d容器
//Scene场景
const scene=new THREE.Scene()
//东西-家具家电
//物体:geometry几何体+material材质
const geometry=new THREE.BoxGeometry(1,1,1)
const material=new THREE.MeshBasicMateral()
const cube = new THREE.Mesh(geometry,material)
scene.add(cube)
//对物体进行操作
//Position
cube.position.x=-1
cube.position.y=-1
cube.position.z=-1
cube.position.set(1,1,1,)
//Routation
cube.rotation.z=45/100*Math.PI
cube.rotation.y=45/100*Math.PI
//Scale
cube.scale.x=2
cube.scale.y=2

//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
//默认画布大小是300*150
renderer.render(scene,camera)

//把dom添加上去
document.body.append(renderer.domElemnent)
```

norubias.top

 



####  组成three

![75006304986](C:\Users\zxh\Desktop\前端\threejs\three.assets\1750063049864.png)

renderer.domElemnent是一个canvas，使用dom操作把它发送到html上去

 



####  坐标系与变换操作

变换，位置，旋转，缩放，在threejs里面对他的属性进行操作

 在上面的案例有相关的操作



####  运动的相关操作
```javascript
import * as THREE from 'three'

const w=window.innerWidth
const h=window.innerHeight
//房间-3d容器
//Scene场景
const scene=new THREE.Scene()
//东西-家具家电
//物体:geometry几何体+material材质
const geometry=new THREE.BoxGeometry(1,1,1)
const material=new THREE.MeshBasicMateral()
const cube = new THREE.Mesh(geometry,material)
scene.add(cube)
//对物体进行操作
//Position
cube.position.x=-1
cube.position.y=-1
cube.position.z=-1
cube.position.set(1,1,1,)
//Routation
cube.rotation.z=45/100*Math.PI
cube.rotation.y=45/100*Math.PI
//Scale
cube.scale.x=2
cube.scale.y=2

//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
//默认画布大小是300*150
renderer.render(scene,camera)

//把dom添加上去
document.body.append(renderer.domElemnent)




//物体与运动
//使用setInteral
setInterval(()=>{
    cube.rotation.z+=0.01
    renderer.render(scene,camera)
},1000/60)
//对性能不友好

//使用requestArimationFrame
//window下面的方法
function tick(){
    cube.rotation.z+=0.01
    renderer.render(scene,camera)
    requestAnimationFrame(tick)
    //尽可能按照刷新率刷新
}//动画基本上都是用这个
tick()
//刷新率不同,刷新率大的,转的比较快
//这样是不行的


//解决刷新率问题
let time =Data.now()
function tick(){
    const currentTime=Date.now()
    const dataTime=currenTime-time
    time=currentTime
    cube.rotation.z+=dataTime*0.001
    renderer.render(scene,camera)
    requestAnimationFrame(tick)
    //每次增加一个固定值,
}
tick()

//THREE自带的解决刷新率问题的函数
const clock = new THREE.clock()
tick()
function tick(){
    const time=clock.getElapsedTime()
    cube.rotation.z=time
    cube.position.x=Math.sin(time*2)*2
    cube.position.y=Math.cos(time*2)*2
    //旋转的代码
    //time是一个不断均匀增加的值
    requestAnimationFrame(tick)
    renderer,render(scene,camera)
}

```







####  帧率检测

可以检测帧率的大小

```javascript
import * as THREE from 'three'

const w=window.innerWidth
const h=window.innerHeight
//房间-3d容器
//Scene场景
const scene=new THREE.Scene()
//坐标轴
const axes=new THREE.AxesHelper(2,2,2)
scene.add(axes)

//不把物体加到场景中
let cubes=[]
const stat=new Stat()
//帧率检查

//多个物体
function createCube(){
    //东西-家具家电
   //物体:geometry几何体+material材质
   const geometry=new THREE.BoxGeometry(1,1,1)
   //大小也可以随机
   const material=new THREE.MeshBasicMateral({
       color:0xfffff*Math.random()
       //随机颜色
   })
   const cube = new THREE.Mesh(geometry,material)]
   cube.postion.x=(Math.random()-0.5)*4;
   cube.postion.y=(Math.random()-0.5)*4;
   cube.postion.z=(Math.random()-0.5)*4;
   cubes.push(cube)
}
createCube()
//这样创建多个会方便很多

//
for(let i =0;i<20;i++){
    createCube()
}

cubes.forEach(cube=>{
    scence.add(cube)
})


//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
//默认画布大小是300*150
renderer.render(scene,camera)
document.body.append(renderer.domElement)
document.body.append(stat.dom)
//添加FPS检测

//运动
//THREE自带的解决刷新率问题的函数
const clock = new THREE.clock()
tick()
function tick(){
    const time=clock.getElapsedTime()
    
    cubes.forEach((cube,index)=>{
        cube.rotation.x=time*0.2+index
        cube.rotation.y=time*0.2+index
    })
    //旋转的代码
    //只要使用两个轴就能任意方向转
    //time是一个不断均匀增加的值
    requestAnimationFrame(tick)
    renderer.render(scene,camera)
    stat.update()
}

//把dom添加上去
document.body.append(renderer.domElemnent)
```




####  Contorls鼠标交互

在3d场景进行操作,左键旋转,已经右键平移和放大缩小等等操作

```javascript
import * as THREE from 'three'
import Stat from 'three/examples/jsm/lib/stats.module'
import {OrbitControls} from 'three/exmaples/jsm/Controls/OrbitControls'

const w=window.innerWidth
const h=window.innerHeight
const stat=new Stat()


//Scene场景
const scene=new THREE.Scene()

const g=new THREE.BoxGeometry(1,1,1)
const m=new THREE.MeshNormalMaterial()
//法线材质,每个面都不一样 
const cube=new THREE.Mesh(g,h)
scence.add(cube)

//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
document.body.append(renderer.domElement)
document.body.append(stat.dom)


const orbitControls=new OrbitControls(camera,renderer.domElement)

const clock = new THREE.clock()
tick()
function tick(){
    const time=clock.getElapsedTime()
    //time是一个不断均匀增加的值
    requestAnimationFrame(tick)
    renderer.render(scene,camera)
    stat.update()
    //更新统计消息
    orbitControls.update()
    //左键旋转,已经右键平移和放大缩小
}

//把dom添加上去
document.body.append(renderer.domElemnent)
```








####  物体的组合运动
```javascript
import * as THREE from 'three'
import Stat from 'three/examples/jsm/lib/stats.module'
import {OrbitControls} from 'three/exmaples/jsm/Controls/OrbitControls'

const w=window.innerWidth
const h=window.innerHeight
const stat=new Stat()
//Scene场景
const scene=new THREE.Scene()

const g=new THREE.BoxGeometry(1,1,1)
const m=new THREE.MeshNormalMaterial()
//法线材质,每个面都不一样 
const cube=new THREE.Mesh(g,h)
scence.add(cube)

const axes = new THREE.AxesHelper(2,2,2)
scene.add(axes)

//组合物体
const group=new THREE.Group()
//组合可以放到组合里面去,再分别加运动也可以


const geometry=new THREE.BoxGeometry(1,1,1)
const material=new THREE.MeshBasicMaterial({color:0xff0000})
const cubel = new THREE.Mesh(geometry,material)
group.add(cubel)

cuvel.position.y=1.5

const geometry2=new THREE.BoxGeometry(1,1,1)
const material2=new THREE.MeshBasicMaterial({color:0xff0000})
const cubel2 = new THREE.Mesh(geometry,material)
group.add(cubel2)


const geometry3=new THREE.BoxGeometry(1,1,1)
const material3=new THREE.MeshBasicMaterial({color:0xff0000})
const cubel3 = new THREE.Mesh(geometry,material)
group.add(cubel3 )

cuvel.position.y=-1.5


scene.add(group)
//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
document.body.append(renderer.domElement)
document.body.append(stat.dom)


const orbitControls=new OrbitControls(camera,renderer.domElement)

const clock = new THREE.clock()
tick()
function tick(){
    const time=clock.getElapsedTime()
    //time是一个不断均匀增加的值
    
    group.rotation.z=time
    //整个组合都会在运动
    
    requestAnimationFrame(tick)
    renderer.render(scene,camera)
    stat.update()
    //更新统计消息
    orbitControls.update()
    //左键旋转,已经右键平移和放大缩小
}

//把dom添加上去
document.body.append(renderer.domElemnent)
```

![75023801663](C:\Users\zxh\Desktop\前端\threejs\three.assets\1750238016631.png)

完成一个行星模型

![75023805985](C:\Users\zxh\Desktop\前端\threejs\three.assets\1750238059859.png)







####  例子,车子运动



```javascript
import * as THREE from 'three'
import Stat from 'three/examples/jsm/lib/stats.module'
import {OrbitControls} from 'three/exmaples/jsm/Controls/OrbitControls'

const w=window.innerWidth
const h=window.innerHeight
const stat=new Stat()
//Scene场景
const scene=new THREE.Scene()

const g=new THREE.BoxGeometry(1,1,1)
const m=new THREE.MeshNormalMaterial()
//法线材质,每个面都不一样 
const cube=new THREE.Mesh(g,h)
scence.add(cube)

const axes = new THREE.AxesHelper(2,2,2)
scene.add(axes)

//组合物体汽车
const car=new THREE.Group()
//组合可以放到组合里面去,再分别加运动也可以

//车身
const body = new THREE.Group()
const material=new THREE.MeshNormalMater()
const bodyCubel=new THREE.Mesh(
    new THREE.BoxGeometry(1,2,0.5)
    material
)
const bodyCube2=new THREE.Mesh(
    new THREE.BoxGeometry(0.5,0.5,0.5)
    new THREE.MeshbasicMaterial()
)
bodyCube2.position.z=0.5

body.add(bodyCubel)
body.add(bodyCube2)
car.add(body)

//轮子
const wheelGroup1=new THREE.Group()
const wheel1=new RHREE.Mesh({
    new THREE.BoxGeometry(0.1,0.7,0.7)
    material
})
wheelGroup1.position.set(-0.7,0.6,0)
wheelGroup1.add(wheel1)
car.add(wheelGroup1)

scene.add(car)
//光线-灯
const light=new THREE.AmbientLight()
scene.add(light)
//相机-
const camera=new THREE.perspectiveCamera(75,w/h,0.1,100)
camera.position.ser(0,0,5)
camera.lookAt(0,0,0)
//相机在0,0,5拍摄0,0,0位置的东西
//Renderer渲染器
const renderer=new THREE.WebGLRenderer()
renderer.setSize(w,h)
document.body.append(renderer.domElement)
document.body.append(stat.dom)

const orbitControls=new OrbitControls(camera,renderer.domElement)
//鼠标操控
const clock = new THREE.clock()
tick()
function tick(){
    const time=clock.getElapsedTime()
    //time是一个不断均匀增加的值
    
    group.rotation.z=time
    //整个组合都会在运动
    
    wheelGroup1.rotation.y=time
    //轮子自转
    
    car.position.y=time
    //车子转动
    
    
    
    requestAnimationFrame(tick)
    renderer.render(scene,camera)
    stat.update()
    //更新统计消息
    orbitControls.update()
    //左键旋转,已经右键平移和放大缩小
}

//把dom添加上去
document.body.append(renderer.domElemnent)
```















