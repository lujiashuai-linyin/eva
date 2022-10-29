package system

import (
	"eva/global"
)

type JwtBlacklist struct {
	global.EVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
