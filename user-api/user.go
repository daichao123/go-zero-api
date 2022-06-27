package main

import (
	"flag"
	"fmt"
	"go-zero-test/common"

	"go-zero-test/user-api/internal/config"
	"go-zero-test/user-api/internal/handler"
	"go-zero-test/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//test 使用全局中间件
	server.Use(common.NewCommonGlobalMiddleware("test middleware ...").Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
