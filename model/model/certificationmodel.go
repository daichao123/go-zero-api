package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CertificationModel = (*customCertificationModel)(nil)

type (
	// CertificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCertificationModel.
	CertificationModel interface {
		certificationModel
	}

	customCertificationModel struct {
		*defaultCertificationModel
	}
)

// NewCertificationModel returns a model for the database table.
func NewCertificationModel(conn sqlx.SqlConn) CertificationModel {
	return &customCertificationModel{
		defaultCertificationModel: newCertificationModel(conn),
	}
}
