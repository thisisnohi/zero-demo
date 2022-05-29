package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMIKESTORELogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMIKESTORELogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMIKESTORELogic {
	return &UpdateMIKESTORELogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMIKESTORELogic) UpdateMIKESTORE(in *pb.UpdateMIKESTOREReq) (*pb.UpdateMIKESTOREResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMIKESTOREResp{}, nil
}
