package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	// MikeUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMikeUser
	MikeUserModel interface {
		mikeUserModel
	}

	customMikeUserModel struct {
		*defaultMikeUserModel
	}
)

// NewMikeUserModel returns a model for the database table.
func NewMikeUserModel(conn sqlx.SqlConn, c cache.CacheConf) *customMikeUserModel {
	return &customMikeUserModel{
		defaultMikeUserModel: newMikeUserModel(conn, c),
	}
}
