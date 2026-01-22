# *** 小众点评

本项目使用go语言重构黑马点评项目，方便使用go语言的同学学习黑马程序员的redis课程，欢迎一起交流学习。
[说明文档](./doc/xzdp.md)

## 介绍

> 当前代码经过 AI 大幅重构，如果出现问题，请切换至[human](https://github.com/lhpqaq/xzdp-go/tree/human)分支，也欢迎您提交代码修复bug☺️  

因为[黑马程序员redis教程实战篇](https://www.bilibili.com/video/BV1cr4y1671t?p=24  )使用的语言是`java`，所以想用golang重构一下。项目使用字节的[Hertz](https://www.cloudwego.io/zh/docs/hertz/)框架。

当前已经初步完成项目的基本功能，大家可以查看[issues](https://github.com/lhpqaq/xzdp-go/issues)中的需求或自行创建需求为项目提交代码，包括但不限于优化代码，添加文档，添加单元测试等。    

### Start
#### 前端
前端代码在`resources/nginx-1.18.0.zip`中，Windows系统可以双击`nginx.exe`运行，Mac 或 Linux 安装 nginx 后参考以下命令执行：
```shell
nginx -c ~/nginx-1.18.0/conf/nginx.conf -p ~/nginx-1.18.0
```

浏览器打开`http://127.0.0.1:8080`
* 使用该前端，发布博客要使用上传图片，需要配置`nginx.conf`增加以下配置：
```shell
        location /imgs {
            proxy_pass http://127.0.0.1:8081/imgs;
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

### TODO

- [ ] 优化代码

- [ ] Anything

	


### 如何合作

欢迎以任何格式提交 Issue 和 PR ！或者➕我v：`lhpqaq1`. 有疑问也请联系我。   

点个🌟吧 😘   

贡献指南：https://juejin.cn/post/7196940857308069945  



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
