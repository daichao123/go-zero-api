package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminRolesModel = (*customAdminRolesModel)(nil)

type (
	// AdminRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminRolesModel.
	AdminRolesModel interface {
		adminRolesModel
	}

	customAdminRolesModel struct {
		*defaultAdminRolesModel
	}
)

// NewAdminRolesModel returns a model for the database table.
func NewAdminRolesModel(conn sqlx.SqlConn) AdminRolesModel {
	return &customAdminRolesModel{
		defaultAdminRolesModel: newAdminRolesModel(conn),
	}
}
