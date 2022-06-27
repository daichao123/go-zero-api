package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminMenuPermissionModel = (*customAdminMenuPermissionModel)(nil)

type (
	// AdminMenuPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminMenuPermissionModel.
	AdminMenuPermissionModel interface {
		adminMenuPermissionModel
	}

	customAdminMenuPermissionModel struct {
		*defaultAdminMenuPermissionModel
	}
)

// NewAdminMenuPermissionModel returns a model for the database table.
func NewAdminMenuPermissionModel(conn sqlx.SqlConn) AdminMenuPermissionModel {
	return &customAdminMenuPermissionModel{
		defaultAdminMenuPermissionModel: newAdminMenuPermissionModel(conn),
	}
}
