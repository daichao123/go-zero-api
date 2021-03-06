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
	userTemporaryRelationsFieldNames          = builder.RawFieldNames(&UserTemporaryRelations{})
	userTemporaryRelationsRows                = strings.Join(userTemporaryRelationsFieldNames, ",")
	userTemporaryRelationsRowsExpectAutoSet   = strings.Join(stringx.Remove(userTemporaryRelationsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userTemporaryRelationsRowsWithPlaceHolder = strings.Join(stringx.Remove(userTemporaryRelationsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	userTemporaryRelationsModel interface {
		Insert(ctx context.Context, data *UserTemporaryRelations) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserTemporaryRelations, error)
		FindOneByUserId(ctx context.Context, userId int64) (*UserTemporaryRelations, error)
		Update(ctx context.Context, newData *UserTemporaryRelations) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTemporaryRelationsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserTemporaryRelations struct {
		Id           int64        `db:"id"`
		UserId       int64        `db:"user_id"`        // 用户id
		ParentUserId int64        `db:"parent_user_id"` // 父id（绑定用户）
		TeamNumber   int64        `db:"team_number"`
		CreatedAt    sql.NullTime `db:"created_at"`
		UpdatedAt    sql.NullTime `db:"updated_at"`
	}
)

func newUserTemporaryRelationsModel(conn sqlx.SqlConn) *defaultUserTemporaryRelationsModel {
	return &defaultUserTemporaryRelationsModel{
		conn:  conn,
		table: "`user_temporary_relations`",
	}
}

func (m *defaultUserTemporaryRelationsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserTemporaryRelationsModel) FindOne(ctx context.Context, id int64) (*UserTemporaryRelations, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userTemporaryRelationsRows, m.table)
	var resp UserTemporaryRelations
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

func (m *defaultUserTemporaryRelationsModel) FindOneByUserId(ctx context.Context, userId int64) (*UserTemporaryRelations, error) {
	var resp UserTemporaryRelations
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userTemporaryRelationsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserTemporaryRelationsModel) Insert(ctx context.Context, data *UserTemporaryRelations) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userTemporaryRelationsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ParentUserId, data.TeamNumber, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultUserTemporaryRelationsModel) Update(ctx context.Context, newData *UserTemporaryRelations) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userTemporaryRelationsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserId, newData.ParentUserId, newData.TeamNumber, newData.CreatedAt, newData.UpdatedAt, newData.Id)
	return err
}

func (m *defaultUserTemporaryRelationsModel) tableName() string {
	return m.table
}
