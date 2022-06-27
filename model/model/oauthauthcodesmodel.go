package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OauthAuthCodesModel = (*customOauthAuthCodesModel)(nil)

type (
	// OauthAuthCodesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthAuthCodesModel.
	OauthAuthCodesModel interface {
		oauthAuthCodesModel
	}

	customOauthAuthCodesModel struct {
		*defaultOauthAuthCodesModel
	}
)

// NewOauthAuthCodesModel returns a model for the database table.
func NewOauthAuthCodesModel(conn sqlx.SqlConn) OauthAuthCodesModel {
	return &customOauthAuthCodesModel{
		defaultOauthAuthCodesModel: newOauthAuthCodesModel(conn),
	}
}
