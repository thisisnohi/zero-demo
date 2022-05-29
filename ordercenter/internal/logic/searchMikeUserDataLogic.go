package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMikeUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMikeUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMikeUserDataLogic {
	return &SearchMikeUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMikeUserDataLogic) SearchMikeUserData(in *pb.SearchMikeUserDataReq) (*pb.SearchMikeUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchMikeUserDataResp{}, nil
}
