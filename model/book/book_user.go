package book

import "eva/global"

type BookUser struct {
	global.EVA_MODEL
	BookID     int    `json:"book_id" gorm:"column: bool_id;comment:书id"`
	UserID     int    `json:"user_id" gorm:"column: user_id;comment:用户id"`
	Praise     bool   `json:"praise" gorm:"column: praise;comment:赞"`
	Collection bool   `json:"collection" gorm:"column: collection;comment:收藏"`
	Subscribe  bool   `json:"subscribe" gorm:"column: subscribe;comment:订阅"`
	Extra      string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}

func (BookUser) TableName() string {
	return "book_users"
}
