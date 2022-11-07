package book

import "eva/global"

type BookType struct {
	global.EVA_MODEL
	TypeName string `json:"type_name" form:"type_name" gorm:"column:type_name;type:varchar(10);comment:标签"`
}
