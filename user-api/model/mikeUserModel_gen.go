// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	mikeUserFieldNames          = builder.RawFieldNames(&MikeUser{})
	mikeUserRows                = strings.Join(mikeUserFieldNames, ",")
	mikeUserRowsExpectAutoSet   = strings.Join(stringx.Remove(mikeUserFieldNames, "`create_time`", "`update_time`"), ",")
	mikeUserRowsWithPlaceHolder = strings.Join(stringx.Remove(mikeUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	mikeUserModel interface {
		Insert(ctx context.Context, data *MikeUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*MikeUser, error)
		Update(ctx context.Context, data *MikeUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultMikeUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	MikeUser struct {
		Id       int64  `db:"id"`
		UserId   int64  `db:"user_id"`
		Nickname string `db:"nickname"`
	}
)

func newMikeUserModel(conn sqlx.SqlConn) *defaultMikeUserModel {
	return &defaultMikeUserModel{
		conn:  conn,
		table: "`mike_user`",
	}
}

func (m *defaultMikeUserModel) Insert(ctx context.Context, data *MikeUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, mikeUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.UserId, data.Nickname)
	return ret, err
}

func (m *defaultMikeUserModel) FindOne(ctx context.Context, id int64) (*MikeUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", mikeUserRows, m.table)
	var resp MikeUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMikeUserModel) Update(ctx context.Context, data *MikeUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, mikeUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Nickname, data.Id)
	return err
}

func (m *defaultMikeUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultMikeUserModel) tableName() string {
	return m.table
}
