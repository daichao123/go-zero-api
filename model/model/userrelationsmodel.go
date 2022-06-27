package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserRelationsModel = (*customUserRelationsModel)(nil)

type (
	// UserRelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRelationsModel.
	UserRelationsModel interface {
		userRelationsModel
	}

	customUserRelationsModel struct {
		*defaultUserRelationsModel
	}
)

// NewUserRelationsModel returns a model for the database table.
func NewUserRelationsModel(conn sqlx.SqlConn) UserRelationsModel {
	return &customUserRelationsModel{
		defaultUserRelationsModel: newUserRelationsModel(conn),
	}
}
