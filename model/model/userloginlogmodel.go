package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserLoginLogModel = (*customUserLoginLogModel)(nil)

type (
	// UserLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLoginLogModel.
	UserLoginLogModel interface {
		userLoginLogModel
	}

	customUserLoginLogModel struct {
		*defaultUserLoginLogModel
	}
)

// NewUserLoginLogModel returns a model for the database table.
func NewUserLoginLogModel(conn sqlx.SqlConn) UserLoginLogModel {
	return &customUserLoginLogModel{
		defaultUserLoginLogModel: newUserLoginLogModel(conn),
	}
}
