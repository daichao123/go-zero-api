package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-test/model/model"
	"go-zero-test/user-api/internal/config"
	"go-zero-test/user-api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	UserMiddleware  rest.Middleware
	UserModel       model.UsersModel
	UserLevelsModel model.UserLevelsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//mysql 配置
		UserMiddleware:  middleware.NewUserMiddleware().Handle,
		UserModel:       model.NewUsersModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
		UserLevelsModel: model.NewUserLevelsModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
	}
}
