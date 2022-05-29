package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMikeUserDataByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMikeUserDataByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMikeUserDataByIdLogic {
	return &GetMikeUserDataByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMikeUserDataByIdLogic) GetMikeUserDataById(in *pb.GetMikeUserDataByIdReq) (*pb.GetMikeUserDataByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMikeUserDataByIdResp{}, nil
}
