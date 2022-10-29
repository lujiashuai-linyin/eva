package game

import "eva/global"

type GameInfo struct {
	global.EVA_MODEL
	TitleImage   string `json:"title_page"  form:"title_page" gorm:"column:title_page;default:http://127.0.0.1:9999/uploads/file/default_game_image.jpeg;comment:封面"`
	GameName     string `json:"game_name" form:"game_name" gorm:"index;column:game_name;type:varchar(50);comment:game_name"`
	PanUrl       string `json:"pan_url" form:"pan_url" gorm:"column:pan_url;type:varchar(200);comment:pan_url"`
	PanCode      string `json:"pan_code" form:"pan_code" gorm:"column:pan_code;type:varchar(6);comment:pan_code"`
	RelieveCode  string `json:"relieve_code" form:"relieve_code" gorm:"column:relieve_code;type:varchar(6);comment:relieve_code"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;type:varchar(500);comment:简介"`
	Extra        string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}
