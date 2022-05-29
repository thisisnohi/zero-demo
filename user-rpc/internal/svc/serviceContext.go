package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/user-api/model"
	"zero-demo/user-rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	MikeUserModel     model.MikeUserModel
	MikeUserDataModel model.MikeUserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		MikeUserModel:     model.NewMikeUserModel(conn, c.CacheRedis),
		MikeUserDataModel: model.NewMikeUserDataModel(conn, c.CacheRedis),
	}
}
