package handler

import "eva/biz/handler/private"

type RpcGroup struct {
	BytestRpcGroup private.RpcGroup
}

var RpcGroupApp = new(RpcGroup)
