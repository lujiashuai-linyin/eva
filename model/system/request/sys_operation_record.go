package request

import (
	"eva/model/common/request"
	"eva/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
