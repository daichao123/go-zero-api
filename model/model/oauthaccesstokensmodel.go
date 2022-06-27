package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OauthAccessTokensModel = (*customOauthAccessTokensModel)(nil)

type (
	// OauthAccessTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthAccessTokensModel.
	OauthAccessTokensModel interface {
		oauthAccessTokensModel
	}

	customOauthAccessTokensModel struct {
		*defaultOauthAccessTokensModel
	}
)

// NewOauthAccessTokensModel returns a model for the database table.
func NewOauthAccessTokensModel(conn sqlx.SqlConn) OauthAccessTokensModel {
	return &customOauthAccessTokensModel{
		defaultOauthAccessTokensModel: newOauthAccessTokensModel(conn),
	}
}
