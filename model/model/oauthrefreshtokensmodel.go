package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OauthRefreshTokensModel = (*customOauthRefreshTokensModel)(nil)

type (
	// OauthRefreshTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthRefreshTokensModel.
	OauthRefreshTokensModel interface {
		oauthRefreshTokensModel
	}

	customOauthRefreshTokensModel struct {
		*defaultOauthRefreshTokensModel
	}
)

// NewOauthRefreshTokensModel returns a model for the database table.
func NewOauthRefreshTokensModel(conn sqlx.SqlConn) OauthRefreshTokensModel {
	return &customOauthRefreshTokensModel{
		defaultOauthRefreshTokensModel: newOauthRefreshTokensModel(conn),
	}
}
