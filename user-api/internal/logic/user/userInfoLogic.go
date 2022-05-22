package user

import (
	"context"
	"fmt"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

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

	return &types.UserInfoResp{
		UserId:   req.UserId,
		NickName: nickname,
	}, nil
}
