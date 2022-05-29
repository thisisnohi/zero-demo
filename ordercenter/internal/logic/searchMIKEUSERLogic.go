package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMIKEUSERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMIKEUSERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMIKEUSERLogic {
	return &SearchMIKEUSERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMIKEUSERLogic) SearchMIKEUSER(in *pb.SearchMIKEUSERReq) (*pb.SearchMIKEUSERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchMIKEUSERResp{}, nil
}
