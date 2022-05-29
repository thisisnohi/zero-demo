package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMikeUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMikeUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMikeUserDataLogic {
	return &DelMikeUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMikeUserDataLogic) DelMikeUserData(in *pb.DelMikeUserDataReq) (*pb.DelMikeUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelMikeUserDataResp{}, nil
}
