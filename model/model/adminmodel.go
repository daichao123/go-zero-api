package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

// NewAdminModel returns a model for the database table.
func NewAdminModel(conn sqlx.SqlConn) AdminModel {
	return &customAdminModel{
		defaultAdminModel: newAdminModel(conn),
	}
}
