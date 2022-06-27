package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InvitationCodeModel = (*customInvitationCodeModel)(nil)

type (
	// InvitationCodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvitationCodeModel.
	InvitationCodeModel interface {
		invitationCodeModel
	}

	customInvitationCodeModel struct {
		*defaultInvitationCodeModel
	}
)

// NewInvitationCodeModel returns a model for the database table.
func NewInvitationCodeModel(conn sqlx.SqlConn) InvitationCodeModel {
	return &customInvitationCodeModel{
		defaultInvitationCodeModel: newInvitationCodeModel(conn),
	}
}
