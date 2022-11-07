package memorandum

import (
	"eva/global"
	"eva/model/system"
)

type MemorandumTopic struct {
	global.EVA_MODEL
	SysUserID int
	SysUser   system.SysUser
	Name      string `json:"name" form:"name" gorm:"column:name;type:varchar(50);comment:主题"`
}
