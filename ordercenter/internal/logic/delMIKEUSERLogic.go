package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMIKEUSERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMIKEUSERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMIKEUSERLogic {
	return &DelMIKEUSERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMIKEUSERLogic) DelMIKEUSER(in *pb.DelMIKEUSERReq) (*pb.DelMIKEUSERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelMIKEUSERResp{}, nil
}
