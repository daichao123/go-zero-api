package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminPermissionsModel = (*customAdminPermissionsModel)(nil)

type (
	// AdminPermissionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminPermissionsModel.
	AdminPermissionsModel interface {
		adminPermissionsModel
	}

	customAdminPermissionsModel struct {
		*defaultAdminPermissionsModel
	}
)

// NewAdminPermissionsModel returns a model for the database table.
func NewAdminPermissionsModel(conn sqlx.SqlConn) AdminPermissionsModel {
	return &customAdminPermissionsModel{
		defaultAdminPermissionsModel: newAdminPermissionsModel(conn),
	}
}
