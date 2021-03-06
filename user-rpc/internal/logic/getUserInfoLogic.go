package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net/url"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		if len(tmp) > 0 {
			str, _ := url.QueryUnescape(tmp[0])
			fmt.Printf("GetUserInfo.metadata ==>%+v\n", str)
		}
		fmt.Printf("GetUserInfo.metadata ==> %+v", tmp)
	}

	m := map[int32]string{
		1: "张三RPC",
		2: "李四RPC",
	}

	nickname := "unknown"
	if name, ok := m[int32(in.Id)]; ok {
		nickname = name
	}
	fmt.Println("NickName:", nickname)
	return &pb.GetUserInfoResp{
		Id:       in.Id,
		Nickname: nickname,
	}, nil
}
