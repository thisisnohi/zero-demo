package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMIKEUSERByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMIKEUSERByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMIKEUSERByIdLogic {
	return &GetMIKEUSERByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMIKEUSERByIdLogic) GetMIKEUSERById(in *pb.GetMIKEUSERByIdReq) (*pb.GetMIKEUSERByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMIKEUSERByIdResp{}, nil
}
