package request

import (
	"eva/model/common/request"
	"eva/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
