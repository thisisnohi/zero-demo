// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"zero-demo/user-rpc/internal/logic"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}