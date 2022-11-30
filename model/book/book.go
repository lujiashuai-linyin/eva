package book

import "eva/global"

type BookInfo struct {
	global.EVA_MODEL
	TitleImage   string `json:"title_page"  form:"title_page" gorm:"column:title_page;default:http://127.0.0.1:9999/uploads/file/book_default_title_image.jpg;comment:封面"`
	BookName     string `json:"book_name" form:"book_name" gorm:"column:book_name;type:varchar(50);comment:书名"`
	Type         string `json:"type" form:"type" gorm:"column:type;type:varchar(10);comment:类型"`
	LabelList    string `json:"label_list"  form:"label_list" gorm:"column:label_list;type:varchar(100);comment:label_list"`
	Author       string `json:"author" form:"author" gorm:"column:author;type:varchar(10);comment:作者"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;type:varchar(500);comment:简介"`
	ChapterList  string `json:"chapter_list" form:"chapter_list" gorm:"column:chapter_list;comment:章节"`
	// 订阅数量
	// 收藏数量
	// 点赞数量
	Extra string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}

func (BookInfo) TableName() string {
	return "book_infos"
}
