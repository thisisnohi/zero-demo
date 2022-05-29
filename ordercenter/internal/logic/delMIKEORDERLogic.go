package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMIKEORDERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMIKEORDERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMIKEORDERLogic {
	return &DelMIKEORDERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMIKEORDERLogic) DelMIKEORDER(in *pb.DelMIKEORDERReq) (*pb.DelMIKEORDERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelMIKEORDERResp{}, nil
}
