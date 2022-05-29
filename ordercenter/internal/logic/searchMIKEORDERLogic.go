package logic

import (
	"context"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMIKEORDERLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMIKEORDERLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMIKEORDERLogic {
	return &SearchMIKEORDERLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMIKEORDERLogic) SearchMIKEORDER(in *pb.SearchMIKEORDERReq) (*pb.SearchMIKEORDERResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchMIKEORDERResp{}, nil
}
