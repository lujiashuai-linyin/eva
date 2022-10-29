package book

import "eva/global"

type Chapter struct {
	global.EVA_MODEL
	BookId       int    `json:"book_id" form:"book_id" gorm:"index;column:book_id;comment:book_id"`
	ChapterName  string `json:"chapter_name" form:"chapter_name" gorm:"column:chapter_name;comment:chapter_name"`
	ChapterOrder int    `json:"chapter_order" form:"chapter_order" gorm:"column:chapter_order;comment:chapter_order"`
	Content      string `json:"content" form:"content" gorm:"column:content;type:text;comment:content"`
}
