package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-test/model/model"

	"go-zero-test/user-api/internal/svc"
	"go-zero-test/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LevelUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLevelUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LevelUpLogic {
	return &LevelUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LevelUpLogic) LevelUp(req *types.UserLevelUpReq) (resp *types.UserLevelUpResp, err error) {
	//users := model.Users{}
	levelResult, _ := l.svcCtx.UserLevelsModel.FindOneByWhere(l.ctx, "user_id", req.UserId)
	if levelResult == nil { //如果为空
		userLevel := &model.UserLevels{
			UserId:    req.UserId,
			LevelType: "Vip等级",
			Level:     req.Level,
			LevelName: "普通Vip",
		}
		insert, err := l.svcCtx.UserLevelsModel.Insert(l.ctx, userLevel)
		affected, _ := insert.RowsAffected()
		if err != nil || affected <= 0 {
			return &types.UserLevelUpResp{Code: 10003, Message: "插入记录失败"}, err
		}
	} else {
		//重写insert 方法 使其在事务中使用同一个会话、保证事务一致性
		if err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			levelResult.Level = req.Level
			levelResult.UserId = req.UserId
			if err := l.svcCtx.UserLevelsModel.TransUpdate(ctx, levelResult, session); err != nil {
				return err
			}
			users, _ := l.svcCtx.UserModel.FindOne(ctx, req.UserId)
			users.RegisterIp = "127.0.0.1"
			fmt.Println("debug：", users)
			if err := l.svcCtx.UserModel.TransUpdate(ctx, users, session); err != nil {
				return err
			}
			return nil
		}); err != nil {
			fmt.Println(err)
			return nil, errors.New("修改失败")
		}
	}
	return &types.UserLevelUpResp{Code: 200, Message: "修改成功"}, err

}
