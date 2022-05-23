package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	var _UserId int64

	if err := l.svcCtx.MikeUserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.MikeUser{}
		user.Mobile = req.Mobile
		user.Nickname = req.Nickname
		// 添加user
		dbResult, err := l.svcCtx.MikeUserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		userId, _ := dbResult.LastInsertId()
		fmt.Println("userId:", userId)
		// 添加userData
		userData := &model.MikeUserData{
			UserId: userId,
			Data:   "123123",
		}
		_UserId = userId
		if _, err := l.svcCtx.MikeUserDataModel.TransInsert(ctx, session, userData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &types.UserCreateResp{
		Flag:   true,
		UserId: _UserId,
	}, nil

}
