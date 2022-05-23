package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	TestMiddleware    rest.Middleware
	MikeUserModel     model.MikeUserModel
	MikeUserDataModel model.MikeUserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		TestMiddleware:    middleware.NewTestMiddleware().Handle,
		MikeUserModel:     model.NewMikeUserModel(conn, c.CacheRedis),
		MikeUserDataModel: model.NewMikeUserDataModel(conn, c.CacheRedis),
	}
}
