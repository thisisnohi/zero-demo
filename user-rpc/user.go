package main

import (
	"context"
	"flag"
	"fmt"

	"zero-demo/user-rpc/internal/config"
	"zero-demo/user-rpc/internal/server"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUsercenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(TestRpcIntercepter)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestRpcIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("TestRpcIntercepter ==> start")
	fmt.Printf("TestRpcIntercepter ==> req %+v \n", req)
	resp, err = handler(ctx, req)
	fmt.Println("TestRpcIntercepter ==> end")
	return resp, err
}
