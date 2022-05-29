package user

import (
	"context"
	"fmt"
	"zero-demo/user-api/model"
	"zero-demo/user-rpc/pb"

	"github.com/pkg/errors"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	logx.Info("this is info.........======")

	if err := l.testOne(); err != nil {
		logx.Errorf("err %+v \n", err)
	}

	m := map[int32]string{
		1: "张三",
		2: "李四",
	}

	nickname := "unknown"
	if name, ok := m[int32(req.UserId)]; ok {
		nickname = name
	}
	fmt.Println("NickName:", nickname)

	mikeUser, err := l.svcCtx.MikeUserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	if mikeUser == nil {
		return nil, errors.New("用户不存在")
	}

	//return &types.UserInfoResp{
	//	UserId:   mikeUser.Id,
	//	NickName: mikeUser.Nickname,
	//}, nil

	//// rpc
	fmt.Println("============rpc===============")
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	fmt.Println("============rpc end ===============", userResp)
	if err != nil {
		fmt.Println("==== error =====")
		return nil, err
	}
	fmt.Println("===> 结束，返回", userResp.Id, "  n1:", userResp.Nickname, "  n2:", userResp.Nickname)

	// endpotin rpc
	fmt.Println("============endpotin===============")
	userRespEndpoint, err := l.svcCtx.UserEndpointRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	fmt.Println("============userRespEndpoint end ===============", userRespEndpoint)
	if err != nil {
		fmt.Println("==== error =====")
		return nil, err
	}

	fmt.Println("===> 结束，返回", userResp.Id, "  n1:", userResp.Nickname, "  n2:", userResp.Nickname)

	return &types.UserInfoResp{
		UserId:   req.UserId,
		NickName: nickname + "===" + userResp.Nickname + ":userRespEndpoint:" + userRespEndpoint.Nickname,
	}, nil

}

func (l *UserInfoLogic) testOne() error {
	return l.testTwo()
}
func (l *UserInfoLogic) testTwo() error {
	return l.testThree()
}

func (l *UserInfoLogic) testThree() error {
	fmt.Println("=========")
	return errors.Wrap(errors.New("这是一个错误，请看堆栈..."), "enen")
}
