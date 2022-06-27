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
	oauthRefreshTokensFieldNames          = builder.RawFieldNames(&OauthRefreshTokens{})
	oauthRefreshTokensRows                = strings.Join(oauthRefreshTokensFieldNames, ",")
	oauthRefreshTokensRowsExpectAutoSet   = strings.Join(stringx.Remove(oauthRefreshTokensFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	oauthRefreshTokensRowsWithPlaceHolder = strings.Join(stringx.Remove(oauthRefreshTokensFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	oauthRefreshTokensModel interface {
		Insert(ctx context.Context, data *OauthRefreshTokens) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*OauthRefreshTokens, error)
		Update(ctx context.Context, newData *OauthRefreshTokens) error
		Delete(ctx context.Context, id string) error
	}

	defaultOauthRefreshTokensModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OauthRefreshTokens struct {
		Id            string       `db:"id"`
		AccessTokenId string       `db:"access_token_id"`
		Revoked       int64        `db:"revoked"`
		ExpiresAt     sql.NullTime `db:"expires_at"`
	}
)

func newOauthRefreshTokensModel(conn sqlx.SqlConn) *defaultOauthRefreshTokensModel {
	return &defaultOauthRefreshTokensModel{
		conn:  conn,
		table: "`oauth_refresh_tokens`",
	}
}

func (m *defaultOauthRefreshTokensModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOauthRefreshTokensModel) FindOne(ctx context.Context, id string) (*OauthRefreshTokens, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", oauthRefreshTokensRows, m.table)
	var resp OauthRefreshTokens
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

func (m *defaultOauthRefreshTokensModel) Insert(ctx context.Context, data *OauthRefreshTokens) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, oauthRefreshTokensRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.AccessTokenId, data.Revoked, data.ExpiresAt)
	return ret, err
}

func (m *defaultOauthRefreshTokensModel) Update(ctx context.Context, data *OauthRefreshTokens) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, oauthRefreshTokensRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AccessTokenId, data.Revoked, data.ExpiresAt, data.Id)
	return err
}

func (m *defaultOauthRefreshTokensModel) tableName() string {
	return m.table
}