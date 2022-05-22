package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	MikeUserModel model.MikeUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		MikeUserModel: model.NewMikeUserModel(conn),
	}
}
