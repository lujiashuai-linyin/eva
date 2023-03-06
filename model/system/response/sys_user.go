package response

import (
	"eva/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	ExpiresAt int64  `json:"expiresAt"`
}
