package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMIKEUSERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMIKEUSERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMIKEUSERLogic {
	return &AddMIKEUSERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------MIKEUSER-----------------------
func (l *AddMIKEUSERLogic) AddMIKEUSER(in *pb.AddMIKEUSERReq) (*pb.AddMIKEUSERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMIKEUSERResp{}, nil
}
