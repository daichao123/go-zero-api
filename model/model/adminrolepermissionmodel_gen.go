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
	adminRolePermissionFieldNames          = builder.RawFieldNames(&AdminRolePermission{})
	adminRolePermissionRows                = strings.Join(adminRolePermissionFieldNames, ",")
	adminRolePermissionRowsExpectAutoSet   = strings.Join(stringx.Remove(adminRolePermissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	adminRolePermissionRowsWithPlaceHolder = strings.Join(stringx.Remove(adminRolePermissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	adminRolePermissionModel interface {
		Insert(ctx context.Context, data *AdminRolePermission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AdminRolePermission, error)
		Update(ctx context.Context, newData *AdminRolePermission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminRolePermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AdminRolePermission struct {
		Id           int64        `db:"id"`
		PermissionId int64        `db:"permission_id"`
		RoleId       int64        `db:"role_id"`
		CreatedAt    sql.NullTime `db:"created_at"`
		UpdatedAt    sql.NullTime `db:"updated_at"`
	}
)

func newAdminRolePermissionModel(conn sqlx.SqlConn) *defaultAdminRolePermissionModel {
	return &defaultAdminRolePermissionModel{
		conn:  conn,
		table: "`admin_role_permission`",
	}
}

func (m *defaultAdminRolePermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAdminRolePermissionModel) FindOne(ctx context.Context, id int64) (*AdminRolePermission, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRolePermissionRows, m.table)
	var resp AdminRolePermission
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

func (m *defaultAdminRolePermissionModel) Insert(ctx context.Context, data *AdminRolePermission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, adminRolePermissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PermissionId, data.RoleId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultAdminRolePermissionModel) Update(ctx context.Context, data *AdminRolePermission) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminRolePermissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PermissionId, data.RoleId, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultAdminRolePermissionModel) tableName() string {
	return m.table
}