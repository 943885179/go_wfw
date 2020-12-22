.PHONY:run
run:
	# 本地启动
	go run main.go
.PHONY:build
.PHONY:pro
pro:
	protoc  --go_out=plugins=grpc:. --micro_out=. -I=../../proto/basic ../../proto/basic/basic.proto
build:
	# 打包linux,部署到docker使用
	set GOOS=linux
	set GOARCH=amd64
	go build -o build/basic
.PHONY:buildexe
buildexe:
	# 打包exe
	set GOOS=windows
	set GOARCH=amd64
	go build -o build/basic.exe
.PHONY:docker
docker:
	#镜像构建，确保本目录下有Dockerfile文件
	docker build -t weixiao.basic:v1 .
.PHONY:dockerrun
dockerrun:
	docker run -it -d --name weixiao.basic -p 8001:8001 weixiao.basic:v1
