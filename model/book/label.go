package book

import "eva/global"

type Label struct {
	global.EVA_MODEL
	LabelName string `json:"label_name" form:"label_name" gorm:"column:label_name;type:varchar(10);comment:标签名称"`
}
