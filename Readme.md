# 微服务项目

## 系统架构

>go micro + gin + grpc

## 项目注意

>gprc版本太高会导致micro v2版本下载错误，go.mod 添加`replace google.golang.org/grpc => google.golang.org/grpc v1.26.0`
>由于micro v3版本在开发，所以最好不要拉去，然后>go get github.com/micro/protoc-gen-micro 不要加-u 否则拉去最新版本，报错
>go get  github.com/golang/protobuf/protoc-gen-go

## 配置文件说明
```yaml
{
    "dbConfig":{
        "driverName":"mysql",
        "server":"127.0.0.1",
        "port":3306,
        "user":"root",
        "password":"123456",
        "database":"test",
        "dateSourceName":"",
        "isDebug":true
    },
    "filePath":"/static/upload",
    "redisConfig":{
        "network":"",
        "addr":"127.0.0.1:6379",
        "password":"",
        "db":"0"
    },
    "txOcrApi":{
        "region":"ap-beijing",
        "endpoint":"ocr.tencentcloudapi.com",
        "secretId":"xxxxxx",
        "secretKey":"xxxxx"
    },
    "jwt":{
        "secret":"weixiao_keji_007",
        "TimeOut":30000000
    },
    "yzm":{
        "width":6,
        "TimeOut":1000000
    },
    "emailConfig":{
        "emailType":0,
        "userName":"xxx@qq.com",
        "password":"xxxx",
        "bcc":[

        ],
        "cc":[

        ],
        "to":[

        ],
        "subject":"weixiaoqaq",
        "text":"",
        "html":""
    },
    "services":{
        "userSrv":{
            "etcd":"",
            "name":"com.weixiao.userSrv",
            "version":"latest",
            "ip":"",
            "port":0
        },
        "userWeb":{
            "etcd":"",
            "name":"com.weixiao.userWeb",
            "version":"latest",
            "ip":"",
            "port":8701
        },
        "sendSrv":{
            "etcd":"",
            "name":"com.weixiao.sendSrv",
            "version":"latest",
            "ip":"",
            "port":8702
        },
        "sendWeb":{
            "etcd":"",
            "name":"com.weixiao.sendWeb",
            "version":"latest",
            "ip":"",
            "port":8703
        },
        "fileSrv":{
            "etcd":"",
            "name":"com.weixiao.fileSrv",
            "version":"latest",
            "ip":"",
            "port":8704
        },
        "fileWeb":{
            "etcd":"",
            "name":"com.weixiao.fileWeb",
            "version":"latest",
            "ip":"",
            "port":8705
        }
    }
}
```
## 微服务划分

### 用户服务（sysuser）

~~protoc --go_out=plugins=grpc:. --micro_out=. -I=proto/sysuser ./proto/sysuser/*.proto~~

`protoc --go_out=plugins=grpc:. --micro_out=. -I=proto/sysuser ./proto/sysuser/sysuser.proto`

请在powershell模式下执行命令`Get-ChildItem *.proto |Resolve-Path -Relative | %{protoc $_ --go_out=. --micro_out=.}`这样就可以使用*.proto了


## 消息服务（send）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/send/send.proto`

## 文件服务（file）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/file/file.proto`
## 图片webp服务（第三方服务，将图片转成webp压缩）

###本项目中我直接下载了源码，然后build,下面介绍docker中如何部署
1. `docker pull webpsh/webps`

2. 找到图片资源的文件夹 

3. 创建容器`docker run -d -p 3333:3333 -v /path/to/pics:/opt/pics --name webps webpsh/webps` 
4. 我的是window.`docker run -d -p 3333:3333 -v E/go/qshapi/static/upload:/opt/pics --name webps webpsh/webps` 挂在文件夹
5. 测试：`docke ps` 查看是否有容器webps,然后在浏览器中输入图片地址`http://localhost:3333/no.gif`查看
### 创建nginx反向代理,处理图片
1. [安装nginx](https://github.com/943885179/dockerStu/blob/master/docker_nginx.md)
    1. `docker pull nginx`
    2. `docker run --name dockernginx -d -p 8080:80 -v e/docker/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v e/docker/nginx/www:/usr/share/nginx/html -v e/docker/nginx/logs:/var/log/nginx nginx` 
    3. 配置nginx.conf 然后重启Nginx`docker exec -it dockernginx service nginx reload`
         ```conf
            user  nginx;
            worker_processes  1;
            
            error_log  /var/log/nginx/error.log warn;
            pid        /var/run/nginx.pid;
            
            
            events {
                worker_connections  1024;
            }
            
            
            http {
                fastcgi_buffers 8 16k;
                fastcgi_buffer_size 32k;
                fastcgi_connect_timeout 300;
                fastcgi_send_timeout 300;
                fastcgi_read_timeout 300;
                include       /etc/nginx/mime.types;
                default_type  application/octet-stream;
            
                log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                                  '$status $body_bytes_sent "$http_referer" '
                                  '"$http_user_agent" "$http_x_forwarded_for"';
            
                access_log  /var/log/nginx/access.log  main;
            
                sendfile        on;
                #tcp_nopush     on;
            
                keepalive_timeout  65;
            
                gzip  on;
                # include /etc/nginx/conf.d/*.conf; # 这个要注释掉，不然它将使用conf.d下的default.conf文件作为配置文件
               server {
                    listen  80;
                    server_name  img.weixiaoqaq.com;
                    root /usr/share/nginx/html;
                    index login.html;
                    location ~ .*\.(gif|jpg|jpeg|png)$ {  
                        proxy_pass  http://192.168.0.9:3333; # 必须要用特定ip否则好像有错误
                    }
                    location / {  
                       index  login.html login.htm; #html文件名称
                    }
               } 
            }
          ```
    4. 范围是否配置成功代理：[http://localhost:8080/1.jpg](http://localhost:8080/1.jpg)
## 三方api服务注册到micro中（三方服务注册到api）,下面是测试用webp注入，但是它是个非api接口程序，所以
1. 启动`micro registry`
2. 启动`micro web`
3. 查看micro registry服务，发现有registry服务，选择service看到有个reigistry方法，赋值需要传入的参数，在client调用该接口，请求参数为
    ```json
    {
            "name": "com.weixiao.imgCli",
            "version": "1.0",
            "endpoints": [],
            "nodes": [{
                "address": "192.168.0.9:3333",
                "id": "userservice-img"
            }]
        }
    ```
4. 注册服务后如果成功则返回{}，失败返回失败内容，然后再去service去看是否有该服务生成
6. 备注(如果使用etcd的话需要先指定使用etcd才能注册到etcd中)
    1. `set micro_registry=etcd`
    2. `set micro_registry_address=127.0.0.1:6379`
##启动micro网关
1. 指定服务的命名空间为自己的命名空间`micro web --namespace com.weixiao.web `特别注意：web程序必须要以xxx.xxx.web.xxx命名，否则虽然注册到服务中但是不能直接调用
2. 使用micro api 启动api同样如初，命名规范必须是[xxx...].api.xxx
```cmd
# 启动服务
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=localhost:2379
go run srv/sysuser/main.go --server_address :9090
# web管理界面启动
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=localhost:2379
micro web
# api网关启动
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=localhost:2379
set MICRO_API_NAMESPACE=xxx.xxx.api.xxx
set MICRO_CLIENT=grpc
set MICRO_SERVER=grpc
micro api --handler=rpc
```