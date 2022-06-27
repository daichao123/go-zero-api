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
	userLoginLogFieldNames          = builder.RawFieldNames(&UserLoginLog{})
	userLoginLogRows                = strings.Join(userLoginLogFieldNames, ",")
	userLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(userLoginLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(userLoginLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	userLoginLogModel interface {
		Insert(ctx context.Context, data *UserLoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserLoginLog, error)
		Update(ctx context.Context, newData *UserLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserLoginLog struct {
		Id        int64          `db:"id"`      // ID
		UserId    int64          `db:"user_id"` // 用户ID
		Ips       sql.NullString `db:"ips"`     // IP
		Address   sql.NullString `db:"address"` // 地址
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
	}
)

func newUserLoginLogModel(conn sqlx.SqlConn) *defaultUserLoginLogModel {
	return &defaultUserLoginLogModel{
		conn:  conn,
		table: "`user_login_log`",
	}
}

func (m *defaultUserLoginLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserLoginLogModel) FindOne(ctx context.Context, id int64) (*UserLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userLoginLogRows, m.table)
	var resp UserLoginLog
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

func (m *defaultUserLoginLogModel) Insert(ctx context.Context, data *UserLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userLoginLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Ips, data.Address, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultUserLoginLogModel) Update(ctx context.Context, data *UserLoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userLoginLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Ips, data.Address, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultUserLoginLogModel) tableName() string {
	return m.table
}