package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMIKESTORELogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMIKESTORELogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMIKESTORELogic {
	return &DelMIKESTORELogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMIKESTORELogic) DelMIKESTORE(in *pb.DelMIKESTOREReq) (*pb.DelMIKESTOREResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelMIKESTOREResp{}, nil
}
