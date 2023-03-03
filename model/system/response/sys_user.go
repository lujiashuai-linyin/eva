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
	Avatar    string `json:"avatar"`
	ExpiresAt int64  `json:"expiresAt"`
}
