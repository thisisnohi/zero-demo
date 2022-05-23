package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	mikeUserFieldNames          = builder.RawFieldNames(&MikeUser{})
	mikeUserRows                = strings.Join(mikeUserFieldNames, ",")
	mikeUserRowsExpectAutoSet   = strings.Join(stringx.Remove(mikeUserFieldNames, "`create_time`", "`update_time`"), ",")
	mikeUserRowsWithPlaceHolder = strings.Join(stringx.Remove(mikeUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGoZeroMikeMobilePrefix = "cache:goZero:mikeUser:id:"
)

type (
	mikeUserModel interface {
		TransInsert(ctx context.Context, session sqlx.Session, data *MikeUser) (sql.Result, error)
		Insert(ctx context.Context, data *MikeUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*MikeUser, error)
		Update(ctx context.Context, data *MikeUser) error
		Delete(ctx context.Context, id int64) error
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
	}

	defaultMikeUserModel struct {
		sqlc.CachedConn
		table string
	}

	MikeUser struct {
		Id       int64  `db:"id"`
		Nickname string `db:"nickname"`
		Mobile   string `db:"mobile"`
	}
)

func newMikeUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMikeUserModel {
	return &defaultMikeUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`mike_user`",
	}
}

// TransInsert
/**
带事务的插入
  改动点: conn -> session
*/
func (m *defaultMikeUserModel) TransInsert(ctx context.Context, session sqlx.Session, data *MikeUser) (sql.Result, error) {
	goZeroMikeMobileKey := fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, mikeUserRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Id, data.Mobile, data.Nickname)
	}, goZeroMikeMobileKey)
	return ret, err
}

func (m *defaultMikeUserModel) Insert(ctx context.Context, data *MikeUser) (sql.Result, error) {
	goZeroMikeMobileKey := fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, mikeUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Mobile, data.Nickname)
	}, goZeroMikeMobileKey)
	return ret, err
}

func (m *defaultMikeUserModel) FindOne(ctx context.Context, id int64) (*MikeUser, error) {
	goZeroMikeMobileKey := fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, id)
	var resp MikeUser
	err := m.QueryRowCtx(ctx, &resp, goZeroMikeMobileKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", mikeUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
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
	goZeroMikeMobileKey := fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, mikeUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Mobile, data.Nickname, data.Id)
	}, goZeroMikeMobileKey)
	return err
}

func (m *defaultMikeUserModel) Delete(ctx context.Context, id int64) error {
	goZeroMikeMobileKey := fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, goZeroMikeMobileKey)
	return err
}

func (m *defaultMikeUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroMikeMobilePrefix, primary)
}

func (m *defaultMikeUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", mikeUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMikeUserModel) tableName() string {
	return m.table
}

// TransCtx 暴露给logic开启事务
func (m *defaultMikeUserModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}
