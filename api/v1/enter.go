package v1

import (
	"eva/api/v1/example"
	"eva/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	FileApiGroup   example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
