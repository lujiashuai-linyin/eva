package anime_video

import "eva/global"

type AnimeInfo struct {
	global.EVA_MODEL
	TitleImage     string `json:"title_page"  form:"title_page" gorm:"column:title_page;default:http://127.0.0.1:9999/uploads/file/book_default_title_image.jpg;type:varchar(600);comment:封面"`
	AnimeName      string `json:"anime_name" form:"anime_name" gorm:"column:anime_name;type:varchar(50);comment:动画名称"`
	Region         string `json:"region" form:"region" gorm:"column:region;type:varchar(10);comment:地区"`
	Score          string `json:"score" form:"score" gorm:"column:score;type:varchar(10);comment:分数"`
	Language       string `json:"language" form:"language" gorm:"column:language;type:varchar(10);comment:语言"`
	Years          string `json:"years" form:"years" gorm:"column:years;type:varchar(10);comment:发行年份"`
	LabelList      string `json:"label_list"  form:"label_list" gorm:"column:label_list;type:varchar(100);comment:label_list"`
	AuthorRelated  string `json:"author_related" form:"author_related" gorm:"column:author_related;type:varchar(500);comment:作者相关"`
	UpdateProgress string `json:"update_progress" form:"update_progress" gorm:"column:update_progress;type:varchar(10);comment:更新进度"`
	Introduction   string `json:"introduction" form:"introduction" gorm:"column:introduction;type:varchar(1000);comment:简介"`
	ChapterList    string `json:"chapter_list" form:"chapter_list" gorm:"column:chapter_list;type:varchar(300);comment:剧集"`
	Extra          string `json:"extra" form:"extra" gorm:"column:extra;type:varchar(300);comment:extra"`
}

func (AnimeInfo) TableName() string {
	return "anime_info"
}
