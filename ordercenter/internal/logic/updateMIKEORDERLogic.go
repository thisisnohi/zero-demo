package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMIKEORDERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMIKEORDERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMIKEORDERLogic {
	return &UpdateMIKEORDERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMIKEORDERLogic) UpdateMIKEORDER(in *pb.UpdateMIKEORDERReq) (*pb.UpdateMIKEORDERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMIKEORDERResp{}, nil
}
