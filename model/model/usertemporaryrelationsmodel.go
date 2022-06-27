package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserTemporaryRelationsModel = (*customUserTemporaryRelationsModel)(nil)

type (
	// UserTemporaryRelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTemporaryRelationsModel.
	UserTemporaryRelationsModel interface {
		userTemporaryRelationsModel
	}

	customUserTemporaryRelationsModel struct {
		*defaultUserTemporaryRelationsModel
	}
)

// NewUserTemporaryRelationsModel returns a model for the database table.
func NewUserTemporaryRelationsModel(conn sqlx.SqlConn) UserTemporaryRelationsModel {
	return &customUserTemporaryRelationsModel{
		defaultUserTemporaryRelationsModel: newUserTemporaryRelationsModel(conn),
	}
}
