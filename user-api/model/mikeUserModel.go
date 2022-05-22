package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MikeUserModel = (*customMikeUserModel)(nil)

type (
	// MikeUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMikeUserModel.
	MikeUserModel interface {
		mikeUserModel
	}

	customMikeUserModel struct {
		*defaultMikeUserModel
	}
)

// NewMikeUserModel returns a model for the database table.
func NewMikeUserModel(conn sqlx.SqlConn) MikeUserModel {
	return &customMikeUserModel{
		defaultMikeUserModel: newMikeUserModel(conn),
	}
}
