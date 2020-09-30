# 微服务项目

## 系统架构

>go micro + gin + grpc

## 项目注意

>gprc版本太高会导致micro v2版本下载错误，go.mod 添加`replace google.golang.org/grpc => google.golang.org/grpc v1.26.0`
>由于micro v3版本在开发，所以最好不要拉去，然后>go get github.com/micro/protoc-gen-micro 不要加-u 否则拉去最新版本，报错
>go get  github.com/golang/protobuf/protoc-gen-go
## 微服务划分

### 用户服务（sysuser）

~~protoc --go_out=plugins=grpc:. --micro_out=. -I=proto/sysuser ./proto/sysuser/*.proto~~

`protoc --go_out=plugins=grpc:. --micro_out=. -I=proto/sysuser ./proto/sysuser/sysuser.proto`

## 消息服务（send）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/send/send.proto`

## 文件服务（file）
`protoc --go_out=plugins=grpc:. --micro_out=.  ./proto/file/file.proto`