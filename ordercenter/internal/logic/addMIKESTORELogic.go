package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMIKESTORELogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMIKESTORELogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMIKESTORELogic {
	return &AddMIKESTORELogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------MIKESTORE-----------------------
func (l *AddMIKESTORELogic) AddMIKESTORE(in *pb.AddMIKESTOREReq) (*pb.AddMIKESTOREResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMIKESTOREResp{}, nil
}
