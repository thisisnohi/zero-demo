package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMIKEORDERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMIKEORDERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMIKEORDERLogic {
	return &AddMIKEORDERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------MIKEORDER-----------------------
func (l *AddMIKEORDERLogic) AddMIKEORDER(in *pb.AddMIKEORDERReq) (*pb.AddMIKEORDERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMIKEORDERResp{}, nil
}
