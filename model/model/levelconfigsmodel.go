package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LevelConfigsModel = (*customLevelConfigsModel)(nil)

type (
	// LevelConfigsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLevelConfigsModel.
	LevelConfigsModel interface {
		levelConfigsModel
	}

	customLevelConfigsModel struct {
		*defaultLevelConfigsModel
	}
)

// NewLevelConfigsModel returns a model for the database table.
func NewLevelConfigsModel(conn sqlx.SqlConn) LevelConfigsModel {
	return &customLevelConfigsModel{
		defaultLevelConfigsModel: newLevelConfigsModel(conn),
	}
}
