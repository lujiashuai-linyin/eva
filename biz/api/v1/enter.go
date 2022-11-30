package v1

import (
	"eva/biz/api/v1/example"
	"eva/biz/api/v1/private"
	"eva/biz/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	FileApiGroup    example.ApiGroup
	PrivateApiGroup private.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
