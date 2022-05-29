package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMIKEUSERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMIKEUSERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMIKEUSERLogic {
	return &UpdateMIKEUSERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMIKEUSERLogic) UpdateMIKEUSER(in *pb.UpdateMIKEUSERReq) (*pb.UpdateMIKEUSERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMIKEUSERResp{}, nil
}
