package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminMenusModel = (*customAdminMenusModel)(nil)

type (
	// AdminMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminMenusModel.
	AdminMenusModel interface {
		adminMenusModel
	}

	customAdminMenusModel struct {
		*defaultAdminMenusModel
	}
)

// NewAdminMenusModel returns a model for the database table.
func NewAdminMenusModel(conn sqlx.SqlConn) AdminMenusModel {
	return &customAdminMenusModel{
		defaultAdminMenusModel: newAdminMenusModel(conn),
	}
}
