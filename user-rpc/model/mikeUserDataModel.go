package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MikeUserDataModel = (*customMikeUserDataModel)(nil)

type (
	// MikeUserDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMikeUserDataModel.
	MikeUserDataModel interface {
		mikeUserDataModel
	}

	customMikeUserDataModel struct {
		*defaultMikeUserDataModel
	}
)

// NewMikeUserDataModel returns a model for the database table.
func NewMikeUserDataModel(conn sqlx.SqlConn, c cache.CacheConf) MikeUserDataModel {
	return &customMikeUserDataModel{
		defaultMikeUserDataModel: newMikeUserDataModel(conn, c),
	}
}
