package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminRolePermissionModel = (*customAdminRolePermissionModel)(nil)

type (
	// AdminRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminRolePermissionModel.
	AdminRolePermissionModel interface {
		adminRolePermissionModel
	}

	customAdminRolePermissionModel struct {
		*defaultAdminRolePermissionModel
	}
)

// NewAdminRolePermissionModel returns a model for the database table.
func NewAdminRolePermissionModel(conn sqlx.SqlConn) AdminRolePermissionModel {
	return &customAdminRolePermissionModel{
		defaultAdminRolePermissionModel: newAdminRolePermissionModel(conn),
	}
}
