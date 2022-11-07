package common

import "eva/global"

type Banner struct {
	global.EVA_MODEL
	ServiceType   string `json:"service_type"`
	BannerType    string `json:"banner_type"`
	BannerUrl     string `json:"banner_url"`
	BannerTitle   string `json:"banner_title"`
	BannerJumpUrl string `json:"banner_jump_url"`
}
