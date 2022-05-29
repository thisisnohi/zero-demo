package svc

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/url"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-api/model"
	"zero-demo/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                config.Config
	TestMiddleware        rest.Middleware
	MikeUserModel         model.MikeUserModel
	MikeUserDataModel     model.MikeUserDataModel
	UserRpcClient         usercenter.Usercenter
	UserEndpointRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		TestMiddleware:    middleware.NewTestMiddleware().Handle,
		MikeUserModel:     model.NewMikeUserModel(conn, c.CacheRedis),
		MikeUserDataModel: model.NewMikeUserDataModel(conn, c.CacheRedis),
		UserRpcClient:     usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
		UserEndpointRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.EndpointsRpcConf, zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			fmt.Println("UserEndpointRpcClient.WithUnaryClientInterceptor ==> start")
			fmt.Printf("UserEndpointRpcClient.WithUnaryClientInterceptor ==> req %+v", req)

			// metadata
			md := metadata.New(map[string]string{"username": url.QueryEscape("metadata.username=中文")})
			ctx = metadata.NewOutgoingContext(ctx, md)

			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return err
			}
			fmt.Printf("UserEndpointRpcClient.WithUnaryClientInterceptor ==> resp %+v", reply)
			fmt.Println("UserEndpointRpcClient.WithUnaryClientInterceptor ==> end")
			return nil
		}))),
	}
}
