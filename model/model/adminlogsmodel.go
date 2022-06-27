package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminLogsModel = (*customAdminLogsModel)(nil)

type (
	// AdminLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminLogsModel.
	AdminLogsModel interface {
		adminLogsModel
	}

	customAdminLogsModel struct {
		*defaultAdminLogsModel
	}
)

// NewAdminLogsModel returns a model for the database table.
func NewAdminLogsModel(conn sqlx.SqlConn) AdminLogsModel {
	return &customAdminLogsModel{
		defaultAdminLogsModel: newAdminLogsModel(conn),
	}
}
