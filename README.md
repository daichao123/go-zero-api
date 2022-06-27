# go-zero-api
Go-zero microservices framework API for individual services 
1. Functions include goCTL scaffold use 
  Go 1.15 及之前版本:
  GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest
  Go 1.16 及以后版本:
  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
  2.goctl 常用命令:
  goctl -v ------打印版本

  goctl rpc go --help  生成rpc 文件

  goctl api validate  验证api 文件语法是否有错误

  goctl api go -api user.api -dir ./ -style goZero 生成api文件命令 具体参数参考文档
  goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/user_service" -table="*"  -dir="./model" 生成model文件命令 具体参数参考文档

3. Database operation and transaction usage 
4. Configure individual services and use logs  
5. Use of local middleware and global middleware
