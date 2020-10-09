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

## 消息服务（send）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/send/send.proto`

## 文件服务（file）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/file/file.proto`