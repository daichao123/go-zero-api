package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MigrationsModel = (*customMigrationsModel)(nil)

type (
	// MigrationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMigrationsModel.
	MigrationsModel interface {
		migrationsModel
	}

	customMigrationsModel struct {
		*defaultMigrationsModel
	}
)

// NewMigrationsModel returns a model for the database table.
func NewMigrationsModel(conn sqlx.SqlConn) MigrationsModel {
	return &customMigrationsModel{
		defaultMigrationsModel: newMigrationsModel(conn),
	}
}
