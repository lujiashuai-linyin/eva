package request

import (
	"eva/model/common/request"
	"eva/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
