package logic

import (
	"context"
	"errors"
	"go-zero-test/user-api/internal/svc"
	"go-zero-test/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return &types.UserInfoResp{
		Username: user.Username,
		UserId:   user.Id,
	}, nil

}
