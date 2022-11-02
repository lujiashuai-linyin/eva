package example

import (
	"eva/global"
	"eva/model/system"
)

type ExaFileUploadAndDownload struct {
	global.EVA_MODEL
	Name      string `json:"name" gorm:"comment:文件名"` // 文件名
	Url       string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag       string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key       string `json:"key" gorm:"comment:编号"`   // 编号
	SysUserID int
	SysUser   system.SysUser
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
