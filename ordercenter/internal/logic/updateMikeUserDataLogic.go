package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMikeUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMikeUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMikeUserDataLogic {
	return &UpdateMikeUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMikeUserDataLogic) UpdateMikeUserData(in *pb.UpdateMikeUserDataReq) (*pb.UpdateMikeUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMikeUserDataResp{}, nil
}
