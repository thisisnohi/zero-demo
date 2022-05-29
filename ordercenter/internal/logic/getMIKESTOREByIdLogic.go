package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMIKESTOREByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMIKESTOREByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMIKESTOREByIdLogic {
	return &GetMIKESTOREByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMIKESTOREByIdLogic) GetMIKESTOREById(in *pb.GetMIKESTOREByIdReq) (*pb.GetMIKESTOREByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMIKESTOREByIdResp{}, nil
}
