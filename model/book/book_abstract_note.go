package book

import "eva/global"

type BookAbstractNode struct {
	global.EVA_MODEL
	BookID       int    `json:"book_id" gorm:"column: bool_id;comment:书id"`
	UserID       int    `json:"user_id" gorm:"column: user_id;comment:用户id"`
	ChapterID    int    `json:"chapter_id" gorm:"column: chapter_id;comment:章节id"`
	Type         int    `json:"type" gorm:"column: type;comment:1、摘要-无标题内容，2、笔记"`
	Title        string `json:"title"  form:"title" gorm:"column:title;type:varchar(200);comment:title"`
	Content      string `json:"content" form:"content" gorm:"column:content;type:text;comment:content"`
	OriginalText string `json:"original_text" form:"original_text" gorm:"column:original_text;type:text;comment:content"`
	Extra        string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}

func (BookAbstractNode) TableName() string {
	return "book_abstract_nodes"
}
