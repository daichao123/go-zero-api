package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OauthPersonalAccessClientsModel = (*customOauthPersonalAccessClientsModel)(nil)

type (
	// OauthPersonalAccessClientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOauthPersonalAccessClientsModel.
	OauthPersonalAccessClientsModel interface {
		oauthPersonalAccessClientsModel
	}

	customOauthPersonalAccessClientsModel struct {
		*defaultOauthPersonalAccessClientsModel
	}
)

// NewOauthPersonalAccessClientsModel returns a model for the database table.
func NewOauthPersonalAccessClientsModel(conn sqlx.SqlConn) OauthPersonalAccessClientsModel {
	return &customOauthPersonalAccessClientsModel{
		defaultOauthPersonalAccessClientsModel: newOauthPersonalAccessClientsModel(conn),
	}
}
