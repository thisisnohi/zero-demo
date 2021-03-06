package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf      zrpc.RpcClientConf
	EndpointsRpcConf zrpc.RpcClientConf
	Mysql            struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
