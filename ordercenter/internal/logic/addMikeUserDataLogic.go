package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMikeUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMikeUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMikeUserDataLogic {
	return &AddMikeUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------mikeUserData-----------------------
func (l *AddMikeUserDataLogic) AddMikeUserData(in *pb.AddMikeUserDataReq) (*pb.AddMikeUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMikeUserDataResp{}, nil
}
