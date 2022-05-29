package logic

import (
	"context"
	"fmt"

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
	return &pb.GetUserModelResp{
		Id:       in.Id,
		Nickname: nickname,
		UserModel: &pb.UserModel{
			UserId:   in.Id,
			Nickname: nickname + "11111",
		},
	}, nil
}
