# *** 小众点评

本项目使用go语言重构黑马点评项目，方便使用go语言的同学学习黑马程序员的redis课程，欢迎一起交流学习。  

## 介绍

因为[黑马程序员redis教程实战篇](https://www.bilibili.com/video/BV1cr4y1671t?p=24  )使用的语言是`java`，不想浪费这个项目所以想用golang重构一下。项目没有采用`gin`框架而是字节的[Hertz](https://www.cloudwego.io/zh/docs/hertz/)框架，目前已经完成了第一部分短信登录功能，**欢迎各位大佬一个合作完成这个项目**  

**没时间完整做的同学可以在issue中挑选一个模块完成。**  

### Start
#### 前端
前端代码在`resources/nginx-1.18.0.zip`中，Windows系统可以双击`nginx.exe`运行，Mac 或 Linux安装nginx后参考以下命令执行：
```shell
nginx -c ~/nginx-1.18.0/conf/nginx.conf -p ~/nginx-1.18.0
```

浏览器打开`http://127.0.0.1:8080`
* 使用该前端，发布博客要使用上传图片，需要配置`nginx.conf`增加以下配置：
```shell
        location /imgs {
            proxy_pass http://127.0.0.1:8080/imgs;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
            proxy_redirect off;
        }
```
#### 后端 
- 在mysql新建数据库表`xzdp`  
- 将`resources/xzdp.sql`导入到表`xzdp`  
- 启动`redis-server`    
- 复制`conf/test/conf.example.yaml`为`conf/test/conf.yaml`并修改其中的配置  
- `go run xzdp`

### 如何添加服务

（To 不熟悉Hertz的同学）  

1. 在idl目录下修改或添加[thrift](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/)文件

2. 修改makefile, 在update_api下添加

	`hz update --mod=xzdp --idl=idl/你的thrift文件 --customize_package=template/package.yaml`

3. `make update_api`

不要修改`model/SERVICE_NAME/SERVICE_NAME.go`的内容，因为会被覆盖。  

项目结构可能有点乱，**欢迎大佬来调整项目结构和修改`template`里的文件。**  

### 如何合作

欢迎以任何格式提交Issue和PR！或者➕我v：`lhpqaq1`. 有疑问也请联系我。   

目前默认分支是我的`dev`分支，请大家fork `master`分支。   

点个🌟吧 😘   



## introduce 

- Use the [Hertz](https://github.com/cloudwego/hertz/) framework
- Integration of pprof, cors, recovery, access_log, gzip and other extensions of Hertz.
- Generating the base code for unit tests.
- Provides basic profile functions.
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| hertz_gen  | Hertz generated model |
| biz/handler  | Used for request processing, validation and return of response. |
| biz/service  | The actual business logic. |
| biz/dal  | Logic for operating the storage layer |
| biz/route  | Routing and middleware registration |
| biz/utils  | Wrapped some common methods |

## How to run

```shell
sh build.sh
sh output/bootstrap.sh
```
