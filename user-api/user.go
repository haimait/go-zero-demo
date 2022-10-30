package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/common/errorx"
	"go-zero-demo/user-api/internal/config"
	"go-zero-demo/user-api/internal/handler"
	"go-zero-demo/user-api/internal/svc"
	"net/http"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	//读取控制台启动命令时带的参数 例 下面命令里-f 带的参数
	//go run user.go -f etc/greet-api.yaml
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) //yaml配置初使化到config里

	ctx := svc.NewServiceContext(c)          //初使化ServiceContext配置文件
	server := rest.MustNewServer(c.RestConf) //启一个server
	defer server.Stop()                      //退出时，服务停时停止服务

	//注册路由
	handler.RegisterHandlers(server, ctx)
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		//处理errors.New生成的错误，支持返回json格式， 兼容go-zero入参校验不过的情况
		if fmt.Sprintf("%T", err) == "*errors.errorString" {
			err = errorx.NewDefaultError(err.Error())
		}

		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}

	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//启动服务
	server.Start()

}
