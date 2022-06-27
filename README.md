# go-zero-api 学习笔记
Go-zero microservices framework API for individual services 
* Functions include goCTL scaffold use:
  Go 1.15 及之前版本:
  GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest
  Go 1.16 及以后版本:
  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
* goctl 常用命令:
  1. goctl -v ------打印版本
  2. goctl rpc go --help  生成rpc 文件
  3. goctl api validate  验证api 文件语法是否有错误
  4. goctl api go -api user.api -dir ./ -style goZero 生成api文件命令 具体参数参考文档
  5. goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/user_service" -table="*"  -dir="./model" 生成model文件命令 具体参数参考文档
* Database operation and transaction usage 
* Configure individual services and use logs  
* Use of local middleware and global middleware
