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
	adminFieldNames          = builder.RawFieldNames(&Admin{})
	adminRows                = strings.Join(adminFieldNames, ",")
	adminRowsExpectAutoSet   = strings.Join(stringx.Remove(adminFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	adminRowsWithPlaceHolder = strings.Join(stringx.Remove(adminFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	adminModel interface {
		Insert(ctx context.Context, data *Admin) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Admin, error)
		Update(ctx context.Context, newData *Admin) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Admin struct {
		Id           int64          `db:"id"`            // 自增id
		Username     string         `db:"username"`      // 用户名
		Password     string         `db:"password"`      // 密码
		Email        string         `db:"email"`         // email
		Mobile       string         `db:"mobile"`        // mobile
		Encrypt      string         `db:"encrypt"`       // 加密因子
		GoogleSecret sql.NullString `db:"google_secret"` // 谷歌验证密钥
		SellerId     int64          `db:"seller_id"`     // 商家id
		RoleId       int64          `db:"role_id"`       // 角色id
		CreatedAt    sql.NullTime   `db:"created_at"`
		UpdatedAt    sql.NullTime   `db:"updated_at"`
		Status       int64          `db:"status"` // 状态（-1删除，0禁用，1正常）
	}
)

func newAdminModel(conn sqlx.SqlConn) *defaultAdminModel {
	return &defaultAdminModel{
		conn:  conn,
		table: "`admin`",
	}
}

func (m *defaultAdminModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAdminModel) FindOne(ctx context.Context, id int64) (*Admin, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRows, m.table)
	var resp Admin
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

func (m *defaultAdminModel) Insert(ctx context.Context, data *Admin) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Email, data.Mobile, data.Encrypt, data.GoogleSecret, data.SellerId, data.RoleId, data.CreatedAt, data.UpdatedAt, data.Status)
	return ret, err
}

func (m *defaultAdminModel) Update(ctx context.Context, data *Admin) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Email, data.Mobile, data.Encrypt, data.GoogleSecret, data.SellerId, data.RoleId, data.CreatedAt, data.UpdatedAt, data.Status, data.Id)
	return err
}

func (m *defaultAdminModel) tableName() string {
	return m.table
}
