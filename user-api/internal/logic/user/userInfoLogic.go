package user

import (
	"context"
	"errors"
	"fmt"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-api/model"

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

	m := map[int32]string{
		1: "张三",
		2: "李四",
	}

	nickname := "unknown"
	if name, ok := m[int32(req.UserId)]; ok {
		nickname = name
	}
	fmt.Println("NickName:", nickname)

	mikeUser, err := l.svcCtx.MikeUserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	if mikeUser == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInfoResp{
		UserId:   mikeUser.Id,
		NickName: mikeUser.Nickname,
	}, nil
}
