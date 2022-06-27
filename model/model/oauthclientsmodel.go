package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OauthClientsModel = (*customOauthClientsModel)(nil)

type (
	// OauthClientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthClientsModel.
	OauthClientsModel interface {
		oauthClientsModel
	}

	customOauthClientsModel struct {
		*defaultOauthClientsModel
	}
)

// NewOauthClientsModel returns a model for the database table.
func NewOauthClientsModel(conn sqlx.SqlConn) OauthClientsModel {
	return &customOauthClientsModel{
		defaultOauthClientsModel: newOauthClientsModel(conn),
	}
}
