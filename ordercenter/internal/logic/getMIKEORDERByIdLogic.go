package logic

import (
	"context"
	"time"

	"zero-demo/ordercenter/internal/svc"
	"zero-demo/ordercenter/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMIKEORDERByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMIKEORDERByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMIKEORDERByIdLogic {
	return &GetMIKEORDERByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMIKEORDERByIdLogic) GetMIKEORDERById(in *pb.GetMIKEORDERByIdReq) (*pb.GetMIKEORDERByIdResp, error) {
	odrder := &pb.MIKEORDER{
		Id:         1,
		GoodsId:    1001,
		GoodsName:  "商品",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	return &pb.GetMIKEORDERByIdResp{
		MIKEORDER: odrder,
	}, nil
}
