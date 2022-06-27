package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminRoleMenuModel = (*customAdminRoleMenuModel)(nil)

type (
	// AdminRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminRoleMenuModel.
	AdminRoleMenuModel interface {
		adminRoleMenuModel
	}

	customAdminRoleMenuModel struct {
		*defaultAdminRoleMenuModel
	}
)

// NewAdminRoleMenuModel returns a model for the database table.
func NewAdminRoleMenuModel(conn sqlx.SqlConn) AdminRoleMenuModel {
	return &customAdminRoleMenuModel{
		defaultAdminRoleMenuModel: newAdminRoleMenuModel(conn),
	}
}
