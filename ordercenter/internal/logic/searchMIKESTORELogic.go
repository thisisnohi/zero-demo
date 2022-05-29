package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMIKESTORELogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMIKESTORELogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMIKESTORELogic {
	return &SearchMIKESTORELogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMIKESTORELogic) SearchMIKESTORE(in *pb.SearchMIKESTOREReq) (*pb.SearchMIKESTOREResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchMIKESTOREResp{}, nil
}
