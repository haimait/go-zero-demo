package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/user-rpc/usercenter"

	"go-zero-demo/model"
	"go-zero-demo/user-api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println("c.DB.DataSource:-->", c.DB.DataSource)
	fmt.Println("Timeout:-->", c.Timeout)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
