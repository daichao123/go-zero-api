package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserLevelsModel = (*customUserLevelsModel)(nil)

type (
	// UserLevelsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLevelsModel.
	UserLevelsModel interface {
		userLevelsModel
	}

	customUserLevelsModel struct {
		*defaultUserLevelsModel
	}
)

// NewUserLevelsModel returns a model for the database table.
func NewUserLevelsModel(conn sqlx.SqlConn) UserLevelsModel {
	return &customUserLevelsModel{
		defaultUserLevelsModel: newUserLevelsModel(conn),
	}
}
