package logic

import (
	"context"
	"errors"
	"fmt"
	"zero-demo/user-rpc/model"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserModelLogic {
	return &GetUserModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserModelLogic) GetUserModel(in *pb.GetUserModelReq) (*pb.GetUserModelResp, error) {
	m := map[int32]string{
		1: "张三RPC",
		2: "李四RPC",
	}

	nickname := "unknown"
	if name, ok := m[int32(in.Id)]; ok {
		nickname = name
	}
	fmt.Println("NickName:", nickname)

	mikeUser, err := l.svcCtx.MikeUserModel.FindOne(l.ctx, in.Id)
	fmt.Println("===============1==================")
	fmt.Println(mikeUser)
	fmt.Println("===============2==================")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	if mikeUser == nil {
		return nil, errors.New("用户不存在")
	}

	return &pb.GetUserModelResp{
		Id:       in.Id,
		Nickname: nickname + ":db:" + mikeUser.Nickname,
		UserModel: &pb.UserModel{
			UserId:   in.Id,
			Nickname: nickname + "11111",
		},
	}, nil
}
